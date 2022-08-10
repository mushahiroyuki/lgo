package main
import (
	"fmt"
)

func main() {
	var x int = 10 //liststart
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d) // 40.2 40   //listend
}
