package main

import (
	"fmt"
	"dict/query"
	"strings"
	"dict/model"
	"flag"
	"dict/common"
	"encoding/json"
	"os"
	"dict/utils"
)

func main() {
	// 取命令行参数
	wordToQuery := flag.String("w", "", "指定要查询的单词")
	dictServer := flag.String("host", "127.0.0.1", "指定dict服务地址")
	dictPort := flag.String("port", "9000", "指定dict服务端口")
	queryFromIciba := flag.Bool("iciba", false, "带上该标记表示直接从iciba爬取")
	flag.Parse()

	if *queryFromIciba {
		result := queryIciba(*wordToQuery)
		show(result)

	} else {
		result := queryDictServer(*wordToQuery, *dictServer, *dictPort)
		show(result)
	}
}

func queryDictServer(word, dictServer, dictPort string) *model.Word {
	host := utils.Concat([]string {
		"http://",
		dictServer,
		":",
		dictPort,
		"/query",
	})

	// 生成请求消息对象
	pkg := buildQueryRequest(word)
	// 序列化
	buf, _ := json.Marshal(pkg)
	// POST请求
	resp := query.Post(host, buf)

	// 反序列化结果
	result := new(common.DictMessage)
	err := json.Unmarshal([]byte(resp), result)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	return result.Result
}

func queryIciba(word string) *model.Word {
	html := query.GetHtml("http://iciba.com/" + word)

	// 解析
	parser := new(query.IcibaParser)
	result := parser.Parse(html)
	result.Literal = word

	return result
}


func trimString(str string) string {
	s1 := strings.TrimLeft(str, "\n ")
	return strings.TrimRight(s1, "\n ")

}

func show(word *model.Word) {
	if word == model.EmptyWord {
		fmt.Println("未查到释义")
		return
	}

	fmt.Printf("%s - %s\n", word.Literal, word.Pronunciation)
	for ix, exp := range word.Exps {
		fmt.Printf("%d. [%s] %s(%s)\n", ix + 1, exp.Property, exp.EnExplanation, exp.CnExplanation)

		for ix, sentence := range exp.Sentences {
			fmt.Printf("\t%d. %s(%s)\n", ix + 1, trimString(sentence.English), trimString(sentence.Chinese))
		}

	}
}

func buildQueryRequest(word string) *common.DictMessage  {
	return &common.DictMessage{
		Op: common.OP_QUERY,
		Word: word,
	}
}
