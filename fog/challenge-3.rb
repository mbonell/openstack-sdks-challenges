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

#3: Listing flavors 
p compute.flavors