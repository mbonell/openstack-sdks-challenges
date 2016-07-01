#4: Choosing images and flavors
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
