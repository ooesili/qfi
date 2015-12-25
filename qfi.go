package main

import (
	"fmt"
	"os"

	"github.com/ooesili/qfi/dispatch"
)

const usage = `Usage:
  qfi <target>
    Edit/chdir to a target
  qfi -a <target> <filename>
    Add a new target
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
`

func main() {
	if err := realMain(); err != nil {
		fmt.Printf("qfi: %s\n", err)

		// print usage info is this is a UsageError
		if _, ok := err.(dispatch.UsageError); ok {
			fmt.Print(usage)
		}

		os.Exit(1)
	}
}

func realMain() error {
	// register commands
	dispatcher := dispatch.New()
	dispatcher.Register("-a", notImplemented)
	dispatcher.Register("-m", notImplemented)
	dispatcher.Register("-d", notImplemented)
	dispatcher.Register("-r", notImplemented)
	dispatcher.Register("-l", notImplemented)
	dispatcher.Register("-s", notImplemented)
	dispatcher.RegisterFallback(dispatch.CommandFunc(func(args []string) error {
		// make sure there is exactly one argument
		if len(args) < 1 {
			return dispatch.UsageError("no targets given")
		}
		if len(args) > 1 {
			return dispatch.UsageError("too many arguments")
		}

		return fmt.Errorf("target '%s' does not exist", args[0])
	}))

	return dispatcher.Run(os.Args[1:])
}

var notImplemented = dispatch.CommandFunc(func(args []string) error {
	fmt.Println("not implemented")
	return nil
})
