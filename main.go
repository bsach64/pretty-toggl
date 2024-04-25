package main

import (
	"github.com/bsach64/pretty-toggl/cmd"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/joho/godotenv"
)

func main() {
	err := util.CreateEnv()
	if err != nil {
		util.PrintError(err.Error())
		return
	}
	godotenv.Load(".env")
	cmd.Execute()
}
