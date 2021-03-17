package utils

import "encoding/json"

type RespMsg struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewSuccessMsg(data interface{}) *RespMsg {
	return NewRespMsg(0, "success", data)
}

func NewRespMsg(code int64, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (resp *RespMsg) JsonByte() []byte {
	bs, _ := json.Marshal(resp)
	return bs
}
func (resp *RespMsg) JsonString() string {
	bs, _ := json.Marshal(resp)
	return string(bs)
}
