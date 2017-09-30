package main

import "encoding/json"
import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user User
	user.Name = "bogerv"
	user.Age = 20

	conJSON, _ := json.Marshal(user)
	fmt.Println(conJSON)
	fmt.Println(string(conJSON))
}
