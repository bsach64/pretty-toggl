package util

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func CreateEnv() error {
	_, err := os.Stat(".env")
	if errors.Is(err, os.ErrNotExist) {
		_, e := os.Create(".env")
		if e != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func WriteAuthToEnv(val string) error {
	authMap := map[string]string{
		"AUTH": val,
	}
	err := godotenv.Write(authMap, ".env")
	if err != nil {
		return err
	}
	return nil
}

func ReadAuthFromEnv() (string, bool) {
	return os.LookupEnv("AUTH")
}
