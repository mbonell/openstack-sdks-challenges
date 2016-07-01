#!/bin/sh
sudo apt-get install apache2 apache2-doc apache2-mpm-prefork apache2-utils libexpat1 ssl-cert -y

wget https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/shade/shade-demo.html
sudo cp shade-demo.html /var/www/html/index.html
