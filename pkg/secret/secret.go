package secret

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func ShouldAsk(cmd *cobra.Command, questionMessage string, flagKey string) {
	passwordPromptFlag := cmd.Flag(flagKey)
	passwordPromptFlag.Value.String()
	isPasswordPrompt, err := strconv.ParseBool(passwordPromptFlag.Value.String())
	if err != nil {
		log.Fatalf("dev error: %v\n", err)
	}
	if isPasswordPrompt {
		Ask(questionMessage)
	}
}

func Ask(questionMessage string) string {
	// Get the initial state of the terminal.
	initialTermState, err := terminal.GetState(syscall.Stdin)
	if err != nil {
		panic(err)
	}

	// Restore terminal state in the event of an interrupt.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	fmt.Print(questionMessage)
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	return string(p)
}
