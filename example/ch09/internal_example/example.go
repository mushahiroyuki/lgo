package internal_package_example

import "github.com/learning-go-book/internal_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}
