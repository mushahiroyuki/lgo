package main

import (
	"errors"
	"fmt"
)

type treeNode struct {
	val    treeVal
	lchild *treeNode
	rchild *treeNode
}

// treeValはエクスポートされていないマーカーのインタフェース定義する
// treeNodeのvalに代入できるのがどの型かをはっきりさせる
type treeVal interface {
	isToken()
}

type number int

func (number) isToken() {}  // 数値が対象

type operator func(int, int) int // int2つを受け取りintを返す関数

func (operator) isToken() {} // オペレータが対象

func (o operator) process(n1, n2 int) int { // 演算を行う
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func walkTree(t *treeNode) (int, error) { //liststart
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("不正な式")
	case number:
		// t.valがnumber型だとわかったので
		// int型の値を返す
		return int(val), nil
	case operator:
		// t.valがoperator型だとわかったので
		// 左右の子の値を取得し
		// operatorのメソッドprocess()を呼び出し
		// 値を処理した結果を返す
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		// treeValに新しい型が定義されたが、
		// walkTreeは更新されていないことがわかる
		return 0, errors.New("不明な型のノード")
	}
} //listend


func main() {
	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err) //  70 <nil>

	parseTree, err = parse("2-10*3")
	if err != nil {
		panic(err)
	}
	result, err = walkTree(parseTree)
	fmt.Println(result, err) // -28 <nil>
}

func parse(s string) (*treeNode, error) {
	// ここではこのメソッドには焦点があたっていないので、手動（ハードコード）で返す
	if s == "5*10+20" { // -> [+ [* 5 10] 20]
		return &treeNode{
			val: operators["+"],
			lchild: &treeNode{
				val:    operators["*"],
				lchild: &treeNode{val: number(5)},
				rchild: &treeNode{val: number(10)},
			},
			rchild: &treeNode{val: number(20)},
		}, nil
	}	else if s == "2-10*3" { // -> [- 2 [* 10 3]]
		return &treeNode{
			val: operators["-"],
			lchild: &treeNode{val: number(2)},
			rchild: &treeNode{
				val:  operators["*"],
				lchild: &treeNode{val: number(10)},
				rchild: &treeNode{val: number(3)},
			},
		}, nil	
	}
	
	return nil, nil
}
