package cmd

import (
	"time"

	"github.com/bsach64/pretty-toggl/internal/togglapi"
	"github.com/bsach64/pretty-toggl/internal/util"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new project or tag!",
		Long:  "Create a new project or tag!",
		Run:   create,
	}
	details struct {
		projectName string
		color       string
	}
)

var colors = map[string]string{
	lipgloss.NewStyle().Foreground(lipgloss.Color("#525266")).Render("Grey"):        "#525266",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#0b83d9")).Render("Blue"):        "#0b83d9",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#9e5bd9")).Render("Purple"):      "#9e5bd9",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#d94182")).Render("Pink"):        "#d94182",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#d94182")).Render("Orange"):      "#d94182",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#bf7000")).Render("Brown"):       "#bf7000",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#2da608")).Render("Green"):       "#2da608",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#06a893")).Render("Teal"):        "#06a893",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#c9806b")).Render("Light Brown"): "#c9806b",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#465bb3")).Render("Dark Blue"):   "#465bb3",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#990099")).Render("Dark Purple"): "#990099",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#c7af14")).Render("Yellow"):      "#c7af14",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#566614")).Render("Olive Green"): "#566614",
	lipgloss.NewStyle().Foreground(lipgloss.Color("#d92b2b")).Render("Red"):         "#d92b2b",
}

func create(cmd *cobra.Command, args []string) {
	c := togglapi.NewClient(time.Minute)
	if details.projectName != "" && details.color != "" {
		err := c.CreateProject(details.projectName, details.color)
		if err != nil {
			util.ErrorPrinter().Println(err)
			return
		}
		util.SuccessPrinter().Printf("Created %v with color %v\n", details.projectName, details.color)
		return
	}
	err := getDetails()
	if err != nil {
		util.ErrorPrinter().Println(err)
	}
	err = c.CreateProject(details.projectName, colors[details.color])
	if err != nil {
		util.ErrorPrinter().Println(err)
		return
	}
	util.SuccessPrinter().Printf("Created Project %v with color %v!\n", details.projectName, details.color)
}

func getDetails() error {
	projectColors := getColors()
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Name").Value(&details.projectName),
			huh.NewSelect[string]().Title("Color").Options(huh.NewOptions(projectColors...)...).Value(&details.color),
		),
	)
	err := form.Run()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&details.projectName, "project", "p", "", "Provide the name of the project")
	createCmd.Flags().StringVarP(&details.color, "color", "c", "", "Hex code for the color of the project (starting with #)")
}

func getColors() []string {
	projectColors := make([]string, 0)

	for k := range colors {
		projectColors = append(projectColors, k)
	}
	return projectColors
}
