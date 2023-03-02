package utils

type Respone struct {
	Result  bool        `json:"result" example:"true"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data" `
}
