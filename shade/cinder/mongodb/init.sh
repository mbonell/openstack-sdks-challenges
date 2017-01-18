#!/bin/sh

# Create a filesystem in the empty volume
sudo mke2fs /dev/vdb

# Link the volume to the MongoDB data files
mkdir /var/lib/mongodb
sudo echo "/dev/vdb /var/lib/mongodb ext4 defaults  1 2" >> /etc/fstab
sudo mount /var/lib/mongodb

# Installing MongoDB
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.4.list
sudo apt-get update
sudo apt-get install mongodb-org -y

# Verify that MongoDB has started successfully
sudo service mongod start
sudo service mongod status
