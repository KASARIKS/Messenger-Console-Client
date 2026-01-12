package actions

import (
	"encoding/json"
	"fmt"

	"github.com/kasariks/messenger_console_client/internal/utils"
	"github.com/kasariks/messenger_console_client/types"
)

func HelpAction() error {
	fmt.Println("help, register, login, send, check")
	return nil
}

func RegisterAction() error {
	var payload types.RegisterPayload

	if err := utils.InputRegisterPayload(&payload); err != nil {
		return err
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error with marshaling data: %v", err)
	}

	if _, err := utils.SendPostRequest("/register", data); err != nil {
		return err
	}

	fmt.Println("The user has been registered.")

	return nil
}

func LoginAction() error {
	var payload types.LoginPayload

	if err := utils.InputLoginPayload(&payload); err != nil {
		return err
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error with marshaling data: %v", err)
	}

	gottenResult, err := utils.SendPostRequest("/login", data)

	if err != nil {
		return err
	}

	fmt.Println(*gottenResult)

	return nil
}
