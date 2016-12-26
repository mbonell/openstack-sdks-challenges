//TODO

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

	// Set cloud credentials
	opts, err := openstack.AuthOptionsFromEnv()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Create connection with the cloud
	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Using the compute service with the cloud
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
		Type:   "computev21",
	})

	// Script to execte after the instance creation
	userData := fmt.Sprint(`#!/usr/bin/env bash
	curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/init.sh | bash -s -- \
	-OS_AUTH_URL %s -OS_PROJECT_NAME %s -OS_USERNAME %s -OS_PASSWORD %s -OS_DOMAIN_ID %s -ORIGINAL_VIDEO_FILE %s -FORMAT_TO_ENCODE %s`,
		os.Getenv("OS_REGION_NAME"),
		os.Getenv("OS_PROJECT_NAME"),
		os.Getenv("OS_USERNAME"),
		os.Getenv("OS_PASSWORD"),
		os.Getenv("OS_DOMAIN_ID"),
		os.Getenv("ORIGINAL_VIDEO_FILE"),
		os.Getenv("FORMAT_TO_ENCODE"),
	)

	// Create an worker instance
	server, err := servers.Create(client, servers.CreateOpts{
		Name:           "worker-" + time.Now().String(),
		FlavorRef:      os.Getenv("WORKER_SERVER_FLAVOR"),
		ImageRef:       os.Getenv("WORKER_SERVER_IMAGE"),
		Networks:       []servers.Network{servers.Network{UUID: os.Getenv("WORKER_SERVER_NETWORK")}},
		SecurityGroups: []string{"worker"},
		UserData:       []byte(userData),
	}).Extract()

	if err != nil {
		fmt.Printf("Unable to create worker: %s", err)
		return
	}

	fmt.Printf("Worker created!, ID: %s", server.ID)

}
