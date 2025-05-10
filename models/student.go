package models

import "strings"

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *Student) GetType() string {
	if strings.HasSuffix(s.Name, "e") {
		return "female"
	} else {
		return "male"
	}

}
