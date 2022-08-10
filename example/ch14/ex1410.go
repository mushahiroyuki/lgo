package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type outExp struct {
	out    []reflect.Value
	expiry time.Time
}

// buildInStruct は動的な構造体を生成する。
// フィールドはメモ化された関数の入力引数に対応する
// メモ化（memoization）についてはウィキペディアなどを参照
func buildInStruct(ft reflect.Type) (reflect.Type, error) {
	if ft.NumIn() == 0 {
		return nil, errors.New("must have at least one param")
	}
	// 動的な構造体を生成するために、reflect.StructFieldのスライスを生成する
	sf := make([]reflect.StructField, 0, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		ct := ft.In(i)
		// この構造体がマップのキーとして使われるので、比較可能（comparable）でなければならない
		// 構造体が比較可能であるためには、すべてのフィールドが比較可能でなければならない
		if !ct.Comparable() {
			return nil, fmt.Errorf("parameter %d of type %s and kind %v is not comparable", i+1, ct.Name(), ct.Kind())
		}
		// 入力引数に対してsfに構造体フィールドを追加する。
		// 入力引数の型を使って名前を作る
		sf = append(sf, reflect.StructField{
			Name: fmt.Sprintf("F%d", i),
			Type: ct,
		})
	}
	// 構造体のフィールドから動的な構造体の型を生成する
	s := reflect.StructOf(sf)
	return s, nil
}

// Memoizer は関数を引数として結果をキャッシュするラッパー関数を戻す
// expirationに指定した時間だけ有効になる
// メモ化（memoization）についてはウィキペディアなどを参照
//
// 関数fには以下のような制限がある
// 1. 関数の実行時間が長い（そうでなければキャッシュする意味がない）
// 2. 副作用をもたない（キャッシュした結果が無効になってしまうかもしれない）
// 3. 関数の入力引数は比較可能でなければならない
func Memoizer(f any, expiration time.Duration) (any, error) {
	// expiration -- 制限時間
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return nil, errors.New("only for functions")
	}

	// 関数の入力引数を表現するのに動的な構造体型を使う
	inType, err := buildInStruct(ft)
	if err != nil {
		return nil, err
	}

	if ft.NumOut() == 0 {
		return nil, errors.New("must have at least one returned value")
	}

	m := map[any]outExp{}
	fv := reflect.ValueOf(f)

	// 関数reflect.MakeFuncを使って、指定の関数と同じ入力と出力引数をもつようにする
	memo := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		// マップのキーを生成
		iv := reflect.New(inType).Elem()
		for k, v := range args {
			iv.Field(k).Set(v)
		}
		ivv := iv.Interface()

		// キーがマップにあるかどうか、制限時間が経過していないかをチェック
		ov, ok := m[ivv]
		now := time.Now()
		if !ok || ov.expiry.Before(now) {
			// マップ内にキーがなければ、あるいは制限時間が経過していれば
			// 関数を実行して結果をマップにキャッシュする
			ov.out = fv.Call(args)
 			ov.expiry = now.Add(expiration)
			m[ivv] = ov
		}
		// キャッシュ内の値を戻す
		return ov.out
	})
	// メモ化した関数を戻す
	return memo.Interface(), nil
}

func AddSlowly(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}

func main() {
	addSlowlyInterface, err := Memoizer(AddSlowly, 2*time.Second)
	if err != nil {
		panic(err)
	}
	// 型アサーションを使って戻されたany（interface{}）から関数を戻す
	addSlowly := addSlowlyInterface.(func(int, int) int)
	for i := 0; i < 5; i++ {
		start := time.Now()
		result := addSlowly(1, 2)
		end := time.Now()
		fmt.Printf("結果（%d）：%d  所要時間：%v\n", i, result, end.Sub(start))
	}
	time.Sleep(3 * time.Second)
	start := time.Now()
	result := addSlowly(1, 2)
	end := time.Now()
	fmt.Printf("結果：%d  所要時間：%v（制限時間経過後）\n", result, end.Sub(start))
}
