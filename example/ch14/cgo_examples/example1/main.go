package main

import "fmt"

/*
	#cgo LDFLAGS: -lm
	#include <stdio.h>
	#include <math.h>
	#include "mylib.h"

	int add(int a, int b) {
		int sum = a + b;
		printf("add内 -- a: %d, b: %d, 合計 %d\n", a, b, sum);
		return sum;
	}
*/
import "C"

func main() {
	sum := C.add(3, 2)
	fmt.Println("sum:", sum)
	fmt.Println("C.sqrt(100):", C.sqrt(100))
	fmt.Println("C.multiply(10, 20):", C.multiply(10, 20))
}
