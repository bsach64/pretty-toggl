package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(currentCmd)
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Displays the current running timer",
	Long:  "Displays the current running timer",
	Run: func(cmd *cobra.Command, args []string) {
		var output strings.Builder
		client := togglapi.NewClient(time.Minute)
		ct, err := client.CurrentTimeEntryReq()
		if ct.ID == 0 {
			fmt.Println("No Current Running Time Entry!")
			return
		}

		if err != nil {
			log.Fatal(err.Error())
		}
		output.WriteString("Start Time: ")
		output.WriteString(fmt.Sprint(ct.Start) + "\n")
		if ct.ProjectID != nil {
			output.WriteString("Project: ")
			output.WriteString(GetProjectNameFromID(client, *ct.ProjectID))
			output.WriteString("\n")
		}

		if ct.Description != "" {
			output.WriteString("Description: " + ct.Description + "\n")
		}
		if len(ct.Tags) > 0 {
			output.WriteString("Tags: ")
		}
		for i := 0; i < len(ct.Tags); i++ {
			output.WriteString(ct.Tags[i])
			if i != len(ct.Tags)-1 {
				output.WriteString(", ")
			}
		}
		output.WriteString("\n")
		output.WriteString("Billable: " + strconv.FormatBool(ct.Billable) + "\n")
		fmt.Print(output.String())
	},
}

func GetProjectNameFromID(client togglapi.Client, id int) string {
	me, err := client.MeReq()
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := 0; i < len(me.Projects); i++ {
		if id == me.Projects[i].ID {
			return me.Projects[i].Name
		}
	}
	return ""
}
