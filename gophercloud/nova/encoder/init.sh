#!/bin/sh
WORKER_BIN_URL=https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/worker/bin/worker

sudo add-apt-repository ppa:mc3man/trusty-media -y
sudo apt-get update
sudo apt-get install ffmpeg -y

# Setting cloud environment variables
export OS_AUTH_URL=$OS_AUTH_URL
export OS_PROJECT_NAME=$OS_PROJECT_NAME
export OS_USERNAME=$OS_USERNAME
export OS_PASSWORD=$OS_PASSWORD
export OS_DOMAIN_ID=$OS_DOMAIN_ID

# Downloading the worker binary and running it
wget $WORKER_BIN_URL
./worker $ORIGINAL_VIDEO_FILE $FORMAT_TO_ENCODE
