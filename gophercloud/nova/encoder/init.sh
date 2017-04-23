#!/bin/sh

WORKER_BIN_URL=https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/worker/bin/worker

sudo add-apt-repository ppa:mc3man/trusty-media -y
sudo apt-get update
sudo apt-get install ffmpeg -y

# Setting cloud environment variables
export OS_AUTH_URL=$OS_AUTH_URL
export OS_REGION_NAME=$OS_REGION_NAME
export OS_USERNAME=$OS_USERNAME
export OS_PASSWORD=$OS_PASSWORD
export OS_DOMAIN_ID=$OS_DOMAIN_ID
export OS_TENANT_NAME=$OS_TENANT_NAME

# Downloading the worker binary and running it
wget $WORKER_BIN_URL
sudo chmod +x worker
./worker $ORIGINAL_VIDEO_CONTAINER $ORIGINAL_VIDEO_NAME $FORMAT_TO_ENCODE
