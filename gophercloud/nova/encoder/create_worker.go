//TODO

package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"time"
)

func main() {

	// Set cloud credentials
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://YOUR_IDENTITY_CONTROLLER",
		Username:         "YOUR_USERNAME",
		Password:         "YOUR_PASSWORD",
		DomainName:       "default",
	}

	// Create connection with the cloud
	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Using the compute service with the cloud
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	// Create an instance
	server, err := servers.Create(client, servers.CreateOpts{
		Name:       "worker-" + time.Now().String(),
		FlavorName: "m1.medium",
		ImageName:  "ubuntu-server-16.04.1",
	}).Extract()

	if err != nil {
		fmt.Println("Unable to create server: %s", err)
		return
	}

	fmt.Println("Server created!, ID: %s", server.ID)

}
