#6: Launching a web instance
from shade import *

simple_logging(debug=True)
conn = openstack_cloud(cloud='myfavoriteopenstack')
 
print "Selected image:"       
image_id = '3c76334f-9644-4666-ac3c-fa090f175655'
image = conn.get_image(image_id)
print(image)

print "\nSelected flavor:"
flavor_id = 'A1.1'
flavor = conn.get_flavor(flavor_id)
print(flavor)

ex_userdata = '''#!/usr/bin/env bash
curl -L -s https://raw.githubusercontent.com/MBonell/openstack-sdks-challenges/master/shade/init.sh | bash -s --
'''

print "\nServer creation:"
instance_name = 'mbonell-001'
testing_instance = conn.create_server(wait=True, auto_ip=True,
    name=instance_name,
    image=image_id,
    flavor=flavor_id,
    userdata=ex_userdata)

