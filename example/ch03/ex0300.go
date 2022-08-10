// 3章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。

package main

import "fmt"

func main() {
	fmt.Println("===== 3.1　配列 =====")
	{
		var x [3]int
		fmt.Println(x)
	}
	{
		var x = [3]int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	fmt.Println("----- 配列の比較 -----")
	{
		var x = [...]int{1, 2, 3}
		var y = [3]int{1, 2, 3}
		fmt.Println(x == y) // true が出力される
	}

	fmt.Println("----- 多次元配列のシミュレート -----")
	{
		var x [2][3]int
		fmt.Println(x)
	}

	fmt.Println("----- インデックス（添字）の指定 -----")
	{
		var x [3]int
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])

		var y [2][3]int
		y[0][2] = 12
		y[1][1] = 3
	}

	fmt.Println("----- len -----")
	{
		var x [3]int
		fmt.Println("len(x):", len(x))

		var y [2][3]int
		fmt.Println("len(y):", len(y))
	}

	fmt.Println("===== 3.2　スライス =====")
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}

	{
		var x []int
		fmt.Println(x)
	}

	fmt.Println("----- スライスの比較 -----")
	{
		var x []int
		var y []int
		// fmt.Println(x==y) // エラー
		fmt.Println("x == nil:", x == nil) // true
		fmt.Println("y != nil:", y != nil) // false
	}

	fmt.Println("===== 3.2.1 len =====")
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len(y):", len(y)) // 0
	}

	fmt.Println("----- スライスのスライス -----")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])    // [2 3]
		fmt.Println(y[1][1]) // 3
	}

	fmt.Println("----- lenの引数 -----")
	{
		var x [][]int
		fmt.Println("len(x):", len(x))
		var y = []int{2, 3, 4, 5}
		fmt.Println("len(y):", len(y))
		var z string = "abc"
		fmt.Println("len(z):", len(z))

		var a = 12
		fmt.Println("a:", a)
		// fmt.Println(len(a)) // エラー
	}

	fmt.Println("===== 3.2.2　append =====")

	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println("append(x, 10):", x)
	}

	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println("append(x, 4):", x)
		x = append(x, 5, 6, 7)
		fmt.Println("append(x, 5, 6, 7):", x)

		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println("x:", x)
		fmt.Println("append(x, y...):", append(x, y...))

		// append(x, y) // エラー
	}

	fmt.Println("===== 3.2.4　make =====")

	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}

	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [0 0 0 0 0] 5 5
		x = append(x, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [0 0 0 0 0 10] 6 10
	}

	fmt.Println("----- キャパシティ（第3引数）を指定 -----")
	{
		x := make([]int, 5, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	{
		x := make([]int, 0, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [] 0 10
	}

	{
		x := make([]int, 0, 10)
		x = append(x, 5, 6, 7, 8)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	fmt.Println("===== 3.2.5　スライスの生成方法の選択 =====")
	{
		var data []int
		fmt.Println(data, len(data), cap(data))
		fmt.Println("data == nil:", data == nil)

		var x = []int{}
		fmt.Println(x, len(x), cap(x))
		fmt.Println("x == nil:", x == nil)
	}

	{
		data := []int{2, 4, 6, 8}
		fmt.Println(data, len(data), cap(data))
	}

	fmt.Println("===== 3.3　文字列、rune、バイト =====")
	{
		var s string = "Hello there"
		var b byte = s[6]
		fmt.Println(b)                  // 116
		fmt.Printf("printfで指定：%c\n", b) // 「t」。フォーマット指定で%c（文字）を指定すると「t」を出力
		var c rune = rune(s[6])         // 116
		fmt.Println(c)
		fmt.Printf("printfで指定：%c\n", c) // 「t」
	}

	fmt.Println("----- 文字列に対する「スライス式」 -----")
	{
		var s string = "Hello there"
		// 01234567890
		fmt.Println("s:", s)   // Hello there
		var s2 string = s[4:7] // o t
		var s3 string = s[:5]  // Hello
		var s4 string = s[6:]  // there
		fmt.Printf("s2:%s\ns3:%s\ns4:%s\n", s2, s3, s4)
	}

	fmt.Println("----- 絵文字 -----")
	{
		var s string = "Hello ☀"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println(s)  // Hello ☀
		fmt.Println(s2) // o ?（文字化け）
		fmt.Println(s3) // Hello
		fmt.Println(s4) // ☀
		fmt.Println("len(s):", len(s))
	}

	fmt.Println("----- 日本語 -----")
	{
		var s string = "こんにちは"
		fmt.Println("s:", s)
		fmt.Println("s[0]:", s[0])
		var b byte = s[6]
		fmt.Println("b:", b)
		var c rune = rune(s[6])
		fmt.Println("c:", c)
	}

	{
		var s string = "こんにちは、みなさん"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println("s:", s)
		fmt.Println("s2:", s2)
		fmt.Println("s3:", s3)
		fmt.Println("s4:", s4)
	}

	{
		var a rune = 'x'
		var s string = string(a)
		var b byte = 'y'
		var s2 string = string(b)
		fmt.Println("a, s, b, s2:", a, s, b, s2)
	}

	{
		var x int = 65
		var y = string(x)
		fmt.Println("y:", y) // A （65ではない）
	}

	fmt.Println("===== 3.4　マップ =====")
	{
		var nilMap map[string]int
		fmt.Println("nilMap == nil:", nilMap == nil)   // true
		fmt.Println("len(nilMap):", len(nilMap))       // 0
		fmt.Println("nilMap[\"abc\"]:", nilMap["abc"]) // 0
		// nilMap["abc"] = 3 // ←パニックになる

		totalWins := map[string]int{}                      // マップリテラル。string→intのマップを要素なしで初期化。nilではない
		fmt.Println("totalWins == nil:", totalWins == nil) // false

		fmt.Println("totalWins[\"abc\"]:", totalWins["abc"]) //		 -> 0
		totalWins["abc"] = 3                                 // 書き込みもできる
		fmt.Println("totalWins[\"abc\"]:", totalWins["abc"]) //
	}
	{
		totalWins := map[string]int{
			"セネターズ":  14,
			"スパローズ":  15,
			"ファルコンズ": 22,
		}
		fmt.Println(totalWins)
	}
	{
		teams := map[string][]string{
			"ライターズ":    []string{"夏目", "森", "国木田"},
			"ナイツ":      []string{"武田", "徳川", "明智"},
			"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
		}

		fmt.Println(teams)          // map[ミュージシャンズ:[ラベル ベートーベン リスト] ライターズ:[夏目 森 国木田] ロビンズ:[池長 渕田 山村]]
		fmt.Println(teams["ライターズ"]) // [夏目 森 国木田]
		fmt.Println(teams["ロビンズ"])  // [池長 渕田 山村]

		/*
			teams := map[string][]string {
				"Orcas": []string{"Fred", "Ralph", "Bijou"},
				"Lions": []string{"Sarah", "Peter", "Billie"},
				"Kittens": []string{"Waldo", "Raul", "Ze"},
			}

			fmt.Println(teams) // map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
			fmt.Println(teams["Lions"]) // [Sarah Peter Billie]
			fmt.Println(teams["Kittens"]) // [Waldo Raul Ze]
		*/

		teams2 := map[string][]string{
			"シャチチーム":  []string{"謙信", "信長", "家康"},
			"ライオンチーム": []string{"レオ", "たか子", "カナ"},
			"猫チーム":    []string{"AKB", "MNB", "FNB"},
		}

		fmt.Println(teams2)           // map[シャチチーム:[謙信 信長 家康] ライオンチーム:[レオ たか子 カナ] 猫チーム:[AKB MNB FNB]]
		fmt.Println(teams2["シャチチーム"]) // [謙信 信長 家康]
		fmt.Println(teams2["チャチチーム"]) // []
		fmt.Println(teams2["猫チーム"])   // [AKB MNB FNB]

		fmt.Println("len(teams2[\"猫チーム\"]):", len(teams2["猫チーム"])) // 3
	}

	{
		ages := make(map[int][]string, 10)
		fmt.Println("ages:", ages)           // map[]
		fmt.Println("len(ages):", len(ages)) // 0
	}

	fmt.Println("===== 3.4.2　カンマOKイディオム =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok) // 5 true

		v, ok = m["world"]
		fmt.Println(v, ok) // 0 true

		v, ok = m["goodbye"]
		fmt.Println(v, ok) // 0 false

	}

	fmt.Println("===== 3.4.3 マップからの削除 =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		fmt.Println(m)
		v, ok := m["hello"]
		fmt.Println(v, ok)
		delete(m, "hello")
		fmt.Println(m)
		v, ok = m["hello"]
		fmt.Println(v, ok)
		delete(m, "こんにちは")
		fmt.Println(m)
	}

	fmt.Println("===== 3.5　構造体 =====")
	{
		type person struct {
			name string
			age  int
			pet  string
		}

		var fred person
		fmt.Println("fred:", fred) // fred: { 0 } // 空文字列2つあり

		bob := person{} // 構造体リテラル。全フィールドがゼロ値で初期化される

		fmt.Println("bob:", bob) // bob: { 0 } // 空文字列2つあり
		bob.name = "ボブ"
		fmt.Println("bob:", bob) // bob: {ボブ 0 }

		julia := person{
			"Julia", // name
			40,      // age
			"cat",   // pet
		}
		fmt.Println("julia:", julia) // julia: {Julia 40 cat}

		/*	  コンパイル時エラー（全フィールドを指定する必要がある）
		fred := person{
			"フレッド",  // name
			"cat",    // pet
		}
		*/

		beth := person{
			age:  30,
			name: "ベス",
		}
		fmt.Println("beth:", beth) // beth: {ベス 30 }

		fmt.Println("bob.name:", bob.name) // bob.name: ボブ
		fmt.Println("bob:", bob)           // bob: {ボブ 0 }

	}

	fmt.Println("===== 3.5.1　無名構造体 =====")
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println("person:", person) // person: {bob 50 dog}

	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}
	fmt.Println("pet:", pet) // pet: {ポチ dog}

	fmt.Println("===== 3.5.2　構造体の比較と変換 =====")

	{
		type firstPerson struct {
			name string
			age  int
		}

		type secondPerson struct {
			name string
			age  int
		}

		type thirdPerson struct {
			age  int
			name string
		}

		x1 := firstPerson{
			"太郎",
			24,
		}
		//		x1.name = "太郎"
		//		x1.age = 24
		fmt.Println("x1:", x1)

		var x2 secondPerson
		x2 = secondPerson(x1)
		fmt.Println("x2:", x2)
		//  fmt.Println("x1==x2:", x1==x2) // コンパイル時エラー
		x12 := firstPerson(x2)
		fmt.Println("x1==x12:", x1 == x12) // x1==x12: true

		var x3 thirdPerson
		// x3 = thirdPerson(x1)  //コンパイル時エラー
		fmt.Println("x3:", x3)

		type fourthPerson struct {
			firstName string
			age       int
		}
		var x4 fourthPerson
		// x4 = fourthPerson(x1)  // コンパイル時エラー
		fmt.Println("x4:", x4)

		type fifthPerson struct {
			name          string
			age           int
			favoriteColor string // 好みの色
		}
		var x5 fifthPerson
		// x5 = fifthPerson(x1)  // コンパイル時エラー
		fmt.Println("x5:", x5)
	}

	fmt.Println("\n----- 無名構造体との変換 -----")
	{
		type firstPerson struct {
			name string
			age  int
		}
		f := firstPerson{
			name: "Bob",
			age:  50,
		}
		var g struct {
			name string
			age  int
		}
		fmt.Println("f:", f)
		fmt.Println("g:", g)
		g = f
		fmt.Println("g（fの代入後）:", g)
		fmt.Println("f==g:", f == g)
	}

} // main
