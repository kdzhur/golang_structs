package zoo

import (
	"fmt"
)

type animal struct {
	Name         string
	habitatBiome string
}

func NewAnimal(name, habitatBiome string) *animal {
	a := &animal{
		Name:         name,
		habitatBiome: habitatBiome,
	}

	return a
}

type Zookeeper struct {
	Name  string
	Cages []*Cage // cages that zookeeper owns
}

func (z *Zookeeper) CageAnimal(a *animal) {
	for i := range z.Cages {
		if z.Cages[i].canFitAnimal(a) {
			z.Cages[i].fitAnimal(a)
			fmt.Printf("Zookeper %v has caged animal %v to the cage #%v\n", z.Name, a.Name, z.Cages[i].Number)
			return
		}
	}
	fmt.Printf("Zookeper %v has no free cage with %v biome for animal %v\n", z.Name, a.habitatBiome, a.Name)
}

type Cage struct {
	Number int
	Biome  string
	Animal *animal
}

func (c *Cage) fitAnimal(a *animal) {
	c.Animal = a
}

func (c *Cage) canFitAnimal(a *animal) bool {
	return (c.Animal == nil && a.habitatBiome == c.Biome)
}
