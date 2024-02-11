package main

import (
	"fmt"
	"os"
)

func init() {
	printBanner()
}

func main() {
	if len(os.Args) < 2 {
		printHelps()
		os.Exit(2)
	}

	if !isGitProject() {
		fmt.Println(errNotGitProject)
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "install", "isntall", "i":
		doInstall()
	case "uninstall", "unisntall", "u":
		doUninstall()
	case "edit", "ed", "e":
		doEdit()
	case "-h", "--h", "--help":
		printHelps()
		os.Exit(2)
	default:
		fmt.Printf(errInvalidCommand, cmd)
		os.Exit(1)
	}
}
