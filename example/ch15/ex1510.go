package main

import (
	"fmt"
	"strings"
)

func main() {
	t1 := NewTree(BuiltInOrderable[int]) //liststart6
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15)) // true
	fmt.Println(t1.Contains(40)) // false //listend6

	t2 := NewTree(OrderPeople) //liststart8
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30})) // true
	fmt.Println(t2.Contains(Person{"Fred", 25})) // false //listend8

	t3 := NewTree(Person.Order) //liststartA
	t3.Add(Person{"Bob", 30})
	t3.Add(Person{"Maria", 35})
	t3.Add(Person{"Bob", 50})
	fmt.Println(t3.Contains(Person{"Bob", 30})) // true
	fmt.Println(t3.Contains(Person{"Fred", 25})) // false //listendA
}

type BuiltInOrdered interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type OrderableFunc[T any] func(t1, t2 T) int  //liststart0
//listend0

type Tree[T any] struct { //liststart1
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
} //listend1

func NewTree[T any](f OrderableFunc[T]) *Tree[T] { //liststart2
	return &Tree[T]{
		f: f,
	}
} //listend2

func (t *Tree[T]) Add(v T) { //liststart3
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
} //listend3

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] { //liststart4
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
} //listend4

func BuiltInOrderable[T BuiltInOrdered](t1, t2 T) int { //liststart5
	if t1 < t2 {
		return -1
	}
	if t1 > t2 {
		return 1
	}
	return 0
} //listend5

type Person struct {  //liststart7
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := strings.Compare(p1.Name, p2.Name) // Nameで比較する
	if out == 0 { // 同じならば
		out = p1.Age - p2.Age  // Ageで比較する
	}
	return out
} //listend7

func (p Person) Order(other Person) int { //liststart9
	out := strings.Compare(p.Name, other.Name)
	if out == 0 {
		out = p.Age - other.Age
	}
	return out
} //listend9
