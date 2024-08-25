package util

import (
	"reflect"
	"testing"
)

// 定义一个结构体，带有 tag
type User struct {
	Name string `map:"name"`
	Age  int    `map:"age"`
}

// TestMapToStruct_Success 用于测试 MapToStructByMapTag 的成功情况
func TestMapToStruct_Success(t *testing.T) {
	data := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	expected := User{
		Name: "Alice",
		Age:  30,
	}

	var user User
	err := MapToStructByMapTag(data, &user)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected %v, but got %v", expected, user)
	}
}

// TestMapToStruct_MissingFields 测试 MapToStructByMapTag 的部分字段缺失情况
func TestMapToStruct_MissingFields(t *testing.T) {
	data := map[string]interface{}{
		"name": "Bob",
		// "age" 字段缺失
	}

	expected := User{
		Name: "Bob",
		Age:  0, // 默认为0
	}

	var user User
	err := MapToStructByMapTag(data, &user)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("Expected %v, but got %v", expected, user)
	}
}

// TestMapToStruct_InvalidType 测试 MapToStructByMapTag 的类型不匹配情况
func TestMapToStruct_InvalidType(t *testing.T) {
	data := map[string]interface{}{
		"name": "Charlie",
		"age":  "thirty", // 类型错误：string 而不是 int
	}

	var user User
	err := MapToStructByMapTag(data, &user)
	if err == nil {
		t.Fatal("Expected an error due to type mismatch, but got no error")
	} else {
		expectedErrMsg := "type mismatch for field Age: expected int but got string"
		if err.Error() != expectedErrMsg {
			t.Fatalf("Expected error message '%v', but got '%v'", expectedErrMsg, err.Error())
		}
	}
}
