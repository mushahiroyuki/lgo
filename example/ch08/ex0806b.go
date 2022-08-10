package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

type StatusErr struct { //liststart
	Status Status
	Message string
	Err error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
} //listend

func fileChecker(name string) error {
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
} 
