from shade import *
import os
from datetime import datetime

simple_logging(debug=True)
#conn = openstack_cloud(cloud='myfavoriteopenstack')
conn = openstack_cloud(cloud='osic-hackathon')

image_id = os.getenv('MONGODB_IMAGE')
flavor_id = os.getenv('MONGODB_FLAVOR')
private_network_id = os.getenv('MONGODB_NETWORK')

if image_id == None or flavor_id == None or private_network_id == None:
    print('[Error] Please specify the ID of the following values: Image, flavor and external network.')
    quit()

# Creating a Cinder volume
database_volume = conn.create_volume(size=1, display_name='mongodb', description='MongoDB database service volume')

sec_group_name = 'mongodb'
if conn.search_security_groups(sec_group_name):
    print('Security group already exists. Skipping creation.')
else:
    print('Creating security group.')
    conn.create_security_group(sec_group_name, 'network access for mongodb service.')
    conn.create_security_group_rule(sec_group_name, 22, 22, 'TCP')

ex_userdata = '''#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/shade/cinder/mongodb/init.sh | bash -s --
'''

# Creating the database server
database_server = conn.create_server(wait=True, auto_ip=False,
    name = "mongodb-" + str(datetime.now()),
    image = image_id,
    flavor = flavor_id,
    network = private_network_id,
    security_groups = [sec_group_name],
    userdata = ex_userdata)

# Adding persistent storage to the database server
conn.attach_volume(database_server, database_volume, '/dev/vdb')

print('MongoDB service ready!')
