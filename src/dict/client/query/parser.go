package query

import "dict/client/model"

// html解析器
type Parser interface {
	// 将结果页面html解析成单例释义
	Parse(html string) *model.Word
}
