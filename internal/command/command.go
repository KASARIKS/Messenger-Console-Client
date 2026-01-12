package command

type ActionFunc func() error

type Command struct {
	Action ActionFunc
}

type Name string
type CommandMap map[Name]Command
