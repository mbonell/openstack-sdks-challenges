require "fog/openstack"

#1: Connecting to the cloud

@connection_params = {
	openstack_auth_url:     "https://YOUR_IDENTITY_CONTROLLER",
	openstack_username:     "YOUR_USERNAME",
  	openstack_api_key:      "YOUR_PASSWORD",
  	openstack_project_name: "YOUR_PROJECT",
  	openstack_region: 		"YOUR_REGION", 
}

compute = Fog::Compute::OpenStack.new(@connection_params)

#4: Choosing images and flavors
image = compute.images.get "3c76334f-9644-4666-ac3c-fa090f175655"
flavor = compute.flavors.get "A1.1"

#5: Launching simple instances

#5.1: Select your external network
external_network = "f6286f9b-07f2-474b-aeb9-7e10bb0a7b00"

#5.2: Script instance post-creation
user_data = <<END
#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/fog/init.sh | bash -s --
END

#5.3: Script instance post-creation
instance = compute.servers.create name: 		'my-pet-001',
								  image_ref: 	image.id,
                                  flavor_ref: 	flavor.id,
                                  user_data: 	user_data,
                                  nics: 		[net_id: external_network]
instance.wait_for { ready? }

#5.4: List all the instances in your cloud
p compute.servers