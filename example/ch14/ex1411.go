package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	s := "hello"  //liststart1
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // 5 //listend1
	for i := 0; i < sHdr.Len; i++ { //liststart2
		bp := *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()
	runtime.KeepAlive(s) //listend2
}
