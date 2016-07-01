#4: Choosing images and flavors
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
