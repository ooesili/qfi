package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ooesili/qfi/commands"
	"github.com/ooesili/qfi/config"
	"github.com/ooesili/qfi/detect"
	"github.com/ooesili/qfi/dispatch"
	"github.com/ooesili/qfi/edit"
	"github.com/ooesili/qfi/scripts"
	"github.com/ooesili/qfi/summarize"
)

var version = `¯\(°_o)/¯`

const usage = `Usage:
  qfi <target>
    Edit/chdir to a target
  qfi -a [<target>] <filename>
    Add a new target, name defaults to basename of file
  qfi -m <target> <filename>
    Move a target's destination
  qfi -d <target> [<target> [...]]
    Delete one or more targets
  qfi -r <target> <newname>
    Rename a target
  qfi -l [<target>]
    List all targets or resolve a specific target to its destination
  qfi -s
    Show detailed information on all targets
  qfi --shell (zsh|fish|bash) (comp|wrapper)
    Print completion or wrapper script for a shell
  qfi --version
    Display the version
  qfi --help
    Display this help message
`

func main() {
	if err := realMain(); err != nil {
		if err == edit.ErrWrapperShouldChdir {
			os.Exit(2)
		}
		fmt.Printf("qfi: %s\n", err)
		if _, ok := err.(dispatch.UsageError); ok {
			fmt.Print(usage)
		}
		os.Exit(1)
	}
}

func realMain() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}
	editorCmd := getEditor()

	cfg, err := config.New(configDir)
	if err != nil {
		return err
	}

	// create dependencies
	detector := detect.Detector{}
	summarizer := summarize.Summarizer{detector, cfg}
	editor := edit.Editor{editorCmd, detector, edit.RealExecutor{}, cfg}
	scriptGetter := scripts.Scripts{}

	// register commands
	dispatcher := dispatch.New()
	dispatcher.Register("-a", &commands.Add{cfg})
	dispatcher.Register("-m", &commands.Move{cfg})
	dispatcher.Register("-d", &commands.Delete{cfg})
	dispatcher.Register("-r", &commands.Rename{cfg})
	dispatcher.Register("-l", &commands.List{cfg, os.Stdout})
	dispatcher.Register("-s", &commands.Summary{summarizer, os.Stdout})
	dispatcher.Register("--shell", &commands.Shell{scriptGetter, os.Stdout})
	dispatcher.Register("--version", &commands.Print{version + "\n", os.Stdout})
	dispatcher.Register("--help", &commands.Print{usage, os.Stdout})
	dispatcher.RegisterFallback(&commands.Edit{editor})

	return dispatcher.Run(os.Args[1:])
}

func getConfigDir() (string, error) {
	if qfiHome := os.Getenv("QFI_CONFIGDIR"); qfiHome != "" {
		return qfiHome, nil
	}
	if home := os.Getenv("HOME"); home != "" {
		return filepath.Join(home, ".config", "qfi"), nil
	}
	return "", errors.New("neither $QFI_CONFIGDIR nor $HOME are set")
}

func getEditor() string {
	if envEditor := os.Getenv("EDITOR"); envEditor != "" {
		return envEditor
	}
	return "vi"
}
