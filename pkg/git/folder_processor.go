package git

import (
	"encoding/json"
	"fmt"
	"github.com/dyxj/gitx/pkg/message"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

type FolderProcessor struct {
	cmd               *cobra.Command
	args              []string
	projectsProcessed int
	ProcessFn         func(cmd *cobra.Command, args []string)
}

type gitxConfig struct {
	Paths []string `json:"paths"`
}

func loadGitxConfig(cwd string) (*gitxConfig, error) {
	data, err := os.ReadFile(filepath.Join(cwd, "gitx.json"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var cfg gitxConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (f *FolderProcessor) ProcessFolder(cmd *cobra.Command, args []string) {
	f.cmd = cmd
	f.args = args

	originalWd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	cfg, err := loadGitxConfig(originalWd)
	if err != nil {
		log.Fatalf("error reading gitx.json: %v\n", err)
	}

	if cfg != nil {
		for _, p := range cfg.Paths {
			absPath := p
			if !filepath.IsAbs(p) {
				absPath = filepath.Join(originalWd, p)
			}
			if !IsGit(absPath) {
				fmt.Printf("warning: %s is not a git repository, skipping\n", absPath)
				continue
			}
			f.processProject(filepath.Base(absPath), absPath, originalWd)
			f.projectsProcessed++
		}
	} else {
		entries, err := os.ReadDir("./")
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			absPath := filepath.Join(originalWd, entry.Name())
			if !IsGit(absPath) {
				continue
			}
			f.processProject(entry.Name(), absPath, originalWd)
			f.projectsProcessed++
		}
	}

	if f.projectsProcessed < 1 {
		fmt.Println("no git projects to process")
	}
}

func (f *FolderProcessor) processProject(name string, absPath string, originalWd string) {

	defer func() {
		fmt.Println(message.Divider)
		err := os.Chdir(originalWd)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

	fmt.Println("Project: " + name)
	err := os.Chdir(absPath)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	f.ProcessFn(f.cmd, f.args)
}
