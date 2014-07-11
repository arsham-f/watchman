package main

import (
	. "github.com/simonz05/godis/redis"
	"os"
)

var (
	client *Client
	queue  = "captures"
)

/*
	Connect to redis and test connection
*/
func ConnectRedis() error {

	client = New("tcp:162.243.93.174:6379", 0, os.Getenv("REDIS_PWD"))
	_, err := client.Keys("*")

	return err

}

/*
	Return name of next image
*/
func PopFromQueue() string {

	for {
		ret, err := client.Blpop([]string{queue}, 0)
		if HandleError(err, "Popping from queue") {
			continue
		}

		return ret.Elems[1].Elem.String()
	}

}
