package main

import (
	"fmt"
	"test/interfaces"
	"test/models"
)

func main() {
	farm := models.Farm{}
	farm.Name = "Wingert"
	cow := models.Livestock{}
	cow.Name = "Millie"
	chicken := models.Poultry{}
	chicken.Name = "Henrietta"
	farm.Animals = []interfaces.Animal{
		cow,
		chicken,
	}

	fmt.Println(farm)

	farm.FeedAnimals(farm.Animals)

}
