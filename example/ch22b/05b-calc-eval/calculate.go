package main

import (
  "go/token" // Evalに必要
  "go/types" // 関数Eval（文字列を評価: EVALuateする）のため
	"fmt"
)

func calculate(exp string) (string, error) {
	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, exp)
	if err != nil {
		fmt.Println("計算できません")
		return "", err
	}
	r := fmt.Sprintf("%s", tv.Value)
	return r, err
}
