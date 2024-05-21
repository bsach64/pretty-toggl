package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var (
	in struct {
		projectName string
		tags        string
		billable    bool
		description string
	}
	workspaceID int
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "Starts a timer",
		Long:  "Starts a timer\nIf no flag is specified a form pops up asking for details",
		Run:   start,
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&in.billable, "bill", "b", false, "Whether the time entry is billable or not")
	startCmd.Flags().StringVarP(&in.projectName, "project", "p", "", "Provide the name of the project")
	startCmd.Flags().StringVarP(&in.description, "description", "d", "", "Provide a description")
	startCmd.Flags().StringVarP(&in.tags, "tags", "t", "", "Provide tags as a comma separated list")
}

func start(cmd *cobra.Command, args []string) {
	client := togglapi.NewClient(time.Minute)
	me, err := client.MeReq()
	if err != nil {
		util.ErrorPrinter().Println(err.Error())
		return
	}

	err = parseInput(me)
	if err != nil {
		util.ErrorPrinter().Println(err.Error())
		return
	}
	tE := togglapi.CreateNewTimeEntry()
	tE.Billable = in.billable
	tE.WorkspaceID = me.DefaultWorkspaceID
	if in.projectName != "" {
		id, found := getProjectID(me, in.projectName)
		if !found {
			util.ErrorPrinter().Println("Could Not Find Project Name!")
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
		util.ErrorPrinter().Println(err.Error())
		return
	}
	if started {
		util.SuccessPrinter().Println("Started a Time Entry!")
	} else {
		util.ErrorPrinter().Println("Could Not Start Time Entry!")
	}
}

func parseInput(me togglapi.Me) error {
	if !in.billable && in.description == "" && in.projectName == "" && in.tags == "" {
		err := runForm(me)
		if err != nil {
			return err
		}
	}
	return nil
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

func runForm(me togglapi.Me) error {
	projectNames := make([]string, 0)
	for _, project := range me.Projects {
		if project.Active {
			projectNames = append(projectNames, project.Name)
		}
	}
	tagNames := make([]string, 0)
	for _, tag := range me.Tags {
		tagNames = append(tagNames, tag.Name)
		fmt.Println(tag.At)
	}
	var tags []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Project").Options(huh.NewOptions(projectNames...)...).Value(&in.projectName),
			huh.NewMultiSelect[string]().Title("Tags").Options(huh.NewOptions(tagNames...)...).Value(&tags),
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
