package git

import (
	"fmt"
	"github.com/dyxj/gitx/pkg/message"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

type FolderProcessor struct {
	cmd       *cobra.Command
	args      []string
	ProcessFn func(cmd *cobra.Command, args []string)
}

func (f *FolderProcessor) ProcessFolder(cmd *cobra.Command, args []string) {
	f.cmd = cmd
	f.args = args

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
		if !IsGit(filepath.Join(originalWd, entry.Name())) {
			continue
		}

		f.processProject(entry, originalWd)
	}
}

func (f *FolderProcessor) processProject(entry os.DirEntry, originalWd string) {

	defer func() {
		fmt.Println(message.Divider)
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

	f.ProcessFn(f.cmd, f.args)
}
