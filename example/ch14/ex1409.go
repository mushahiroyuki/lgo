package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)


func MakeTimedFunction(f any) any { //liststart1
	rf := reflect.TypeOf(f)
	if rf.Kind() != reflect.Func {
		panic("関数を指定してください")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf,
		func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("%sの所要時間: %v\n", runtime.FuncForPC(vf.Pointer()).Name(),
			end.Sub(start))
		return out
	})
	return wrapperF.Interface()
} //listend1

func timeMe() { //liststart2
	time.Sleep(1 * time.Second)
}

func timeMeToo(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func())
	timed()
	timedToo := MakeTimedFunction(timeMeToo).(func(int) int)
	fmt.Println("timedToo(2):", timedToo(2))
} //listend2
