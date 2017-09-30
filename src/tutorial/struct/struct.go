package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int32
	mess string
}

func newUser(name string, age int32, mess string) *User {
	return &User{Name: name, Age: age, mess: mess}
}

func main() {
	var user User
	user.Name = "nick"
	user.Age = 18
	user.mess = "lover"

	var user1 = &User{
		Name: "dawn",
		Age:  21,
	}
	fmt.Println(*user1)
	fmt.Println(user1.Name, (*user1).Name)

	var user2 = new(User)
	user2.Name = "bogerv"
	user2.Age = 18
	fmt.Println(*user2)
	fmt.Println(user2.Name, (*user2).Name)

	// 模拟构造函数
	user3 := newUser("nater", 20, "lover nater")
	fmt.Println(*user3)
}
