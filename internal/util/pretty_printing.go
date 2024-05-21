package util

import (
	"fmt"

	"github.com/fatih/color"
)

func ErrorPrinter() *color.Color {
	color.New(color.FgRed).Add(color.Bold).Print("\tERROR: ")
	printErr := color.New(color.FgRed)
	return printErr
}

func SuccessPrinter() *color.Color {
	return color.New(color.FgGreen).Add(color.Bold)
}

func PrintKeyValue(key string, values ...string) {
	HeadingPrinter().Print("\t", key, ": ")
	for i, v := range values {
		fmt.Print(v)
		if i != len(values) - 1 {
			fmt.Print(" ,")
		}
	}
	fmt.Println()
}

func HeadingPrinter() *color.Color {
	return color.New(color.FgBlue).Add(color.Bold)
}
