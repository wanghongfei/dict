package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"
	"io/ioutil"
	"dict/common"
)

func main()  {
	log.Println("启动server")

	http.HandleFunc("/query", queryHandler)

	err := http.ListenAndServe(":9000", nil)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println("hello, server")
}

// 处理单词查询请求
func queryHandler(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Add("Content-Type", "application/json;charset=utf8")

	// 只支持POST请求
	if req.Method != "POST" {
		result := common.NewErrorDictMessage(common.ST_INVALID_METHOD)
		buf, _ := json.Marshal(result)
		resp.Write(buf)

		return
	}

	// 读取请求体
	buf, _ := ioutil.ReadAll(req.Body)

	// 反序列化
	payload := &common.DictMessage{}
	json.Unmarshal(buf, payload)

	log.Println(string(buf))
	log.Println(payload)

	io.WriteString(resp, "OK")
}
