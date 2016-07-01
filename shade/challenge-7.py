#6: Deploying My FirstApp
from shade import *

simple_logging(debug=True)
conn = openstack_cloud(cloud='myfavoriteopenstack')
   
print "Selected image:"    
image_id = 'b55d48a9-29af-490c-af8d-ff897f688f0c'
image = conn.get_image(image_id)
print(image)


print "\nSelected flavor:"
flavor_id = '2'
flavor = conn.get_flavor(flavor_id)
print(flavor)


print('\nChecking for existing SSH keypair...')
keypair_name = 'demokey'
pub_key_file = '~/.ssh/demo_key.pub'

if conn.search_keypairs(keypair_name):
    print('Keypair already exists. Skipping import.')
else:
    print('Adding keypair...')
    conn.create_keypair(keypair_name, open(pub_key_file, 'r').read().strip())

for keypair in conn.list_keypairs():
    print(keypair)


print('\nChecking for existing security groups...')
sec_group_name = 'all-in-one'
if conn.search_security_groups(sec_group_name):
    print('Security group already exists. Skipping creation.')
else:
    print('Creating security group.')
    conn.create_security_group(sec_group_name, 'network access for all-in-one application.')
    conn.create_security_group_rule(sec_group_name, 80, 80, 'TCP')
    conn.create_security_group_rule(sec_group_name, 22, 22, 'TCP')

conn.search_security_groups(sec_group_name)


ex_userdata = '''#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/fog/init.sh | bash -s --
'''

print "\nCreating app instance:"
instance_name = 'all-in-one'
testing_instance = conn.create_server(wait=True, auto_ip=False,
    name=instance_name,
    image=image_id,
    flavor=flavor_id,
    key_name=keypair_name,
    security_groups=[sec_group_name],
    userdata=ex_userdata)


print "\nAssigning an public IP to the app instance :"
f_ip = conn.available_floating_ip()
conn.add_ip_list(testing_instance, [f_ip['floating_ip_address']])


print('\nThe Fractals app will be deployed to http://%s' % f_ip['floating_ip_address'] )
