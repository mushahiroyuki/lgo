package main

import(
	"fmt"
	"time"
)

func main() {
	_ = timeTest()
}

func timeTest() error {
//liststart1
	t, err := time.Parse("2006年1月2日 PM3:04:05 -0700", "1986年2月14日 PM6:34:56 +0900")
	
	if err != nil {
		return err
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	fmt.Println(t.Format("2006年1月2日 15時4分5秒"))
	fmt.Println(t.Format("2006.01.02 15:04:05" ))
	fmt.Println(t.Format("1/2/2006 15:04:05 MST"))

	fmt.Println("------------")
	t2, err := time.Parse("2006.01.02 3:04:05PM -0700", "2022.05.05 04:34:00AM +0900")
	fmt.Println(t2.Format("2006年1月2日 午前3時4分5秒"))
	fmt.Println(t2.Format("2006.01.02 15:04:05" ))
	fmt.Println(t2.Format("1/2/2006 15:04:05 MST"))

	fmt.Println(t.Equal(t2)) // false

	t3,err := time.Parse("January 2, 2006 at 3:04:05PM MST", "May 5, 2022 at 04:34:00AM JST")
	fmt.Println(t3.Equal(t2)) // true
	
	return nil
}


