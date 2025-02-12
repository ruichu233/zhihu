package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func TransStructToMap(st interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	b, err := json.Marshal(st)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	for k, v := range res {
		switch v.(type) {
		case float64:
			temp := int64(v.(float64))
			if float64(temp) == v.(float64) {
				res[k] = temp
			}
			break
		}
	}
	return res, nil
}

// 将 map 转换为 struct 的函数
func MapToStruct(data map[string]interface{}, result interface{}) error {
	val := reflect.ValueOf(result)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("result 参数必须是非空的 struct 指针")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("result 必须指向一个 struct")
	}

	for key, value := range data {
		field := val.FieldByName(key)
		if !field.IsValid() {
			continue // 跳过 struct 中不存在的字段
		}
		if !field.CanSet() {
			continue // 跳过不可设置的字段
		}

		// 将 value 转换为字段的类型
		val := reflect.ValueOf(value)
		if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		}
	}

	return nil
}

func ParseTimestamp(timestamp int64) string {
	// 将时间戳转换为时间
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return timeStr
}
