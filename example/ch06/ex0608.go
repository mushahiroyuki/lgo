package main

// import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error { //liststart1
	f.Field1 = "val"
	f.Field2 = 20
	return nil
} //listend1

func main() {
	var f Foo
	_ = MakeFoo(&f)
}
