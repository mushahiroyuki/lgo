package sibling

import "github.com/learning-go-book/internal_example/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}
