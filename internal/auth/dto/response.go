package dto

type SignInResponse struct {
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message"`
}
