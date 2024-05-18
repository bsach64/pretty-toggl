package cmd

import (
	"fmt"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:                   "auth [API TOKEN]",
	Short:                 "Authenticate with an API token!",
	Long:                  "Provide an API Token to enable functionality!",
	Run:                   auth,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
}

func auth(cmd *cobra.Command, args []string) {
	client := togglapi.NewClient(time.Minute)
	valid, err := client.AuthUsingToken(args[0])
	if err != nil {
		util.PrintError(err.Error())
		return
	}
	if valid {
		err := util.WriteAuthToken(args[0])
		if err != nil {
			util.PrintError(fmt.Sprintf("Could not save API Token! %v", err.Error()))
			return
		}
		fmt.Println("Successfully Authenticated!")
	} else {
		fmt.Println("Please Enter a valid API token!")
	}
}
