package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "Stops the current time entry!",
	Long: "Stops the current time entry!",
	Run: func(cmd *cobra.Command, args []string) {
		client := togglapi.NewClient(time.Minute)
		ct, err := client.CurrentTimeEntryReq()
		if err != nil {
			log.Fatal(err.Error())
		}
		suc, err := client.StopReq(ct.WorkspaceID, int(ct.ID))
		if !suc {
			fmt.Println("Could not stop timer!")
		} else {
			fmt.Println("Stopped Time Entry!")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
