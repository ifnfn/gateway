package util

import (
	"encoding/json"
	"fmt"

	log "qiniupkg.com/x/log.v7"

	"github.com/fgrid/uuid"
)

// UUID return uuid string
func UUID() string {
	return uuid.NewV4().String()
}

// StructToString 将结构体转成字符串
func StructToString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case int:
		return fmt.Sprintf("%d", t)
	default:
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Warn(err)
		}
		str := string(b)

		return str
	}
}

// PrintInterface 显示 Interface 接口
func PrintInterface(i interface{}) {
	s := StructToString(i)
	println(s)
}
