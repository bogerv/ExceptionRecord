package main

import (
	"fmt"
)

type Car interface {
	NameGet() string
	Run(n int)
	Stop()
}
type BMW struct {
	Name string
}

func (this *BMW) NameGet() string {
	return this.Name
}
func (this *BMW) Run(n int) {
	fmt.Println("BMW is running of num is %d \n", n)
}
func (this *BMW) Stop() {
	fmt.Println("BMW is stop \n")
}

type Benz struct {
	Name string
}

func (this *Benz) NameGet() string {
	return this.Name
}
func (this *Benz) Run(n int) {
	fmt.Printf("Benz is running of num is %d \n", &n)
}
func (this *Benz) Stop( {
	fmt.Printf("Benz is stop \n")
}
func (this *Benz) ChatUp() {
	fmt.Printf("ChatUp \n")
}

func main() {
	var car Car
	fmt.Println(car)

	var bmw BMW = BMW{Name: "baoma"}
	car = &bmw
	fmt.Println(car.NameGet())
	car.Run(1) //BMW is running of num is 1
	car.Stop() //BMW is stop

	benz := &Benz{Name: "大奔"}
	car = benz
	fmt.Println(car.NameGet()) //大奔
	car.Run(2)                 //Benz is running of num is 2
	car.Stop()                 //Benz is stop
}
