package main

import (
	"fmt"
	"log"

	"github.com/kasariks/messenger_console_client/internal/command"
	"github.com/kasariks/messenger_console_client/types"
)

func main() {
	commander := command.Menu{
		Commands: command.CommandMap{
			"help": command.Command{
				Flags: []string{"-command"},
				Action: func(flags ...command.Flag) error {
					if len(flags) == 0 {
						fmt.Println("help, register, login, send, check")
						return nil
					}

					for _, f := range flags {
						switch f.Name {
						case "-command":
							fmt.Println(f.Value)
						default:
							return fmt.Errorf("unexpected flag %s", f.Name)
						}
					}

					return nil
				},
			},

			"register": command.Command{
				Flags: []string{},
				Action: func(flags ...command.Flag) error {
					var payload types.RegisterPayload

					fmt.Print("Input user id(string): ")

					// resp, err := http.Post("http://localhost:8080/register", "application/json", )

					return nil
				},
			},
		},
	}

	var command string

	for command != "exit" {
		fmt.Print("Input a command:")

		_, err := fmt.Scan(&command)
		if err != nil {
			log.Fatal()
		}
	}
}
