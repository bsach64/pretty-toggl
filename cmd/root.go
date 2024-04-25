package cmd

import (
	"os"

	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pretty-toggl",
	Short: "A pretty CLI app for toggl track!",
	Long:  `A CLI app for toggl track!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		util.PrintError(err.Error())
		os.Exit(1)
	}
}

func init() {
}
