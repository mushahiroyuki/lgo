package main

type Foo struct {
	Field1 string
	Field2 int
}


func MakeFoo() (Foo, error) { //liststart1
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
} //listend1

func main() {
	_, _ = MakeFoo()
}
