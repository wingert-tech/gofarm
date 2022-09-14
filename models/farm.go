package models

import (
	"fmt"
	"reflect"
	"test/interfaces"
)

type Farm struct {
	Name    string
	Animals []interfaces.Animal
}

func (f Farm) String() string {

	return fmt.Sprintf("%v %v", f.Name, f.Animals)
}

func (f Farm) FeedAnimals(animals []interfaces.Animal) {

	defer f.finishedFeeding()
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a))
		switch a.(type) {
		case Livestock:
			fmt.Println("Getting hay")
		case Poultry:
			fmt.Println("Getting scratch")
		}
		a.Eat()
	}

}

func (f Farm) finishedFeeding() {
	fmt.Println("Feeding chores done for", f.Name)
}
