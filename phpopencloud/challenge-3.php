<?php

require 'vendor/autoload.php';
use OpenCloud\OpenStack;

#1: Connecting to the cloud

$cloud = new OpenStack('https://YOUR_IDENTITY_CONTROLLER', array(
    'username'   => 'YOUR_USERNAME',
    'password'   => 'YOUR_PASSWORD',
    'tenantName' => 'YOUR_PROJECT'
));

$compute = $cloud->computeService('nova', 'YOUR_REGION');

#3: Listing flavors 

echo ("**Available flavors** \n");
$flavors = $compute->flavorList();

foreach ($flavors as $flavor) {
     printf("ID: %s, Name: %s, RAM: %s, VCPUs: %s\n", $flavor->id, $flavor->name, $flavor->ram, $flavor->vcpus);
}


?>
