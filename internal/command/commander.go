package command

type Menu struct {
	Commands CommandMap
}

func (m *Menu) ActCommand(name Name, flags ...Flag) error {
	return m.Commands[name].Action(flags...)
}
