package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello"
	b := []byte("goodbye")
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sHdr.Data = bHdr.Data
	sHdr.Len = bHdr.Len
	fmt.Println(s)

	b[0] = 'x'
	fmt.Println(s)
}
