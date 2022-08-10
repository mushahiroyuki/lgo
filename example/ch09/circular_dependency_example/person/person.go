package person

import "github.com/learning-go-book/circular_dependency_example/pet"

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{ //liststart1
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
} //listend1

func (p Person) Pet() pet.Pet {
	return pets[p.PetName]
}
