package models

type ReqSendingKafka struct {
	Topic string                 `json:"topic"`
	Data  map[string]interface{} `json:"data"`
}

type RespSendingKafka struct {
	RespCode string `json:"resp_code"`
	RespMsg  string `json:"resp_msg"`
	RespData string `json:"resp_data"`
}
