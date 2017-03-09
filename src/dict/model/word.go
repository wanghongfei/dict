package model

import "fmt"

// 一个单词
type Word struct {
	// 单词拼写
	Literal string
	// 音标
	Pronunciation string

	// 释译
	Exps []*Explanation
}

var EmptyWord *Word

func init()  {
	EmptyWord = &Word{}
}

func (me Word) String() string {
	return fmt.Sprintf("%s, %s, %s", me.Literal, me.Pronunciation, me.Exps)
}




