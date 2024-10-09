package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Execute() {
	args := os.Args
	if len(args) == 1 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error reading cwd: ", err)
			os.Exit(1)
		}
		formatted, err := formatDirectory(cwd)
		if err != nil { // This should never happen because the cwd should always resolve correctly.
			fmt.Printf("go-ls: cannot access '%s': %s\n", cwd, err)
			os.Exit(1)
		}
		fmt.Print(formatted)
	} else { // If there are arguments passed in.
		// Make list of directories
		d := args[1:]
		for _, p := range d {
			formatted, err := formatDirectory(p)
			if err != nil {
				fmt.Printf("go-ls: cannot access '%s': No such file or directory\n", p)
				continue
			}
			fmt.Print(formatted)
		}
	}

	// fmt.Println("Current dir contents:", files)
	os.Exit(0)
}

func formatDirectory(directory string) (dir string, err error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for _, file := range files {
		if file.IsDir() {
			sb.WriteString(file.Name())
			sb.WriteString("/")
		} else {
			sb.WriteString(file.Name())
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}
