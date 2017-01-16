# My Pet's Photo Album: Sample App
The application uses Swift API to upload pictures in an OpenStack cloud. The web photo album displays those images from the cloud. The code sample uses [Shade](http://docs.openstack.org/infra/shade/) (Python) as SDK.

## App flow
1. Photos uploader
  *  Create a public storage container called "my-pets".
  *  Select the images to upload by specifying their location in your system.

## Cloud services used
* Object Storage (Swift)

## Usage

### Pre-requisites
* Images with cloud-init enabled.
* Python and Shade installed on your development system.

### Setting the cloud profile to use (clouds.yaml)
```
myfavoriteopenstack:
  auth:
    auth_url: https://example.com:5000/v3
    username: admin
    password: admin
    project_name: default
    domain_id: default
  region_name: RegionOne
```

### Running the uploader script
```
$ python uploader/init.py
```

Main steps:
* [Uploader] Create public containers.
* [Uploader] Upload images to the containers.
* [Backend] Get the container public URL with the list of images available in the container (XML format).
* [Backend] Generate the image’s public URLs from the XML.
* [Frontend] Show the images in the photo album frontend.
