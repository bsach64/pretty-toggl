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
	Use:   "auth [API TOKEN]",
	Short: "Authenticate with an API token!",
	Long:  "Provide an API Token to enable functionality!",
	Run: func(cmd *cobra.Command, args []string) {
		client := togglapi.NewClient(time.Minute)
		valid := client.AuthUsingToken(args[0])
		if valid {
			util.CreateEnv()
			err := util.WriteAuthToEnv(args[0])
			if err != nil {
				fmt.Println("Could not save API Token!")
				return
			}
			fmt.Println("Successfully Authenticated!")
		} else {
			fmt.Println("Please Enter a valid API token!")
		}
	},
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
}
