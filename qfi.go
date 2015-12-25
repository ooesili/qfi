package main

import (
	"fmt"
	"os"

	"github.com/ooesili/qfi/dispatch"
)

var usage = `
Usage: qfi TARGET
       qfi [-a|-m] TARGET FILENAME
       qfi -d TARGET1 [TARGET2 [...]]
       qfi -r TARGET NEWNAME
       qfi -l [TARGET]
       qfi -s
`[1:]

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
