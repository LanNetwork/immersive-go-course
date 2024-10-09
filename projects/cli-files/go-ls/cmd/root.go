package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func Execute() {
	args := os.Args
	if slices.Contains(args, "-h") {
		hMessage :=
			`go-ls is a command that performs the same function as the unix ls command.
		By default, it prints a list of directories (marked by a trailing /), and files (with their extension).
		go-ls can take any number of arguments, which can be file or directory paths.
		go-ls -h prints this message and ignores other arguments.
		`
		fmt.Println(hMessage)
		os.Exit(0)
	}
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

			if len(d) > 1 { // If d is longer than 1, place labels of each input unless p is a file
				if !isFile(p) {
					fmt.Printf("%s:\n", p)
				}
			}
			fmt.Print(formatted)
		}
	}
	os.Exit(0)
}

func formatDirectory(directory string) (dir string, err error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		// Try to see if it's a file, and if so return the name
		if isFile(directory) {
			return directory + "\n\n", nil
		}
		// Since directory is not a file (see check above), and it's not a directory (ReadDir failed), return error
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
	sb.WriteString("\n")
	return sb.String(), nil
}

// isFile function will only return true if the filePath leads to a file.
// If it does not exist, or is a directory it will return false.
func isFile(filePath string) bool {
	fileInfo, statErr := os.Stat(filePath)
	if statErr != nil {
		return false
	}

	return !fileInfo.IsDir()
}
