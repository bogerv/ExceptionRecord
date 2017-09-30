package main

import (
	"fmt"
	"time"
)

type user1 struct {
	name string
	Age  int
}

type user2 struct {
	name string
	age  int
	sex  time.Time
}

type User struct {
	u1 user1
	user2
	Name string
	Age  int
}

func main() {
	var user User
	user.Name = "bogerv"
	user.u1.Age = 20
	fmt.Println(user)
	fmt.Println(user.user2)
	fmt.Println(user.user2.name)
	fmt.Println(user.user2.age)
	fmt.Println(user.user2.sex)
}
