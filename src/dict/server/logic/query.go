package logic

import (
	"dict/server/store"
	"dict/server/store/mongodb"
	"log"
	"dict/model"
	"dict/query"
	"dict/common"
	"strings"
	"fmt"
)

var wordDao store.WordStore

// 初始化mongodb连接
func init()  {
	mongoDao, err := mongodb.NewMongodbStore("127.0.0.1")
	if nil != err {
		log.Fatal(err)
	}

	wordDao = mongoDao
}

// 先从库中查询, 没有则爬取远程, 最后保存到库中
func QueryWord(word string) *model.Word {
	// 从库中查询
	result := wordDao.Load(word)
	if "" != result.Literal {
		return result
	}

	// 没查到, 从远程查询
	parser := new(query.IcibaParser)
	remoteResult := parser.Parse(query.GetHtml(common.URL_ICIBA + word))
	if remoteResult == model.EmptyWord {
		// 没查到
		log.Printf("未从远程查到单词%s\n", word)
		return remoteResult
	}

	remoteResult.Literal = word

	// 保存到本地
	TrimResult(remoteResult)
	wordDao.Save(remoteResult)

	return remoteResult
}

// 去掉爬取结果中出现中换行,空格
func TrimResult(word *model.Word) {
	word.Literal = trimString(word.Literal)

	for _, exp := range word.Exps {
		exp.CnExplanation = trimString(exp.CnExplanation)
		exp.EnExplanation = trimString(exp.EnExplanation)
		exp.Property = trimString(exp.Property)

		for _, sentence := range exp.Sentences {
			sentence.Chinese = trimString(sentence.Chinese)
			sentence.English = trimString(sentence.English)
		}
	}

	fmt.Println(word)
}

// 去掉字符串开头和结尾的空格, 换行
func trimString(str string) string {
	middle := strings.TrimLeft(str, "\n ")
	return strings.TrimRight(middle, "\n ")
}