package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type input struct {
	projectName string
	tags        string
	billable    bool
	description string
}

var (
	in          input
	workspaceID int
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Starts a timer",
		Long:  "Starts a timer",
		Run:   start,
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&in.billable, "bill", "b", false, "Whether the time entry is billable or not")
	startCmd.Flags().StringVarP(&in.projectName, "project", "p", "", "Provide the name of the project")
	startCmd.Flags().StringVarP(&in.description, "description", "d", "", "Provide a description")
	startCmd.Flags().StringVarP(&in.tags, "tags", "t", "", "Provide tags")
}

func parseInput() error {
	if !in.billable && in.description == "" && in.projectName == "" && in.tags == "" {
		err := runForm()
		if err != nil {
			return err
		}
	}
	return nil
}

func start(cmd *cobra.Command, args []string) {
	err := parseInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tE := togglapi.CreateNewTimeEntry()
	tE.Billable = in.billable
	client := togglapi.NewClient(time.Minute)
	me, err := client.MeReq()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tE.WorkspaceID = me.DefaultWorkspaceID
	if in.projectName != "" {
		id, found := getProjectID(me, in.projectName)
		if !found {
			fmt.Println("Could not find project")
			return
		}
		tE.ProjectID = &id
	}
	if in.description != "" {
		tE.Description = &in.description
	}
	if in.tags != "" {
		ListOfTags := strings.Split(in.tags, ",")
		tE.Tags = ListOfTags
	}
	tE.Start = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	started, err := client.StartTimeEntry(tE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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

func runForm() error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Project Name").Value(&in.projectName),
			huh.NewInput().Title("Tags [Separated by Commas]").Value(&in.tags),
			huh.NewInput().Title("Description").Value(&in.description),
			huh.NewConfirm().Title("Billable").Value(&in.billable),
		),
	)
	err := form.Run()
	if err != nil {
		return err
	}
	return nil
}
