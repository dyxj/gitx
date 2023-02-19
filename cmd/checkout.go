/*
Copyright © 2023 Darren Yim <darrenyxj@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitx/pkg/git"
	"log"
	"os"
	"path/filepath"
)

// checkoutCmd represents the Checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout [arg:branch name]",
	Args:  cobra.ExactArgs(1),
	Short: "Checkout specified branch from first level of directories",
	Long: `Performs the following actions to all directories in the current folder
1. stashes current work
2. checks out branch specified`,
	Run: runCheckout,
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func runCheckout(cmd *cobra.Command, args []string) {
	originalWd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		processFolder(entry, originalWd, func() {
			git.Stash()
			git.Checkout(args[0])
			fmt.Println(divider)
		})
	}
}

func processFolder(entry os.DirEntry, originalWd string, processFn func()) {
	defer func() {
		err := os.Chdir(originalWd)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	fmt.Println("Project: " + entry.Name())
	entryPath := filepath.Join(originalWd, entry.Name())
	err := os.Chdir(entryPath)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	processFn()
}
