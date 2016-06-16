<?php

require 'vendor/autoload.php';
use OpenCloud\OpenStack;

#1: Connecting to the cloud

$cloud = new OpenStack('https://YOUR_IDENTITY_CONTROLLER', array(
    'username'   => 'YOUR_USERNAME',
    'password'   => 'YOUR_PASSWORD',
    'tenantName' => 'YOUR_PROJECT'
));

$cloud->authenticate();
print_r($cloud->getCatalog()->getItems());

?>
