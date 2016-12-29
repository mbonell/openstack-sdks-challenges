# Video Encoding: Sample App
The application uses Nova API to launch encoding workers that receive video objects and convert them into different formats. The encoder app architecture should have a manager microservice that manages the job requests and launch a job worker to perform the encoding task. The code sample uses Gophercloud (Go) as SDK and ffmpeg for transcoding the video files.

## App flow
1. TODO
1. TODO

## Cloud services used
* Compute (Nova)
* Object Storage (Swift)

## Usage

### Pre-requisites
* Images with cloud-init enabled.
* Golang and Gophercloud installed on your development system.

### Set the credentials of your cloud
```
export OS_AUTH_URL=https://example.com:5000/v3
export OS_USERNAME=admin
export OS_PASSWORD=admin
export OS_DOMAIN_ID=default
export OS_REGION_NAME=RegionOne
```

### Set the infrastructure values for the workers (flavor, image and network IDs)
```
export WORKER_SERVER_FLAVOR=3
export WORKER_SERVER_IMAGE=41ba40fd-e801-4639-a842-e3a2e5a2ebdd
export WORKER_SERVER_NETWORK=7004a83a-13d3-4dcd-8cf5-52af1ace4cae
```

### Set values for the video (container and name) object in your cloud and the format to encode it
```
export ORIGINAL_VIDEO_CONTAINER=original-videos
export ORIGINAL_VIDEO_NAME=prairie-dog.mov
export FORMAT_TO_ENCODE=webm
```
