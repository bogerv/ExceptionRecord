package main

import (
	"fmt"
)

type Phone interface {
	call()
	sales() int
}

type NokiaPhone struct {
	price int
}

func(n NokiaPhone) call() {
	fmt.Println("Nokia call")
}

func(n NokiaPhone) sales() int {
	return n.price
}

type IPhone struct {
	price	int
}

func(n IPhone) call() {
	fmt.Println("IPhone call")
}

func(i IPhone) sales() int {
	return i.price
}

type Person struct {
	phones []Phone
	name string
	age int
}
func (p Person) total_cost() int {
	var sum = 0
	for _, phone := range p.phones {
		sum += phone.sales()
	}
	return sum
}

func main() {
	var bought_phones = [5]Phone {
		NokiaPhone{price: 350},
		IPhone{price: 5000},
		IPhone{price: 3400},
		NokiaPhone{price: 450},
		IPhone{price: 5000},
	}

	var person = Person{name: "Jemy", age: 25, phones: bought_phones[:]}

	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.total_cost())
}