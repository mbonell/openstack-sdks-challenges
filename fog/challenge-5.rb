require "fog/openstack"

#1: Connecting to the cloud
@connection_params = {
	openstack_auth_url:     "https://YOUR_IDENTITY_CONTROLLER",
	openstack_username:     "YOUR_USERNAME",
  	openstack_api_key:      "YOUR_PASSWORD",
  	openstack_project_name: "YOUR_PROJECT",
  	openstack_region: 	"YOUR_REGION", 
}

compute = Fog::Compute::OpenStack.new(@connection_params)

#4: Choosing images and flavors
image = compute.images.get "YOUR_IMAGE_ID"
flavor = compute.flavors.get "YOUR_FLAVOR_ID"

#5: Launching a web instance

#5.1: Select your external network
external_network = "YOUR_NETWORK_ID"

#5.2: Script instance post-creation
user_data = <<END
#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/fog/init.sh | bash -s --
END

#5.3: Script instance post-creation
instance = compute.servers.create name: 	'my-web-cattle-001',
				  image_ref: 	image.id,
                                  flavor_ref: 	flavor.id,
                                  user_data: 	user_data,
                                  nics: 	[net_id: external_network]
instance.wait_for { ready? }

#5.4: List all the instances in your cloud
p compute.servers
