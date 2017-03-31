package query

import (
	"dict/model"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)

// 爱词霸html解析器
type IcibaParser struct {

}

func (me *IcibaParser) Parse(html string) *model.Word {

	// 构造单词对象
	word := &model.Word{}
	word.Exps = make([]*model.Explanation, 0, 5)

	// 读取html字符串
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if nil != err {
		fmt.Println(err)
		// 返回空word对象
		return model.EmptyWord
	}

	// 取出音标
	existNode := doc.Find(".base-speak span")
	if existNode.Size() == 0 {
		// 返回空word对象
		return model.EmptyWord
	}

	existNode.Each(func(i int, node *goquery.Selection) {
		word.Pronunciation = node.Text()
	})


	// 找到释义节点
	node := doc.Find(".collins-section").First()

	// 遍历每条释义
	node.Find(".section-prep").Each(func(i int, node *goquery.Selection) {

		exp := new(model.Explanation)
		exp.Property = node.Find(".size-chinese .family-english").First().Text()
		exp.CnExplanation = node.Find(".size-chinese .family-chinese").Text()
		exp.EnExplanation = node.Find(".size-chinese .prep-en").Text()
		exp.Sentences = make([]model.Sentence, 0, 5)

		word.Exps = append(word.Exps, exp)

		// 例句
		node.Find(".text-sentence").Each(func(i int, node *goquery.Selection) {
			sentence := model.Sentence{}
			sentence.English = node.Find(".family-english").Text()
			sentence.Chinese = node.Find(".family-chinese").Text()

			exp.Sentences = append(exp.Sentences, sentence)
		})
	})


	return word
}
