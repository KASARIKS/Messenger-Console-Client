package main

import (
	"fmt"
	"log"

	"github.com/kasariks/messenger_console_client/internal/actions"
	"github.com/kasariks/messenger_console_client/internal/command"
)

func main() {
	menu := command.Menu{
		Commands: command.CommandMap{
			"help": command.Command{
				Action: actions.HelpAction,
			},

			"register": command.Command{
				Action: actions.RegisterAction,
			},

			"login": command.Command{
				Action: actions.LoginAction,
			},
		},
	}

	var inputCommand string

	for inputCommand != "exit" {
		fmt.Print("Input a command: ")

		_, err := fmt.Scan(&inputCommand)
		if err != nil {
			log.Fatal()
		}

		if err := menu.ActCommand(command.Name(inputCommand)); err != nil {
			fmt.Println(err.Error())
		}
	}
}
