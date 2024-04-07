package main

import (
	"github.com/bsach64/pretty-toggl/cmd"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/joho/godotenv"
)

func main() {
	util.CreateEnv()
	godotenv.Load(".env")
	cmd.Execute()
}
