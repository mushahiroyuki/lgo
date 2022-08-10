package main

import (
	"fmt"
)



func main() {

	fmt.Println("===== 7.2.7　iotaと列挙型 =====")

	{

		type MailCategory int  // メールの分類   //liststart1		
//listend1
		const ( //liststart2
			Uncategorized MailCategory = iota  // 未分類
			Personal  // 個人的
			Spam  // 迷惑メール
			Social  // ソーシャル
			Advertisements  // 広告
		) //listend2

		m := Personal
		fmt.Println("Personal:", m)  // 1
		m = Uncategorized
		fmt.Println("Uncategorized:", m)	// 0
	}

	fmt.Println("----- ビットフィールド -----")	

	{
		type BitField int //liststart3

		const (
			Field1 BitField = 1 << iota // 1が代入される
			Field2                      // 2が代入される
			Field3                      // 4が代入される
			Field4                      // 8が代入される
		) //listend3

		fmt.Println("Field1:", Field1)
		fmt.Println("Field2:", Field2)
		fmt.Println("Field3:", Field3)
		fmt.Println("Field4:", Field4)
	}

	fmt.Println("----- 「_」 -----")	

	{
		type SomeValue int

		const (
			_ SomeValue = iota
			Value1
			Value2
			Value3
			Value4
		)

		fmt.Println("Value1:", Value1)
		fmt.Println("Value2:", Value2)
		fmt.Println("Value3:", Value3)
		fmt.Println("Value4:", Value4)
	}

	fmt.Println("----- 無効を表す定数 -----")	

	{
		type SomeValue int

		const (
			Invalid SomeValue = iota
			Value1
			Value2
			Value3
			Value4
		)

		fmt.Println("Invalid:", Invalid)
		fmt.Println("Value1:", Value1)
		fmt.Println("Value2:", Value2)
		fmt.Println("Value3:", Value3)
		fmt.Println("Value4:", Value4)
	}
	
}
