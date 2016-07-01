#5: Launching/Destroying simple instances
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

ex_userdata = '''#!/usr/bin/env bash

curl -L -s https://git.openstack.org/cgit/openstack/faafo/plain/contrib/install.sh | bash -s -- \
-i faafo -i messaging -r api -r worker -r demo
'''

print "\nServer creation:"
instance_name = 'mbonell-001'
testing_instance = conn.create_server(wait=True, auto_ip=True,
    name=instance_name,
    image=image_id,
    flavor=flavor_id,
    userdata=ex_userdata)

