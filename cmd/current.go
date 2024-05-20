package cmd

import (
	"errors"
	"strconv"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/spf13/cobra"
)

type TimeEntryInfo struct {
	startTime time.Time
	project string
	description string
	tags []string
	billable bool
}


func init() {
	rootCmd.AddCommand(currentCmd)
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Displays the current running timer",
	Long:  "Displays the current running timer",
	Run:   current,
}

func current(cmd *cobra.Command, args []string) {
	client := togglapi.NewClient(time.Minute)
	ct, err := client.CurrentTimeEntryReq()
	var tEntryInfo TimeEntryInfo
	if err != nil {
		util.PrintError(err.Error())
		return
	}
	tEntryInfo.startTime = ct.Start
	if ct.ProjectID != nil {
		pName, err := GetProjectNameFromID(client, *ct.ProjectID)
		if err != nil {
			util.PrintError("Could Not Get Project Name")
			return
		}
		tEntryInfo.project = pName
	}

	tEntryInfo.description = ct.Description
	tEntryInfo.tags = ct.Tags
	tEntryInfo.billable = ct.Billable
	PrintTimeEntryInfo(tEntryInfo)
}

func GetProjectNameFromID(client togglapi.Client, id int) (string, error) {
	me, err := client.MeReq()
	if err != nil {
		return "", err
	}
	for i := 0; i < len(me.Projects); i++ {
		if id == me.Projects[i].ID {
			return me.Projects[i].Name, nil
		}
	}
	return "", errors.New("Could not Find Project..")
}


func PrintTimeEntryInfo(tEntryInfo TimeEntryInfo) {
	util.PrintKeyValue("Start Time", tEntryInfo.startTime.Local().String())
	if tEntryInfo.project != "" {
		util.PrintKeyValue("Project", tEntryInfo.project)
	}
	if len(tEntryInfo.tags) != 0 {
		util.PrintKeyValue("Tags", tEntryInfo.tags...)
	}
	if tEntryInfo.description != "" {
		util.PrintKeyValue("Description", tEntryInfo.description)
	}
	util.PrintKeyValue("Billable", strconv.FormatBool(tEntryInfo.billable))
}
