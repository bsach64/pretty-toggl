package util

import (
	"errors"
	"os"
)

const PrettyTogglPath = "/.local/share/pretty-toggl"

func GetFilePath() (string, error) {
	var path *string
	XDGdataHome, found := os.LookupEnv("XDG_DATA_HOME")
	if !found {
		home, found := os.LookupEnv("HOME")
		if !found {
			return "", errors.New("Could not get XDG_DATA_HOME or HOME")
		}
		path = &home
	} else {
		path = &XDGdataHome
	}
	tokenFilePath := *path + PrettyTogglPath
	return tokenFilePath, nil
}

func WriteAuthToken(token string) error {
	path, err := GetFilePath()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(path + "/auth")
	if err != nil {
		return err
	}
	_, err = file.WriteString(token)
	if err != nil {
		return err
	}
	return nil
}

func ReadAuthToken() (string, error) {
	path, err := GetFilePath()
	if err != nil {
		return "", err
	}
	tokenFilePath := path + "/auth"
	authByte, err := os.ReadFile(tokenFilePath)
	if err != nil {
		if err == os.ErrNotExist {
			return "", errors.New("Please login before using!")
		}
		return "", err
	}
	return string(authByte), nil
}
