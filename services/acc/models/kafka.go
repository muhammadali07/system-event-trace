package models

type RespSendingKafka struct {
	RespCode string `json:"resp_code"`
	RespMsg  string `json:"resp_msg"`
	RespData string `json:"resp_data"`
}
