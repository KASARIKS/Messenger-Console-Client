package command

type Command struct {
	Flags  []string
	Action func(flags ...Flag) error
}

type Name string
type CommandMap map[Name]Command
