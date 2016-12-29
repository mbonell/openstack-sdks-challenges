// Using the OpenStack Nova API (Gophercloud), the create worker script launch encoding workers that
// receive a video name and container from the cloud (Swift object) and convert it into the format selected by the user.

package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"os"
	"time"
)

func main() {

	authUrl := os.Getenv("OS_AUTH_URL")
	username := os.Getenv("OS_USERNAME")
	password := os.Getenv("OS_PASSWORD")
	domain := os.Getenv("OS_DOMAIN_ID")
	region := os.Getenv("OS_REGION_NAME")

	videoContainer := os.Getenv("ORIGINAL_VIDEO_CONTAINER")
	videoName := os.Getenv("ORIGINAL_VIDEO_NAME")
	format := os.Getenv("FORMAT_TO_ENCODE")

	workerImage := os.Getenv("WORKER_SERVER_IMAGE")
	workerFlavor := os.Getenv("WORKER_SERVER_FLAVOR")
	workerNetwork := os.Getenv("WORKER_SERVER_NETWORK")

	// Validate required variables
	if videoContainer == "" || videoName == "" || format == "" {
		fmt.Println("The following env vars are required: video container, video name and format")
		return
	}

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

	// Using the compute service with the cloud
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
		Type:   "computev21",
	})

	// Script to execte after the instance creation
	userData := fmt.Sprintf(`#!/usr/bin/env bash
	wget https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/init.sh
	OS_AUTH_URL=%s OS_USERNAME=%s OS_PASSWORD=%s OS_DOMAIN_ID=%s OS_REGION_NAME=%s \
	ORIGINAL_VIDEO_CONTAINER=%s ORIGINAL_VIDEO_NAME=%s FORMAT_TO_ENCODE=%s bash init.sh`,
		authUrl,
		username,
		password,
		domain,
		region,
		videoContainer,
		videoName,
		format,
	)

	// Create an worker instance
	server, err := servers.Create(client, servers.CreateOpts{
		Name:           "worker-" + time.Now().Format("2006-01-02-15:04:05"),
		FlavorRef:      workerFlavor,
		ImageRef:       workerImage,
		Networks:       []servers.Network{servers.Network{UUID: workerNetwork}},
		SecurityGroups: []string{"worker"},
		UserData:       []byte(userData),
	}).Extract()

	if err != nil {
		fmt.Printf("Unable to create worker: %s", err)
		return
	}

	fmt.Printf("Worker created!, ID: %s", server.ID)

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
