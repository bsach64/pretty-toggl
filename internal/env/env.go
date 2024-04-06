package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateEnv() {
	_, err := os.Stat(".env")
	if errors.Is(err, os.ErrNotExist) {
		_, e := os.Create(".env") 
		if e != nil {
			log.Fatal("Could not create .env file")
		}
	} else if err != nil {
		log.Fatal("Could not stat .env")
	}
}

func WriteAuthToEnv(val string) (error) {
	authMap := map[string]string{
		"AUTH":val,
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
