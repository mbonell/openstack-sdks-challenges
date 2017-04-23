// Using the OpenStack Nova API (Gophercloud), the create worker script launch encoding workers that
// receive a video name and container from the cloud (Swift object) and convert it into the format selected by the user.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {

	// Get cloud credentials
	authURL := os.Getenv("OS_AUTH_URL")
	username := os.Getenv("OS_USERNAME")
	password := os.Getenv("OS_PASSWORD")
	region := os.Getenv("OS_REGION_NAME")
	domain := os.Getenv("OS_DOMAIN_ID")
	tenantName := os.Getenv("OS_TENANT_NAME")

	// Get  workers' infrastructure values
	workerImage := os.Getenv("WORKER_SERVER_IMAGE")
	workerFlavor := os.Getenv("WORKER_SERVER_FLAVOR")
	workerNetwork := os.Getenv("WORKER_SERVER_NETWORK")

	// Get video values
	videoContainer := os.Getenv("ORIGINAL_VIDEO_CONTAINER")
	videoName := os.Getenv("ORIGINAL_VIDEO_NAME")
	format := os.Getenv("FORMAT_TO_ENCODE")

	// Validate required variables
	if videoContainer == "" || videoName == "" || format == "" {
		log.Println("[Error] The following env vars are required: video container, video name and format")
		return
	}

	// Validate compatible formats
	if format != "mp4" && format != "mpeg" && format != "webm" {
		log.Println("[Error] Format not compatible! Only mp4, mpeg and webm are supported")
		return
	}

	provider, err := getOpenStackProvider()
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	// Using the compute service with the cloud
	client, err := getComputeClient(provider, region)
	if err != nil {
		log.Println("[Error] " + err.Error())
		return
	}

	// Script to execute after the instance creation
	userData := fmt.Sprintf(`#!/usr/bin/env bash
	wget https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/init.sh
	OS_AUTH_URL=%s OS_USERNAME=%s OS_PASSWORD=%s OS_DOMAIN_ID=%s OS_TENANT_NAME=%s OS_REGION_NAME=%s \
	ORIGINAL_VIDEO_CONTAINER=%s ORIGINAL_VIDEO_NAME=%s FORMAT_TO_ENCODE=%s bash init.sh`,
		authURL,
		username,
		password,
		domain,
		tenantName,
		region,
		videoContainer,
		videoName,
		format,
	)

	// Create a worker instance
	server, err := servers.Create(client, servers.CreateOpts{
		Name:           "worker-" + time.Now().Format("2006-01-02-15:04:05"),
		FlavorRef:      workerFlavor,
		ImageRef:       workerImage,
		Networks:       []servers.Network{servers.Network{UUID: workerNetwork}},
		SecurityGroups: []string{"worker"},
		UserData:       []byte(userData),
	}).Extract()

	if err != nil {
		log.Println("[Error] Unable to create worker: " + err.Error())
		return
	}

	log.Println("[Success] Worker created!, ID: " + server.ID)
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

// getComputeClient provides a compute client based on the compute service version available.
func getComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {

	clientv21, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
		Type:   "computev21",
	})

	if err != nil {
		client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
			Region: region,
		})

		if err != nil {
			return client, err
		}

		return client, nil
	}

	return clientv21, nil

}
