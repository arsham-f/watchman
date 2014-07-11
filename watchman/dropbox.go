package main

import (
	. "github.com/scottferg/Dropbox-Go/dropbox"
	"os"
)

var (
	session Session
)

func ConnectDropbox() error {
	token := os.Getenv("DROPBOX_TOKEN")
	session = Session{
		Oauth2AccessToken: token,
		AccessType:        "app_folder",
	}

	_, err := GetDelta(session, nil)
	return err
}

func Get(path string) (data []byte, err error) {
	data, _, err = GetFile(session, URI(path), nil)
	return
}

func Del(path string) {
	Delete(session, URI(path), nil)
}

func URI(path string) Uri {
	return Uri{
		Root: "auto",
		Path: path,
	}
}
