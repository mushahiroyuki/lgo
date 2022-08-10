package main

import(
	"fmt"
	"time"
)

func main() {
	_ = timeTest()
}

func timeTest() error {
//	t, err := time.Parse("2006-02-01 15:04:05 -0700", "2022-23-09 12:34:56 +0900")
	t, err := time.Parse("2006年1月2日 PM3:04:05 -0700", "2022年07月15日 PM6:34:56 +0900") //liststart1
	if err != nil {
		return err
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	fmt.Println(t.Format("2006年1月2日 15時4分5秒"))
	fmt.Println(t.Format("2006.01.02 15:04:05" ))
	fmt.Println(t.Format("1/2/2006 15:04:05 MST"))
//listend1
	return nil
}
