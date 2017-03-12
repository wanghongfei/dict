package handler

import (
	"net/http"
	"dict/common"
	"encoding/json"
	"dict/client/model"
)

// 响应错误信息
func SendErrorMessage(w http.ResponseWriter, op int) error {
	result := common.NewErrorDictMessage(op)
	buf, err := json.Marshal(result)

	w.Write(buf)

	return err
}

func SendResultMessage(w http.ResponseWriter, word *model.Word) error {
	result := common.NewResultDictMessage(word, false)
	buf, err := json.Marshal(result)

	w.Write(buf)

	return err;
}
