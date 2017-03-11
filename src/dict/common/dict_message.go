package common

import "dict/client/model"

// 通信数据格式
type DictMessage struct {
	// 操作类型
	Op int `json:"op"`

	// 状态码
	St int `json:"st"`

	// 要查询的单词
	Word string `json:"word"`

	// 查询结果
	Result *model.Word `json:"result"`

	// 为true时表示是从数据库中取到的数据
	// 为false表示从远端抓取的数据
	Cached bool `json:"cached"`
}

const (
	// 查询
	OP_QUERY = 1

	// 响应
	OP_RESULT
)

const (
	ST_SUCCESS = 0
	ST_FAILED = 1
	ST_INVALID_METHOD = 2
)

// 创建单词查询请求数据
// word是要查询的单词
func NewQueryDictMessage(word string) *DictMessage {
	return &DictMessage{
		Op: OP_QUERY,
	}
}

// 创建响应数据
// word: 指向查询结果的指针
// cached: 表明是否是从本地数据库中找到的数据
func NewResultDictMessage(word *model.Word, cached bool) *DictMessage {
	return &DictMessage{
		Result: word,
		Cached: cached,
		St: ST_SUCCESS,
	}
}

func NewErrorDictMessage(code int) *DictMessage {
	return &DictMessage{
		St: code,
	}
}


