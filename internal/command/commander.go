package command

type Menu struct {
	Commands CommandMap
}

func (m *Menu) ActCommand(name Name) error {
	return m.Commands[name].Action()
}
