package util

import (
	"errors"
	"fmt"
	"reflect"
)

// MapToStruct 将 map 反序列化到指定的结构体
func MapToStruct(source map[string]interface{}, dest interface{}) error {
	// 检查 dest 是否是指针且指向结构体
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("dest should be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	// 遍历结构体的字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("map")

		if tag == "" {
			tag = field.Name
		}

		// 从 map 中获取值
		if v, ok := source[tag]; ok {
			fieldVal := val.Field(i)

			// 检查类型是否匹配
			if fieldVal.CanSet() && reflect.TypeOf(v) == fieldVal.Type() {
				fieldVal.Set(reflect.ValueOf(v))
			} else if fieldVal.CanSet() && reflect.TypeOf(v) != fieldVal.Type() {
				return fmt.Errorf("type mismatch for field %s: expected %s but got %s", field.Name, fieldVal.Type(), reflect.TypeOf(v))
			}
		}
	}

	return nil
}
