// パッケージfmtの説明用コード。実行するには 「go run main.go」
package main    // メインプログラムは必ずこうする

import "fmt"    // 読み込むパッケージ（ライブラリ）の宣言

func main() {   // 関数の定義。必ずmainが必要。「{」の前で改行してはダメ！
	i := 3      // 変数iに3が代入される。型は自動判定されるこの場合はint（整数）
	answer := "Go言語"  // "Go言語"は文字列なので、answerの型はstring（文字列）
	fmt.Printf("第%d問の答えは「%s」です。\n", i, answer)     // ←1:
	// %dがiに、%sがanswerに対応

	s1 := "文字列"
	fmt.Printf("%s\n", s1)                                  // ←2:
	s2 := `バッククオートを使った、改行や
	タブが入った
		文字列`
	fmt.Printf("%s\n", s2)  // そのまま表示                   // ←3〜5:
	fmt.Printf("%q\n", s2)  // 「エスケープ」される            // ←6:

	i = 254  // 上でiの型はintになっているので、254は代入できる
	fmt.Printf("d: %d  b: %b  o: %o  x: %x  X: %X\n", i, i, i, i, i)  // ←7:

	f := 3.14159265358      // 小数（型はfloat64）
	fmt.Printf("f: %f  F: %F\ne: %e  E: %E\n", f, f, f, f)  // ←8〜9:
	fmt.Printf("v: %v  %v  %v\n", s1, i, f)                 // ←10:
	fmt.Printf("T: %T  %T  %T\n", s1, i, f)                 // ←11:
	fmt.Printf("「%%%%」で「%%」をひとつ出力\n")               // ←12:

	fmt.Printf("%%fでs1を指定: %f\n", s1)     // ←13:（間違いあり！）
}
