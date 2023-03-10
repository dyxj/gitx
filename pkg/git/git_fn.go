package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func IsGit(path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("failed to read directory: %v err: %v\n", path, err)
		return false
	}

	for _, entry := range entries {
		if entry.Name() == ".git" {
			return true
		}
	}
	return false
}

func Stash() {
	runShell("git stash")
}

func Checkout(branchName string) {
	runShell("git checkout " + branchName)
}

func Pull() {
	runShell("git pull")
}

func Status() {
	runShell("git status")
}

func ShowCurrentBranch() {
	runShell("git branch --show-current")
}

func runShell(command string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	// should expose configuration to select type of shell
	eCmd := exec.Command("zsh", "-c", command)
	eCmd.Stdout = &stdout
	eCmd.Stderr = &stderr
	err := eCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	if len(stdout.String()) > 0 {
		fmt.Print(stdout.String())
	}

	if len(stderr.String()) > 0 {
		fmt.Print(stderr.String())
	}
}
