package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/spf13/cobra"
)

var (
	projectsFlag   bool
	tagsFlag       bool
	workspacesFlag bool
	listCmd        = &cobra.Command{
		Use:   "list",
		Short: "Get list of all projects, tags, workspaces",
		Long:  "Get list of all projects, tags, workspaces",
		Run: func(cmd *cobra.Command, args []string) {
			client := togglapi.NewClient(time.Minute)
			me, err := client.MeReq()
			if err != nil {
				util.PrintError(err.Error())
				return
			}
			fmt.Print(List(me, projectsFlag, tagsFlag, workspacesFlag))
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&projectsFlag, "projects", "p", false, "Specify to list projects")
	listCmd.Flags().BoolVarP(&tagsFlag, "tags", "t", false, "Specify to list tags")
	listCmd.Flags().BoolVarP(&workspacesFlag, "workspaces", "w", false, "Specify to list workspaces")
}

func List(me togglapi.Me, projectsFlag, tagsFlag, workspacesFlag bool) string {
	var output strings.Builder
	if projectsFlag {
		output.WriteString("Projects\n")
		for i := 0; i < len(me.Projects); i++ {
			if me.Projects[i].Active {
				output.WriteString("\t")
				output.WriteString(me.Projects[i].Name)
				output.WriteString("\n")
			}
		}
	}

	if tagsFlag {
		output.WriteString("Tags\n")
		for i := 0; i < len(me.Tags); i++ {
			output.WriteString("\t")
			output.WriteString(me.Tags[i].Name)
			output.WriteString("\n")
		}
	}

	if workspacesFlag {
		output.WriteString("Workspaces\n")
		for i := 0; i < len(me.Workspaces); i++ {
			output.WriteString("\t")
			output.WriteString(me.Workspaces[i].Name)
			output.WriteString("\n")
		}
	}

	return output.String()
}
