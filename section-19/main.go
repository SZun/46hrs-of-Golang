package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	data := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	var user []user
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {

	}
	fmt.Println(user)
}
