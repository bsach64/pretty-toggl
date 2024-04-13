package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/spf13/cobra"
)

var (
	projectName string
	tags string
	billable bool
	description string
	workspaceID int
	startCmd = &cobra.Command{
		Use: "start",
		Short: "Starts a timer",
		Long: "Starts a timer",
		Run: Start,
	}
)


func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&billable, "bill", "b", false, "Whether the time entry is billable or not")
	startCmd.Flags().StringVarP(&projectName, "project", "p", "", "Provide the name of the project")
	startCmd.Flags().StringVarP(&description, "description", "d", "", "Provide a description")
	startCmd.Flags().StringVarP(&tags, "tags", "t", "", "Provide tags")
}

func Start(cmd *cobra.Command, args []string) {
	tE := togglapi.CreateNewTimeEntry()
	tE.Billable = billable
	client := togglapi.NewClient(time.Minute)
	me, err := client.MeReq()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tE.WorkspaceID = me.DefaultWorkspaceID
	if projectName != "" {
		id, found := getProjectID(me, projectName)
		if !found {
			fmt.Println("Could not find project")
			return
		}
		tE.ProjectID = &id
	}
	if description != "" {
		tE.Description = &description
	}
	if tags != "" {
		ListOfTags := strings.Split(tags, " ")
		tE.Tags = ListOfTags
	}
	tE.Start = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	fmt.Println(tE)
	started := client.StartTimeEntry(tE)
	if started {
		fmt.Println("Started a Time Entry!")
	} else {
		fmt.Println("Could not start Time Entry!")
	}
}

func getProjectID(me togglapi.Me, projectName string) (int, bool) {
	lower := strings.ToLower(projectName)
	for _, entry := range me.Projects {
		if strings.ToLower(entry.Name) == lower {
			return entry.ID, true
		}
	}
	return 0, false
}
