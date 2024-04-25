package cmd

import (
	"fmt"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the current time entry!",
	Long:  "Stops the current time entry!",
	Run: func(cmd *cobra.Command, args []string) {
		client := togglapi.NewClient(time.Minute)
		ct, err := client.CurrentTimeEntryReq()
		if err != nil {
			util.PrintError(err.Error())
			return
		}
		suc, err := client.StopReq(ct.WorkspaceID, int(ct.ID))
		if err != nil {
			util.PrintError(err.Error())
			return
		}
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
