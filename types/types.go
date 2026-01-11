package types

type RegisterPayload struct {
	id       string `json:"id"`
	password string `json:"password"`
	nickname string `json:"nickname"`
}

type LoginPayload struct {
	id       string `json:"id"`
	password string `json:"password"`
}
