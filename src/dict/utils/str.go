package utils

import "bytes"

// 高效字符串拼接
func Concat(strs []string) string {
	buffer := bytes.Buffer{}


	LEN := len(strs)
	for ix := 0; ix < LEN; ix++ {
		buffer.WriteString(strs[ix])
	}

	return buffer.String()
}
