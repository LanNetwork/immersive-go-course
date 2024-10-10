package cmd

import "os"

func Execute() {
	os.Stdout.WriteString("Hi, test from Execute")
}
