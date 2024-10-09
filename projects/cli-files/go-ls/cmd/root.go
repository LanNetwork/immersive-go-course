package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	args := os.Args
	if len(args) == 1 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error reading cwd: ", err)
			os.Exit(1)
		}
		printDirectory(cwd)
	}

	// fmt.Println("Current dir contents:", files)
	os.Exit(0)
}

func printDirectory(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error running os.ReadDir(): ", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s/\n", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}
