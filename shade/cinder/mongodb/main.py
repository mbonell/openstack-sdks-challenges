from shade import *
import os
from datetime import datetime

simple_logging(debug=True)
#conn = openstack_cloud(cloud='myfavoriteopenstack')
conn = openstack_cloud(cloud='osic-hackathon')

image_id = os.getenv('MONGODB_IMAGE')
flavor_id = os.getenv('MONGODB_FLAVOR')
private_network_id = os.getenv('MONGODB_NETWORK')

mongodb_database = os.getenv('MONGODB_DATABASE')
mongodb_username = os.getenv('MONGODB_USER')
mongodb_password = os.getenv('MONGODB_PASSWORD')

if image_id == None or flavor_id == None or private_network_id == None:
    print('[Error] Please specify the ID of the following values: Image, flavor and network.')
    quit()

if mongodb_database == None or mongodb_username == None or mongodb_password == None:
    print('[Warning] Please specify the values to configure your database connection string: Database name, username and password.')
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
    conn.create_security_group_rule(sec_group_name, 27017, 27017, 'TCP')

ex_userdata = '''#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/shade/cinder/mongodb/init.sh | bash -s --
'''

# Creating the database server
database_server = conn.create_server(wait=True, auto_ip=False,
    name = "mongodb-" + str(datetime.now()),
    image = image_id,
    flavor = flavor_id,
    network = private_network_id,
    key_name = 'marcela-ws', # DELETE
    security_groups = [sec_group_name],
    userdata = ex_userdata)

# Adding persistent storage to the database server
conn.attach_volume(database_server, database_volume, '/dev/vdb')

if conn.get_server_public_ip(database_server):
    ip_service = conn.get_server_public_ip(database_server)
else:
    ip_service = conn.get_server_private_ip(database_server)

print("MongoDB service ready!")
print("Connection string: MONGO_URL='mongodb://" + mongodb_username + ":" + mongodb_password + "@" + ip_service + ":27017/" + mongodb_database + "'")
