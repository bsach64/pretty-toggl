package util

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(str string) {
	printErr := color.New(color.FgRed).Add(color.Bold)
	printErr.Print("ERROR: ")
	fmt.Println(str)
}

func PrintDone(str string) {
	color.New(color.FgGreen).Add(color.Bold).Println(str)
}
