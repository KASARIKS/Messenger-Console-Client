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

type MessagePayload struct {
	RecipientId string `json:"recipientId"`
	Value       string `json:"value"`
}

type Message struct {
	Id          int    `json:"id"`
	AuthorId    string `json:"authorId"`
	RecipientId string `json:"recipientId"`
	Value       string `json:"value"`
	CreatedAt   string `json:"createdAt"`
}
