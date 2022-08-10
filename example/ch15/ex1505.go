package main

import (
	"fmt"
	"math"
)

func main() { //liststart5
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)  // {{1,1} {5,5}}

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2) // {{10,10,10} {11,5,0}}

	// 同じ型のペアが指定されていない場合
	// closer3 := FindCloser(pair2Da, pair3Da) // コンパイル時のエラー
	// fmt.Println(closer3)

	// Differを実装していない場合  コンパイル時のエラー
	// closer4 := FindCloser(Pair[StringerString]{"a", "b"},
	//             Pair[StringerString]{"hello", "goodbye"})
	// fmt.Println(closer4)
} //listend5

// 同じ型の任意の2つ値を保持する型 //liststart1
type Pair[T fmt.Stringer] struct { // Tはfmt.Stringerを実装する
	Val1 T
	Val2 T
} //listend1

// 型パラメータをもつインタフェース //liststart2
type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
} //listend2

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] { //liststart3
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
} //listend3


type Point2D struct { //liststart4
	X, Y int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}

func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
} //listend4

type StringerString string

func (ss StringerString) String() string {
	return string(ss)
}

