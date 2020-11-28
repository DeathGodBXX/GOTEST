package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s1 := &Student{
		Id:     3,       //通过tag制定json序列化的key值
		Gender: "male",  //json序列化，使用默认字段名作为key
		name:   "Geoge", //小写表示私有属性，外部包无法读取（json）
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("Marshal failed\n")
	}
	fmt.Printf("%s\n", data)
}

type Student struct {
	Id     int8   `json:"id"`
	Gender string `json:"gender"`
	name   string
}
