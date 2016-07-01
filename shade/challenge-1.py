#1: Connecting to the cloud
from shade import *

simple_logging(debug=True)
conn = openstack_cloud(cloud='myfavoriteopenstack')