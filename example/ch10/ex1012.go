package main

import (
	"fmt"
	"sync"
)

func main() { //liststart1
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		doThing1()
	}()

	go func() {
		defer wg.Done()
		doThing2()
	}()

	go func() {
		defer wg.Done()
		doThing3()
	}()

	wg.Wait()
} //listend1

func doThing1() {
	fmt.Println("Thing 1 done!")
}

func doThing2() {
	fmt.Println("Thing 2 done!")
}

func doThing3() {
	fmt.Println("Thing 3 done!")
}
