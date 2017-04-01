package handler

import (
	"net/http"
	"dict/common"
	"encoding/json"
	"log"
)

//JsonHandler实现了http.Handler接口,因此可以直接
//将其赋值给http.Handle()方法的第2个参数.
type JsonHandlerProxy struct {
	handler RestHandler
}

// 创建一个http请求处理器实例
// handler: 封装了用户代码逻辑的对象
func NewHandler(handler RestHandler) *JsonHandlerProxy {
	result := &JsonHandlerProxy{
		handler: handler,
	}

	return result
}

// 该方法在执行用户逻辑之前先设置请求头;
// 并将用户返回的实体对象序列化成json返回
func (me *JsonHandlerProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 设置Content-Type头
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	// 调用用户代码逻辑
	entity := me.handler.Handle(req)

	// json序列化返回对象
	buf, err := json.Marshal(entity)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(buf)
}

// 所有Controller需实现此接口;
// 接收请求参数, 返回实体对象
type RestHandler interface {
	Handle(req *http.Request) *common.DictMessage
}

