# gitx

A git helper tool to run git commands across multiple repositories at once.

## Install

```sh
go install github.com/dyxj/gitx@latest
```

## Usage
```sh
➜  ~ gitx
A git helper tool to run git commands on multiple folders.
Use the help flag on each command to see the implementation details.
Currently only supports zsh.

Usage:
  gitx [command]

Available Commands:
  checkout      Stash current branch and checkout specified branch for first level of directories in current folder
  completion    Generate the autocompletion script for the specified shell
  currentbranch Shows current branch on first level of directories in current folder
  help          Help about any command
  pull          Execute git pull on first level of directories in current folder
  push          Execute git push on first level of directories in current folder
  status        Execute git status on first level of directories in current folder

Flags:
  -h, --help   help for gitx

Use "gitx [command] --help" for more information about a command.
```

## Configuration

By default, `gitx` scans the current directory for first-level git subdirectories and runs the command against each one.

### gitx.json

To run commands against a specific set of repositories instead, create a `gitx.json` file in the directory where you run `gitx`:

```json
{
  "paths": [
    "/absolute/path/to/repo",
    "../relative/path/to/another/repo"
  ]
}
```

When `gitx.json` is present, only the listed paths are processed. Paths can be absolute or relative to the directory containing `gitx.json`. Any path that is not a valid git repository will print a warning and be skipped.
