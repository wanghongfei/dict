package mongodb

import (
	"testing"
	"dict/client/model"
)

var wordApple *model.Word

func init() {
	exp := &model.Explanation{
		Property: "n",
		EnExplanation: "a kind of fruit",
		CnExplanation: "苹果",
	}

	exps := make([]*model.Explanation, 0, 1)
	exps = append(exps, exp)

	wordApple = &model.Word{
		Literal: "apple",
		Pronunciation: "发音",
		Exps: exps,
	}
}

func TestSave(t *testing.T)  {
	store, err := NewMongodbStore("127.0.0.1")
	if nil != err {
		t.Fatal(err)
	}

	err = store.Save(wordApple)
	if nil != err {
		t.Fatal(err)
	}
}


