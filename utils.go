package main

import (
	"fmt"
	"os"
	"strings"

	"os/exec"
)

func printBanner() {
	fmt.Fprintf(
		os.Stderr,
		"\n  %s %s\n  --\n  made with ♥ by %s\n\n",
		banner, version, author,
	)
}

func getUsage() string {
	return fmt.Sprintf(usage, app)
}

func getExamples() string {
	return fmt.Sprintf(examples, app)
}

func printHelps() {
	fmt.Fprintf(os.Stderr, "%s\n\n%s\n\n%s\n", getUsage(), commands, getExamples())
}

func isGitProject() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(output)) == "true"
}

func isHookExists() bool {
	_, err := os.Stat(hookPath)

	return !os.IsNotExist(err)
}

func isOwnedHook() bool {
	hookScript, err := os.ReadFile(hookPath)
	if err != nil {
		return false
	}

	return strings.Contains(string(hookScript), hookHeader)
}

func writeHookFile() error {
	err := os.WriteFile(hookPath, prepareCommitMsgByte, 0644)
	if err != nil {
		return err
	}

	// Convert to binary
	cmd := exec.Command("chmod", "a+x", hookPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func removeHookFile() error {
	return os.Remove(hookPath)
}

func editHookFile() error {
	var args []string

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	} else {
		// NOTE(dwisiswat0): If you know, you know!
		args = strings.Split(editor, " ")
		if len(args) > 1 {
			editor = args[0]
		}
		args = args[1:]
	}

	args = append(args, hookPath)

	cmd := exec.Command(editor, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
