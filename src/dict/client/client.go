package main

import (
	"fmt"
	"dict/query"
	"strings"
	"dict/model"
	"flag"
	"dict/common"
	"encoding/json"
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
	host := "http://" + dictServer + ":" + dictPort

	pkg := buildQueryRequest(word)
	buf, _ := json.Marshal(pkg)
	resp := query.Post(host, buf)

	result := new(model.Word)
	json.Unmarshal([]byte(resp), result)

	return result
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
