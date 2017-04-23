// The worker will receive the original video and container name to download it from the cloud.
// The format to encode the video (MP4, MPEG, WEBM) will be sent to the worker as well.
// Then the worker will execute the encoding task and at the end it will upload the new encoded videos to the cloud
// through the object storage API.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		log.Println("[Error] Video container, name and/or format are required!")
		return
	}

	videoContainer := args[0]
	videoName := args[1]
	format := args[2]

	// Validate compatible formats
	if format != "mp4" && format != "mpeg" && format != "webm" {
		log.Println("[Error] Format not compatible! Only mp4, mpeg and webm are supported.")
		return
	}

	provider, err := getOpenStackProvider()
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	region := os.Getenv("OS_REGION_NAME")
	objectStorage, err := getServiceObjectStorage(provider, region)
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	log.Println("[WORKER] Downloading video from the cloud...")
	objectPath := path.Join(os.TempDir(), videoName)
	err = downloadObject(objectStorage, videoContainer, videoName, objectPath)
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	log.Println("[WORKER] Starting the video encoding...")
	encodedName := fmt.Sprintf("%s-%s.%s", videoName, time.Now().Format("2006-01-02-15:04:05"), format)
	encodedPath := path.Join(os.TempDir(), encodedName)

	cmd := exec.Command("ffmpeg", "-i", objectPath, encodedPath)
	err = cmd.Start()
	if err != nil {
		log.Printf("Error during encoding execution: %s \n", err.Error())
		return
	}

	log.Println("[WORKER] Waiting for video encoding to finish...")
	err = cmd.Wait()
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	log.Println("[WORKER] Uploading encoded video to the cloud...")
	err = uploadObject(objectStorage, "encoded-videos", encodedName, encodedPath)
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	log.Println("[WORKER] The encoded task was completed with success!")
	os.Remove(objectPath)
	os.Remove(encodedPath)
}

// getOpenStackProvider reads credentials from the environment variables and
// creates a connection with the provided OpenStack cloud.
func getOpenStackProvider() (*gophercloud.ProviderClient, error) {

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	return provider, nil

}

// getServiceObjectStorage provides a object storage client.
func getServiceObjectStorage(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {

	service, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: region,
	})

	if err != nil {
		return nil, err
	}

	return service, nil

}

// downloadObject uses the object storage service to download an object from the cloud and
// stored it in the specified filesystem path.
func downloadObject(service *gophercloud.ServiceClient, containerName, objectName, path string) error {

	result := objects.Download(service, containerName, objectName, nil)
	content, err := result.ExtractContent()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content), 0644)
	return err

}

// uploadObject uses the object storage service to upload a file from a path in the filesystem to
// a specific storage container in the cloud.
func uploadObject(service *gophercloud.ServiceClient, containerName, objectName, objectPath string) error {

	f, err := os.Open(objectPath)
	defer f.Close()
	reader := bufio.NewReader(f)

	if err != nil {
		return err
	}

	options := objects.CreateOpts{
		Content: reader,
	}

	res := objects.Create(service, containerName, objectName, options)
	return res.Err

}
