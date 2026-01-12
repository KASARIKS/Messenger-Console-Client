package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	if err := os.WriteFile("token.json", []byte((*gottenResult)["token"]), 0644); err != nil {
		return err
	}

	return nil
}

func DeleteUserAction() error {
	token, err := os.ReadFile("token.json")
	if err != nil {
		return err
	}

	if _, err = utils.SendDeleteRequestWithJWT("/delete", string(token)); err != nil {
		return err
	}

	return nil
}

func SendMessage() error {
	var payload types.MessagePayload

	token, err := os.ReadFile("token.json")
	if err != nil {
		return err
	}

	if err := utils.InputMessagePayload(&payload); err != nil {
		return err
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error with marshaling data: %v", err)
	}

	if _, err = utils.SendPostRequestWithJWT("/sendmessage", data, string(token)); err != nil {
		return err
	}

	return nil
}

func GetMessages() error {
	token, err := os.ReadFile("token.json")
	if err != nil {
		return err
	}

	gottenResult, err := utils.SendGetRequestWithJWT("/getgottenmessages", "0", string(token))
	if err != nil {
		return err
	}

	fmt.Println(gottenResult)

	for i := 1; ; i++ {
		p := strconv.Itoa(i)
		gottenResult, err := utils.SendGetRequestWithJWT("/getmessages", p, string(token))
		if err != nil {
			return err
		}
		if len(gottenResult) <= 0 {
			break
		}
	}

	return nil
}
