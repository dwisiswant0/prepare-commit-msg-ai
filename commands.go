package main

import (
	"log"
)

func doInstall() {
	if isHookExists() {
		log.Println("There is prepare-commit-msg hook installed. Aborting.")
		return
	}

	log.Println("Installing prepare-commit-msg hook...")
	if err := writeHookFile(); err != nil {
		log.Fatal(err)
	}

	println("OK!")
}

func doUninstall() {
	if isHookExists() {
		if isOwnedHook() {
			log.Println("Installing prepare-commit-msg hook...")
			if err := removeHookFile(); err != nil {
				log.Fatal(err)
			}

			println("OK!")
		} else {
			log.Println("There is not prepare-commit-msg AI hook. Aborting.")
		}
	} else {
		log.Println("There is no prepare-commit-msg AI hook. Aborting.")
	}
}

func doEdit() {
	if isHookExists() || isOwnedHook() {
		log.Println("Editing prepare-commit-msg AI hook file...")
		if err := editHookFile(); err != nil {
			log.Fatal(err)
		}

		println("OK!")
		return
	}

	log.Println("There is no prepare-commit-msg AI hook. Aborting.")
}
