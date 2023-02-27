/*
Copyright © 2023 Darren Yim <darrenyxj@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/dyxj/gitx/pkg/git"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the Checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout [arg:branch name]",
	Args:  cobra.ExactArgs(1),
	Short: "Stash current branch and checkout specified branch for first level of directories in current folder",
	Long: `Performs the following actions to all directories in the current folder
1. stashes current work
2. checks out branch specified`,
	Run: runCheckout,
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func runCheckout(cmd *cobra.Command, args []string) {
	processor := git.FolderProcessor{
		ProcessFn: func(cmd *cobra.Command, args []string) {
			git.Stash()
			git.Checkout(args[0])
		},
	}

	processor.ProcessFolder(cmd, args)
}
