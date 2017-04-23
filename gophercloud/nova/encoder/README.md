# Video Encoding: Sample App
The application uses Nova API to launch encoding workers that receive video objects and convert them into different formats. The encoder app architecture should have a manager microservice that manages the job requests and launch a job worker to perform the encoding task. The code sample uses [Gophercloud](https://github.com/gophercloud/gophercloud) (Go) as SDK and [ffmpeg](https://ffmpeg.org/) for transcoding the video files.

## App flow
1. Creating a worker
   1. Select the image, flavor, network and security groups for the worker instance.
   1. Set as env vars the format to encode, the container and video name where the video to encode is stored in the cloud.
   1. Specify the bash script for the cloud-init service that install the worker dependencies (golang, ffmpeg), download the worker binary and run it once the instance is ready.
   1. Through the compute API, launch the worker instance specifying the infrastructure values and user data script.
1. Worker initialization
   1. Once the instance is ready, the init script will update the dependencies, install ffmpeg, set the cloud credentials as environment variables and download and run the encoding app.
1. Worker execution
   1. The encoding app installed in the worker will receive the original video file as a Swift object (container and object name) and the format to encode the video (MP4, MPEG, WEBM). Then the worker will download the original video from the cloud, execute the encoding task and at the end it will upload the new encoded video to the cloud through the object storage API.

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

### Create a worker to encode videos
```
$ go run create_worker.go
```
