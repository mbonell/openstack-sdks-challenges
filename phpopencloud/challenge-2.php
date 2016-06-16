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

#2: Listing images 

echo ("**Available images** \n");
$images = $compute->imageList();

foreach ($images as $image) {
    echo sprintf("Name: %s, ID: %d \n", $image->name, $image->id);
}
?>
