package pet

import "github.com/learning-go-book/circular_dependency_example/person"

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

var owners = map[string]person.Person{ //liststart1
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
} //listend1

func (p Pet) Owner() person.Person {
	return owners[p.OwnerName]
}
