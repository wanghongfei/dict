package handler

import (
	"net/http"
	"dict/common"
	"encoding/json"
	"io/ioutil"
	"dict/server/logic"
	"dict/model"
)

type QueryHandler struct {

}

// 响应查询请求
func (me *QueryHandler) Handle(req *http.Request) *common.DictMessage {
	// 只支持POST请求
	if req.Method != "POST" {
		return common.NewErrorDictMessage(common.ST_INVALID_METHOD)
	}

	// 读取请求体
	buf, _ := ioutil.ReadAll(req.Body)

	// 反序列化
	payload := &common.DictMessage{}
	json.Unmarshal(buf, payload)

	// 调用业务逻辑
	result := logic.QueryWord(payload.Word)

	// 没查到
	if result == model.EmptyWord {
		return common.NewErrorDictMessage(common.ST_FAILED)
	}

	//log.Println(string(buf))
	//log.Println(payload)

	//SendResultMessage(w, result)
	return common.NewResultDictMessage(result, false)
}
