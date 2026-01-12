package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kasariks/messenger_console_client/types"
)

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
	resp, err := http.Post("http://localhost:8080"+path,
		"application/json",
		bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error with sending the request: %v", err)
	}
	defer resp.Body.Close()

	gottenResult, err := getResult(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error has gotten from the server: %s", (*gottenResult)["error"])
	}

	return gottenResult, nil
}

func SendDeleteRequest(path string, token string) (*map[string]string, error) {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080"+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("error with sending the request: %v", err)
	}
	defer resp.Body.Close()

	gottenResult, err := getResult(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error has gotten from the server: %s", (*gottenResult)["error"])
	}

	return gottenResult, nil
}

func getResult(body io.ReadCloser) (*map[string]string, error) {
	var gottenResult map[string]string

	if err := json.NewDecoder(body).Decode(&gottenResult); err != nil {
		return nil, fmt.Errorf("error has gotten while decoding the response: %v", err)
	}

	return &gottenResult, nil
}
