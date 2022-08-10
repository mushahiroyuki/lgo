package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)


func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
//liststart1	
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
//listend1
//liststart2
	if diff := cmp.Diff(expected, result, comparer); diff != "" { 
		t.Error(diff)
	}
//listend2	
	if result.DateAdded.IsZero() {
		t.Error("DateAdded wasn't assigned")
	}
}
