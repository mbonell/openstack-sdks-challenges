#!/bin/sh

WORKER_BIN_URL=https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/gophercloud/nova/encoder/worker/bin/worker

while getopts "OS_AUTH_URL:OS_PROJECT_NAME:OS_USERNAME:OS_PASSWORD:OS_DOMAIN_ID:ORIGINAL_VIDEO_FILE:FORMAT_TO_ENCODE" ARG; do
    case $ARG in 
        OS_AUTH_URL)
            OS_AUTH_URL=$OPTARG
        ;;
        OS_PROJECT_NAME)
            OS_PROJECT_NAME=$OPTARG
        ;;
        OS_USERNAME)
            OS_USERNAME=$OPTARG
        ;;
        OS_PASSWORD)
            OS_PASSWORD=$OPTARG
        ;;
        OS_DOMAIN_ID)
            OS_DOMAIN_ID=$OPTARG
        ;;
        ORIGINAL_VIDEO_FILE)
            ORIGINAL_VIDEO_FILE=$OPTARG
        ;;
        FORMAT_TO_ENCODE)
            FORMAT_TO_ENCODE=$OPTARG
        ;;
    esac
done

sudo add-apt-repository ppa:mc3man/trusty-media -y
sudo apt-get update
sudo apt-get install ffmpeg -y

# Setting cloud environment variables
export OS_AUTH_URL=$OS_AUTH_URL
export OS_PROJECT_NAME=$OS_PROJECT_NAME
export OS_USERNAME=$OS_USERNAME
export OS_PASSWORD=$OS_PASSWORD
export OS_DOMAIN_ID=$OS_DOMAIN_ID

printenv

# Downloading the worker binary and running it
wget $WORKER_BIN_URL
sudo chmod +x worker
./worker $ORIGINAL_VIDEO_FILE $FORMAT_TO_ENCODE
