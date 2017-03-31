package query

import (
	"net/http"
	"io/ioutil"
)

// 下载html代码
func GetHtml(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	return string(buf)
}
