package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func Execute() {
	args := os.Args[1:]
	if len(args) == 0 {
		readPrintStdio()
	} else {
		fmt.Println("IAN: You should put the regular cat logic here.")
	}
}

func readPrintStdio() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}
