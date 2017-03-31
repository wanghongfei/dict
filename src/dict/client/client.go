package main

import (
	"fmt"
	"dict/query"
	"os"
	"strings"
	"dict/model"
)

func main() {
	argWord := "hello"

	// 取命令行参数
	if len(os.Args) > 1 {
		argWord = os.Args[1]
	}

	// GET请求
	html := query.GetHtml("http://iciba.com/" + argWord)

	// 解析
	parser := new(query.IcibaParser)
	word := parser.Parse(html)
	word.Literal = argWord

	// 打印
	show(word)



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
