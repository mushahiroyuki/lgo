package main

import "fmt"

/*
	extern int add(int a, int b);
*/
import "C"

func main() {
	sum := C.add(3, 2)
	fmt.Println(sum)
}

//export doubler  //liststart1
func doubler(i int) int {
	return i * 2
}  //listend1
