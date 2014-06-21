require 'rubygems'
require 'redis'
require 'dropbox_sdk'

r = Redis.new(:password => ENV['REDIS_PWD'])
client = DropboxClient.new(ENV['DB_CLIENT_TOKEN'])

lastFile = nil
currentFile = nil


def compare(first, second) 
	return if first == nil || second == nil
	out = `compare -metric FUZZ #{first} #{second} out 2>&1`
	out = out.split(" ")[0].to_i
	puts out
	return out
end

while true do
	File.delete(lastFile) unless lastFile == nil

	lastFile = currentFile
	currentFile = r.blpop("captures")[1]

	img = client.get_file(currentFile)
	File.write(currentFile, img)
	

	
	client.file_delete(currentFile) if compare(lastFile, currentFile).to_i < 2500
end