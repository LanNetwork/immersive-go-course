package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error reading cwd: ", err)
		return
	}

	files, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error running os.ReadDir(): ", err)
	}

	// fmt.Println("Current dir contents:", files)
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s/\n", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}
