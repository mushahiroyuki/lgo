package main

import(
	"fmt"
	"time"
)

func main() {
	err := testTime()
	if (err != nil) {
		fmt.Println(err)
	}
}

func testTime() error {

	t, err := time.Parse(time.UnixDate, "Mon Feb 14 06:34:56 JST 2022")
	if err != nil {
		return err
	}
	fmt.Println(t) 
	
	fmt.Println("----------")
	
	t1, err := time.Parse("2006.02.01 15:04:05 -0700", "2022.14.02 06:34:56 +0900")
	if err != nil {
		return err
	}
	fmt.Println(t1.Year()) // 2022
	fmt.Println(t1.Month()) // February
	fmt.Println(t1.Day()) // 14
	fmt.Println(t1.Hour()) // 6
	fmt.Println(t1.Minute()) // 34
	fmt.Println(t1.Second()) // 56
	fmt.Println(t1.Weekday()) // Monday

	fmt.Println("----------")
	t2, err := time.Parse("2006.01.02 3:04:05PM -0700", "2022.01.05 04:34:12AM +0900")
	if err != nil {
		return err
	}
	fmt.Println(t2.After(t)) // false
	fmt.Println(t1.After(t2)) // true

	hour, minute, second := t2.Clock()
	fmt.Printf("%d時%d分%d秒\n", hour, minute, second)  // 4時34分12秒
	// 

	fmt.Println("----------")
	t3, err := time.Parse("2006.01.02 3:04:05PM -0700", "2022.01.05 05:35:14AM +0900")
	if err != nil {
		return err
	}
	
	fmt.Println(t3.Sub(t2)) // 1h1m2s
	fmt.Println(t3.Add(time.Minute*30 + time.Second*5)) // 2022-01-05 06:05:19 +0900 JST

	
	return nil
}
