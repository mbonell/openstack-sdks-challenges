# MongoDB database service
The service uses Cinder API to add persistent storage and Nova API to launch a MongoDB server ready to use for your application. The code sample uses [Shade](http://docs.openstack.org/infra/shade/) (Python) as SDK.

## Service flow
1. Creating the Cinder volume
   1. The persistent volume should be created as first step.
1. Creating a MongoDB server
   1. Select the image, flavor, network and security groups for the server instance.
   1. Set as env vars the name of your database, user and password.
   1. Through the compute API, launch the MongoDB instance specifying the infrastructure values and user data script.
1. Adding persistent storage to the MongoDB server
   1. Once the database server is created, the Cinder volume will be attached to it and then the init script will update the dependencies, install MongoDB, create a filesystem the volume, mount it in the MongoDB data files location and create the database with its user and password.

## Cloud services used
* Block Storage (Cinder)
* Compute (Nova)

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

### Set the infrastructure values for the database server (flavor, image and network IDs)
```
export MONGODB_FLAVOR=3
export MONGODB_IMAGE=41ba40fd-e801-4639-a842-e3a2e5a2ebdd
export MONGODB_NETWORK=7004a83a-13d3-4dcd-8cf5-52af1ace4cae
```

### Set values to configure the database connection string
```
export MONGODB_DATABASE=mongo_db
export MONGODB_USER=admin
export MONGODB_PASSWORD=secret
```

### Launching the MongoDB database service
```
$ python mongodb/main.py
```
At the end the script will provide you the MongoDB URL to use as connection string in your application.
```
MONGO_URL='mongodb://admin:secret@<SERVICE_IP>:27017/mongo_db'
```
