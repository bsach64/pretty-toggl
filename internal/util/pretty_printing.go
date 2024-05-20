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

func PrintKeyValue(key string, values ...string) {
	printKey := color.New(color.FgBlue).Add(color.Bold)
	printKey.Print("\t", key, ": ")
	for i, v := range values {
		fmt.Print(v)
		if i != len(values) - 1 {
			fmt.Print(" ,")
		}
	}
	fmt.Println()
}
