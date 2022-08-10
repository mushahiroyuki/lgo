package main

import(
	"fmt"
	"reflect"
)

//liststart1
type Foo struct { // 型（type）の定義
	A int
	B string
}

var x Foo // 変数の定義

func DoSomething(f Foo) { // 関数の定義
	fmt.Println(f.A, f.B)
} //listend1


func main() {

	fmt.Println("===== 14.1 リフレクションを使った実行時の型の処理 =====")
	
	fmt.Println("----- TypeOf -----")
	{
		v := 0
		w := 0.1
	
		vType := reflect.TypeOf(v)
		fmt.Println(vType)  // int
		wType := reflect.TypeOf(w)
		fmt.Println(wType) // float64
	}


	fmt.Println("----- ValueOf -----")
	{
		v := 1.03
		vValue := reflect.ValueOf(v)
		fmt.Println(vValue) // 1.03
	}

	{
		s := []string{"a", "b", "c"} //liststart4
		sv := reflect.ValueOf(s)  // svの型はreflect.Value
		s2 := sv.Interface().([]string) // s2の型は[]string //listend4
		fmt.Println(sv) // [a b c]
		fmt.Printf("%T\n", sv)  // reflect.Value
		fmt.Println(s2) // [a b c]
		fmt.Printf("%T\n", s2)  // []string
	}
	
	fmt.Println("----- 値 -----")
	{
		v := 1.03		
		vValue := reflect.ValueOf(v)
		fmt.Println("vValue:", vValue)
		
	}


	{
		i := 10 //liststart5
		iv := reflect.ValueOf(&i)  //listend5
//liststart6
		ivv := iv.Elem()
//listend6

		ivv.SetInt(20) //liststart7
		fmt.Println("i:", i) // 20   //listend7
		
	}
	
}
