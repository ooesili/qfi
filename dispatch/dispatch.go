// Package dispatch provides a simple command dispatcher, intended to be used
// to create CLI applications.
package dispatch

var (
	ErrNotFound = UsageError{"command not found"}
	ErrNoArgs   = UsageError{"no arguments given"}
)

// UsageError indicates invalid command usage.
type UsageError struct {
	Msg string
}

// Error implements the error interface.
func (err UsageError) Error() string {
	return err.Msg
}

// New creates a new Dispatcher.
func New() *Dispatcher {
	return &Dispatcher{
		commands: make(map[string]Command),
	}
}

// Dispatcher is the command dispatcher. It implements the Command interface,
// allowing arbitrarily nested subcommands.
type Dispatcher struct {
	commands       map[string]Command
	defaultCommand Command
}

// Register saves a Command that will be run when the given arg is seen by the
// dispatcher.
func (d *Dispatcher) Register(arg string, command Command) {
	d.commands[arg] = command
}

// RegisterFallback saves a Command that will be run when Run cannot find a
// registered Command, or when no args were given to Run.
func (d *Dispatcher) RegisterFallback(command Command) {
	d.defaultCommand = command
}

// Run looks at the first argument in args, and tries to run the Command that
// was registered with that argument. If it finds one, it will be run with all
// given args after the first one. If no Command was found or no args were
// given, a ErrNotFound or ErrNoArgs will be returned (respectively), unless a
// default command was registered, in which case it will be run with all given
// args.
func (d *Dispatcher) Run(args []string) error {
	// make sure at least one arg was given
	if len(args) < 1 {
		if d.defaultCommand != nil {
			return d.defaultCommand.Run(args)
		}
		return ErrNoArgs
	}

	// try to find a registered command that matches
	if command, ok := d.commands[args[0]]; ok {
		return command.Run(args[1:])
	}

	// try to use the default command
	if d.defaultCommand != nil {
		return d.defaultCommand.Run(args)
	}

	// nothing worked
	return ErrNotFound
}

// Command is a subcommand of the main application.
type Command interface {
	Run(args []string) error
}

// CommandFunc is an adapter to allow plain functions to be used as Commands.
type CommandFunc func(args []string) error

// Run calls f(args).
func (f CommandFunc) Run(args []string) error {
	return f(args)
}
