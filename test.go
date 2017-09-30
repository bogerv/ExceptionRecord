package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("first")
}

type car struct {
	color string
}

func (c car) speak() {
	fmt.Println("My color is: ", c.color)
}

func foo(h human) {
	fmt.Println(h)
}

// type User struct {
// 	Id       int
// 	UserName string
// 	Url      string
// 	Age      int
// }

func main() {
	// db, err := gorm.Open("mysql", "root:Mit000@/bogerv?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	// user := User{UserName: "Jinzhu", Age: 18}
	// db.NewRecord(user)
	// db.Create(&user)
	// var user2 User
	// db.First(&user2, 1)
	// fmt.Println(user2)

	// guid
	// u1 := uuid.NewV4()
	// fmt.Printf("UUIDv4: %s\n", u1)

	p1 := person{"bogerv"}
	p2 := person{"Wang"}
	c := car{"Red"}

	c.speak()

	foo(p1)
	foo(p2)
	foo(c)
}
