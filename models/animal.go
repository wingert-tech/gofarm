package models

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) String() string {
	return a.Name
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.Name)
}

func (a *Animal) Speak() {
	fmt.Printf("Hi I am %v\n", a.Name)
}
