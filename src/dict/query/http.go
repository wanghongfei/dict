package query

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"bytes"
)

// 下载html代码
func GetHtml(url string) string {
	resp, err := http.Get(url)
	if nil != err {
		log.Fatal(err)
		return ""
	}

	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf)
}

func Post(url string, body []byte) string {
	resp, err := http.Post(url, "application/json;charset=utf8", bytes.NewReader(body))
	if nil != err {
		log.Fatal(err)
		return ""
	}

	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf)

}
