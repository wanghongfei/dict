package model

import "fmt"

type Sentence struct{
	English string
	Chinese string
}


// 单词解释
type Explanation struct {
	// 词性
	Property string
	// 英文解释
	EnExplanation string
	// 中文解释
	CnExplanation string
	// 例句
	Sentences []Sentence
}

func (me Explanation) String() string {
	return fmt.Sprintf("[%s] %s(%s)", me.Property, me.EnExplanation, me.CnExplanation)
}
