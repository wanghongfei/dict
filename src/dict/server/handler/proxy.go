package handler

import (
	"net/http"
	"dict/common"
	"encoding/json"
	"io/ioutil"
	"log"
	"io"
)

//JsonHandler实现了http.Handler接口,因此可以直接
//将其赋值给http.Handle()方法的第2个参数.
type JsonHandlerProxy struct {
	handler http.Handler
}

// 创建一个http请求处理器实例
// handler: 封装了用户代码逻辑的对象
func NewHandler(handler http.Handler) *JsonHandlerProxy {
	result := &JsonHandlerProxy{
		handler: handler,
	}

	return result
}

// 该方法在执行用户逻辑之前先设置请求头
func (me JsonHandlerProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 设置Content-Type头
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// 调用用户代码逻辑
	me.handler.ServeHTTP(w, req)
}

type QueryHandler struct {

}

func (me QueryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 只支持POST请求
	if req.Method != "POST" {
		result := common.NewErrorDictMessage(common.ST_INVALID_METHOD)
		buf, _ := json.Marshal(result)
		w.Write(buf)

		return
	}

	// 读取请求体
	buf, _ := ioutil.ReadAll(req.Body)

	// 反序列化
	payload := &common.DictMessage{}
	json.Unmarshal(buf, payload)

	log.Println(string(buf))
	log.Println(payload)

	io.WriteString(w, "OK")

}
