package main

import (
	"fmt"
)

type LinkedList struct { //liststart
	Value interface{} // any も可
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
} //listend

func (ll *LinkedList) String() string {
	if ll == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v->%s", ll.Value, ll.Next)
}

func main() {
	var ll *LinkedList
	fmt.Println(ll)
	ll = ll.Insert(0, 10)
	fmt.Println(ll)
	ll = ll.Insert(0, "foo")
	fmt.Println(ll)
	ll = ll.Insert(1, 4.5)
	fmt.Println(ll)
	ll = ll.Insert(10, false)
	fmt.Println(ll)
	ll = ll.Insert(2, struct{}{})
	fmt.Println(ll)
}
