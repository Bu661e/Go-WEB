package web_func

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func Test09_json() {
	jsonStr := `{
		"name": "Bu661e",
		"age": 20,
		"hobbies": ["coding", "reading", "running"]
	}`
	com := Company{}

	// 1.将json字符串转换为结构体
	_ = json.Unmarshal([]byte(jsonStr), &com)
	fmt.Println(com)

	// 2.将结构体转换为json字符串
	bytes, _ := json.Marshal(com)
	fmt.Println(string(bytes))
	bytes1, _ := json.MarshalIndent(com, "", "    ")
	fmt.Println(string(bytes1))

}
