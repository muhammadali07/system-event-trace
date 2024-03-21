package models

type ResponseApp struct {
	RespCode string `json:"response_code"`
	RespMsg  string `json:"response_msg"`
	Respdata any    `json:"data"`
}
