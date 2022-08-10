package main

import (
	"fmt"

	"github.com/learning-go-book/circular_dependency_example/person"
)

func main() {
	bob := person.Person{PetName: "Fluffy"}
	fmt.Println(bob.Pet())
}
