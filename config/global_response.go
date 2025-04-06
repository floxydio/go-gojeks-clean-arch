package config

type GlobalResponseMsg struct {
	Status  int    `json:"status"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type GlobalResponseData struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
