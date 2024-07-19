package main

import (
	"encoding/json"
	"fmt"
)




type User struct {
    Name string
}

func test() {
    user := &Message{Id: 1,Text: "Frank",IsChecked: true}
    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
}