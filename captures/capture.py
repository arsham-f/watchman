import dropbox
import time
import picamera
import thread
import os
import redis


# Initialize Dropbox
client = dropbox.client.DropboxClient(os.environ.get("DB_CLIENT_TOKEN"))

#Initialize Camera
c = picamera.PiCamera()
c.vflip = True
c.hflip = True
c.resolution = (640, 480)
c.led = False

#Initialize Redis
pwd = os.environ.get("REDIS_PWD")
print pwd
r = redis.StrictRedis(host='arsh.am', port=6379, db=0, password=pwd)

#Get to work
def Capture():
	name = `round(time.time())` + ".jpg"

	print "capture"
	c.capture(name)

	print "upload"
	f = open(name, "rb")
	client.put_file(name, f)
	os.remove(name)

	print "redis"
	r.rpush('captures', name)

while True:
	if r.llen('captures') < 10:
		Capture()
	time.sleep(1)