package command

type Command struct {
	Action func() error
}

type Name string
type CommandMap map[Name]Command
