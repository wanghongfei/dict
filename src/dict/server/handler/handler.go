package handler

import (
	"net/http"
	"dict/common"
	"encoding/json"
	"io/ioutil"
	"log"
)

type QueryHandler struct {

}

// 响应查询请求
func (me QueryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 只支持POST请求
	if req.Method != "POST" {
		SendErrorMessage(w, common.ST_INVALID_METHOD)
		return
	}

	// 读取请求体
	buf, _ := ioutil.ReadAll(req.Body)

	// 反序列化
	payload := &common.DictMessage{}
	json.Unmarshal(buf, payload)

	log.Println(string(buf))
	log.Println(payload)

	SendResultMessage(w, nil)
}
