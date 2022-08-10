package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello"
	sHdrData := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // prints 5

	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(uintptr(sHdrData) + uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()
	sHdr.Len = sHdr.Len + 10
	fmt.Println(s)
	bp := (*byte)(unsafe.Pointer(uintptr(sHdrData) + 2))
	*bp = *bp + 1
	fmt.Println(s)
}
