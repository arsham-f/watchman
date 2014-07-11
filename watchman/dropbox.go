package main

import (
	. "github.com/scottferg/Dropbox-Go/dropbox"
	"os"
)

var (
	session Session
)

/*
	Connect to Dropbox using OAuth and test the connection
*/
func ConnectDropbox() error {

	token := os.Getenv("DROPBOX_TOKEN")
	session = Session{
		Oauth2AccessToken: token,
		AccessType:        "app_folder",
	}

	_, err := GetDelta(session, nil)
	return err
}

/*
	Download file from app folder
*/
func Get(path string) (data []byte, err error) {
	data, _, err = GetFile(session, URI(path), nil)
	return
}

/*
	Delete file from app folder
*/
func Del(path string) {
	Delete(session, URI(path), nil)
}

/*
	All URIs should be rooted in the app folder
*/
func URI(path string) Uri {
	return Uri{
		Root: "auto",
		Path: path,
	}
}
