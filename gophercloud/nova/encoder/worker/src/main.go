// The worker will receive the original video and container name to download it from the cloud.
// The format to encode the video (MP4, MPEG, WEBM) will be sent to the worker as well.
// Then the worker will execute the encoding task and at the end it will upload the new encoded videos to the cloud
// though the object storage API.

package main

import (
	"bufio"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Video container, name  and/or format is required!")
		return
	}

	videoContainer := args[0]
	videoName := args[1]
	format := args[2]

	// Validate compatible formats
	if format != "mp4" && format != "mpeg" && format != "webm" {
		fmt.Println("Format not compatible! Only mp4, mpeg and webm are supported")
		return
	}

	provider, err := GetOpenStackProvider()
	if err != nil {
		fmt.Println(err)
		return
	}

	region := os.Getenv("OS_REGION_NAME")
	objectStorage, err := GetServiceObjectStorage(provider, region)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[WORKER] Downloading video from the cloud...")
	objectPath := path.Join(os.TempDir(), videoName)
	err = DownloadObject(objectStorage, videoContainer, videoName, objectPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[WORKER] Starting the video encoding...")
	encodedName := fmt.Sprintf("%s-%s.%s", videoName, time.Now().Format("2006-01-02-15:04:05"), format)
	encodedPath := path.Join(os.TempDir(), encodedName)

	cmd := exec.Command("ffmpeg", "-i", objectPath, encodedPath)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error during encoding execution: %s \n", err)
		return
	}

	fmt.Println("[WORKER] Waiting for video encoding to finish...")
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[WORKER] Uploading encoded video to the cloud...")
	err = UploadObject(objectStorage, "encoded-videos", encodedName, encodedPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[WORKER] The encoded task was completed with success!")
	os.Remove(objectPath)
	os.Remove(encodedPath)
}

func GetOpenStackProvider() (*gophercloud.ProviderClient, error) {

	// Set cloud credentials
	opts, err := openstack.AuthOptionsFromEnv()

	if err != nil {
		return nil, err
	}

	// Create connection with the cloud
	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		return nil, err
	}

	return provider, nil

}

func GetServiceObjectStorage(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {

	service, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: region,
	})

	if err != nil {
		return nil, err
	}

	return service, nil

}

func DownloadObject(service *gophercloud.ServiceClient, containerName, objectName, path string) error {

	result := objects.Download(service, containerName, objectName, nil)
	content, err := result.ExtractContent()

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(content), 0644)
	return err

}

func UploadObject(service *gophercloud.ServiceClient, containerName, objectName, objectPath string) error {

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
