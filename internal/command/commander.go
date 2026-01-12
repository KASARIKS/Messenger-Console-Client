package command

import "fmt"

type Menu struct {
	Commands CommandMap
}

func (m *Menu) ActCommand(name Name) error {
	action, ok := m.Commands[name]
	if !ok {
		return fmt.Errorf("undefined action")
	}

	return action.Action()
}
