package main

import "fmt"

type User struct {
	Name string
	Age  int
	sex  string
}

func (this *User) init(name string, age int, sex string) {
	this.Name = name
	this.Age = age
	this.sex = sex
}

func (this User) getName() string {
	return this.Name
}

func main() {
	var user User
	user.init("bogerv", 18, "男")

	fmt.Println(user.getName())
}
