package main

import (
	"fmt"
	"structs/zoo"
)

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

const (
	Snowy    string = "snowy"
	Aquatic  string = "aquatic"
	Savannah string = "savannah"
	Jungle   string = "jungle"
)

func main() {
	animal1 := zoo.NewAnimal("Lion", Savannah)
	animal2 := zoo.NewAnimal("Penguin", Snowy)
	animal3 := zoo.NewAnimal("Hippopotamus", Aquatic)
	animal4 := zoo.NewAnimal("Arctic_bear", Snowy)
	animal5 := zoo.NewAnimal("Monkey", Jungle)
	animal6 := zoo.NewAnimal("Snake", Jungle)
	animal7 := zoo.NewAnimal("Panther", Jungle)

	cages := []zoo.Cage{
		{
			Number: 1,
			Biome:  Snowy,
		},
		{
			Number: 2,
			Biome:  Savannah,
			Animal: animal1, // This animal didn't run away
		},
		{
			Number: 3,
			Biome:  Snowy,
		},
		{
			Number: 4,
			Biome:  Aquatic,
		},
		{
			Number: 5,
			Biome:  Jungle,
		},
		{
			Number: 6,
			Biome:  Jungle,
		},
	}

	z1 := zoo.Zookeeper{
		Name:  "Oleksiy",
		Cages: []*zoo.Cage{&cages[2], &cages[3], &cages[4], &cages[5]},
	}

	z2 := zoo.Zookeeper{
		Name:  "Kyrylo",
		Cages: []*zoo.Cage{&cages[0], &cages[1], &cages[2], &cages[3]},
	}

	// Animals ran away
	for _, cage := range cages {
		if cage.Animal != nil {
			fmt.Printf("An animal %v is in a cage #%v\n", cage.Animal.Name, cage.Number)
		} else {
			fmt.Printf("There is no animal in a cage #%v\n", cage.Number)
		}
	}

	fmt.Println("----------")
	// Catch animals
	z1.CageAnimal(animal2)
	z2.CageAnimal(animal3)
	z2.CageAnimal(animal4)
	z1.CageAnimal(animal5)
	z1.CageAnimal(animal6)

	// Oleksiy has no free room left. All cages are taken
	z1.CageAnimal(animal7)
	// Aswell as Kyrylo
	z2.CageAnimal(animal7)

	fmt.Println("----------")
	// Animals were catched
	for _, cage := range cages {
		if cage.Animal != nil {
			fmt.Printf("An animal %v is in a cage #%v\n", cage.Animal.Name, cage.Number)
		} else {
			fmt.Printf("There is no animal in a cage #%v\n", cage.Number)
		}
	}
}
