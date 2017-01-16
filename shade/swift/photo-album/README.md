# My Pet's Photo Album: Sample App
The application uses Swift API to upload pictures in an OpenStack cloud. The web photo album displays those images from the cloud. The code sample uses [Shade](http://docs.openstack.org/infra/shade/) (Python) as SDK.

![My Pet's Photo Album Example](https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/shade/swift/photo-album/my-pet-photo-album.png)

## App flow
1. Photos uploader
  *  Create a public storage container called "my-pets".
  *  Select the images to upload by specifying their location in your system.
1. Web photo album
  *  Once all your images are available from the cloud, you can get the container public URL with the list of images available in the container (XML format). E.g: [https://cloud1.osic.org:8080/v1/AUTH_e92735e996c44a758f12262d0501c79b/my-pets](https://cloud1.osic.org:8080/v1/AUTH_e92735e996c44a758f12262d0501c79b/my-pets)
  * Each image can be accesible by the URL provided by the Swift API (Swift endpoint + container name + image name). **Challenge:** Add a photo album's backend and try to automate the generation of these URL. Send me the PR! :wink:
  * In your web server publish the web app (html+css) with your photos URLs.

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
