package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kasariks/messenger_console_client/internal/command"
	"github.com/kasariks/messenger_console_client/types"
)

func main() {
	menu := command.Menu{
		Commands: command.CommandMap{
			"help": command.Command{
				Action: func() error {
					fmt.Println("help, register, login, send, check")
					return nil
				},
			},

			"register": command.Command{
				Action: func() error {
					var payload types.RegisterPayload

					if err := InputRegisterPayload(&payload); err != nil {
						return err
					}

					data, err := json.Marshal(payload)
					if err != nil {
						return fmt.Errorf("error with marshaling data: %v", err)
					}

					if _, err := SendPostRequest("/register", data); err != nil {
						return err
					}

					fmt.Println("The user has been registered.")

					return nil
				},
			},

			"login": command.Command{
				Action: func() error {
					var payload types.LoginPayload

					if err := InputLoginPayload(&payload); err != nil {
						return err
					}

					data, err := json.Marshal(payload)
					if err != nil {
						return fmt.Errorf("error with marshaling data: %v", err)
					}

					gottenResult, err := SendPostRequest("/login", data)

					if err != nil {
						return err
					}

					fmt.Println(*gottenResult)

					return nil
				},
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

func InputRegisterPayload(payload *types.RegisterPayload) error {
	fmt.Print("Input user id(string):")
	if _, err := fmt.Scan(&payload.Id); err != nil {
		return fmt.Errorf("error with inputing the id: %v", err)
	}

	fmt.Print("Input user password(string):")
	if _, err := fmt.Scan(&payload.Password); err != nil {
		return fmt.Errorf("error with inputing the password: %v", err)
	}

	fmt.Print("Input user nickname(string):")
	if _, err := fmt.Scan(&payload.Nickname); err != nil {
		return fmt.Errorf("error with inputing the nickname: %v", err)
	}

	return nil
}

func InputLoginPayload(payload *types.LoginPayload) error {
	fmt.Print("Input user id(string):")
	if _, err := fmt.Scan(&payload.Id); err != nil {
		return fmt.Errorf("error with inputing the id: %v", err)
	}

	fmt.Print("Input user password(string):")
	if _, err := fmt.Scan(&payload.Password); err != nil {
		return fmt.Errorf("error with inputing the password: %v", err)
	}

	return nil
}

func SendPostRequest(path string, data []byte) (*map[string]string, error) {
	var gottenResult map[string]string

	resp, err := http.Post("http://localhost:8080"+path,
		"application/json",
		bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error with sending the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&gottenResult); err != nil {
			return nil, fmt.Errorf("error has gotten while decoding the response: %v", err)
		}

		return nil, fmt.Errorf("error has gotten from the server: %s", gottenResult["error"])
	}

	if err := json.NewDecoder(resp.Body).Decode(&gottenResult); err != nil {
		return nil, fmt.Errorf("error has gotten while decoding the response: %v", err)
	}

	return &gottenResult, nil
}
