package types

type RegisterPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginPayload struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
