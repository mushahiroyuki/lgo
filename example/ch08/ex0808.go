package main

import (
	"errors"
	"fmt"
)

func main() {
	err := ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	if errors.Is(err, ResourceErr{Resource: "Database"}) { //liststart3
		fmt.Println("The database is broken:", err) // データベースが壊れている
		// ... コード ...
	} //listend3

	if errors.Is(err2, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err2)
	}

	if errors.Is(err, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err)
	}

	if errors.Is(err2, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err2)
	}

	if errors.Is(err, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered:", err)
	}

	if errors.Is(err, ResourceErr{Resource: "Network", Code: 123}) {
		fmt.Println("Network is broken and Code 123 triggered:", err)
	}
}

type ResourceErr struct { //liststart1
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
} //listend1

func (re ResourceErr) Is(target error) bool { //liststart2
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
} //listend2
