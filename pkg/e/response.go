package e

type ResponseData struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}
