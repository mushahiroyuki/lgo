package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error { //liststart
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("in main, wrappedErr:", wrappedErr)
		}
	}
} //listend
