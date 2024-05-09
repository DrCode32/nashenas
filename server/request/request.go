package request

type SetUsername struct {
	Username string `json:"username"`
}

type SetPublicKey struct {
	PublicKey string `json:"public_key"`
}

type SendMessage struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}