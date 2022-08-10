package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i64 int64 = 1_000_000
	i64ptr := unsafe.Pointer(&i64)
	bytes := *(*[unsafe.Sizeof(i64)]byte)(i64ptr)
	fmt.Println(bytes)
	i32 := *(*int32)(i64ptr)
	fmt.Println(i32)

	i64ptrShifted := unsafe.Pointer(uintptr(unsafe.Pointer(&i64)) + 4)
	bytesShifted := *(*[unsafe.Sizeof(i64)]byte)(i64ptrShifted)
	fmt.Println(bytesShifted)

	i32Shifted := *(*int32)(unsafe.Pointer(&bytesShifted))
	fmt.Println(i32Shifted)

	i64Shifted := *(*int64)(unsafe.Pointer(&bytesShifted))
	fmt.Println(i64Shifted)
}
