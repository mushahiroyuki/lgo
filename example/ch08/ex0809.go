package main

import "os"

func doPanic(msg string) { //liststart
	panic(msg)
}

func main() {
	doPanic(os.Args[0])
} //listend

