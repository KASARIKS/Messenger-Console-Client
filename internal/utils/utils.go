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
	fmt.Print("Input user id(string): ")
	if _, err := fmt.Scan(&payload.Id); err != nil {
		return fmt.Errorf("error with inputing the id: %v", err)
	}

	fmt.Print("Input user password: ")
	if _, err := fmt.Scan(&payload.Password); err != nil {
		return fmt.Errorf("error with inputing the password: %v", err)
	}

	fmt.Print("Input user nickname: ")
	if _, err := fmt.Scan(&payload.Nickname); err != nil {
		return fmt.Errorf("error with inputing the nickname: %v", err)
	}

	return nil
}

func InputLoginPayload(payload *types.LoginPayload) error {
	fmt.Print("Input user id(string): ")
	if _, err := fmt.Scan(&payload.Id); err != nil {
		return fmt.Errorf("error with inputing the id: %v", err)
	}

	fmt.Print("Input user password: ")
	if _, err := fmt.Scan(&payload.Password); err != nil {
		return fmt.Errorf("error with inputing the password: %v", err)
	}

	return nil
}

func InputMessagePayload(payload *types.MessagePayload) error {
	fmt.Print("Input recipient id(string): ")
	if _, err := fmt.Scan(&payload.RecipientId); err != nil {
		return fmt.Errorf("error with inputing the recipient id: %v", err)
	}

	fmt.Print("Input message itself(we don't have newlines, sorry): ")
	if _, err := fmt.Scan(&payload.Value); err != nil {
		return fmt.Errorf("error with inputing the message: %v", err)
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

func SendGetRequestWithJWT(path, page, token string) ([]types.Message, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+path+"?page="+page, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("error with sending the request: %v", err)
	}
	defer resp.Body.Close()

	gottenResult, err := getMultipleResult(resp.Body)
	if err != nil {
		return nil, err
	}

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("error has gotten from the server: %s", (*gottenResult)["error"])
	// }

	return gottenResult, nil
}

func SendPostRequestWithJWT(path string, data []byte, token string) (*map[string]string, error) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080"+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
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

func SendDeleteRequestWithJWT(path string, token string) (*map[string]string, error) {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080"+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
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

func getMultipleResult(body io.ReadCloser) ([]types.Message, error) {
	var gottenResult []types.Message

	if err := json.NewDecoder(body).Decode(&gottenResult); err != nil {
		return nil, fmt.Errorf("error has gotten while decoding the response: %v", err)
	}

	return gottenResult, nil
}
