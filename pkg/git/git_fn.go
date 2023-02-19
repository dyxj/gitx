package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Stash() {
	runShell("git stash")
}

func Checkout(branchName string) {
	runShell("git checkout " + branchName)
}

func runShell(command string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	// TODO define type of shell
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
