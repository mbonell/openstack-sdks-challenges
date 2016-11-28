from shade import *

simple_logging(debug=True)
conn = openstack_cloud(cloud='myfavoriteopenstack')

# Create the containers for the photo album
container_name = 'my-pets'
container = conn.create_container(container_name, public=True)

# Upload images to the photo album containers
pets = {'birthday': 'my-pets/birthday.jpg', 'cute': 'my-pets/cute.jpg', 'tie': 'my-pets/tie.jpg'}
for object_name, file_path in pets.items():
	conn.create_object(container=container_name, name=object_name, filename=file_path)

print '\nListing photos in  the album %s' % container_name
print(conn.list_objects(container_name))
