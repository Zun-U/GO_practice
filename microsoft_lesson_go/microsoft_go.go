package main

import (
	"fmt"
)


func main() {

	fmt.Println("This is the Microsoft's Go langage course.")




	Variable()




}

func Variable() {

	// 変数の宣言（☆宣言だけ）
	// 『var』キーワードを使用する
	var firstName string // string型の「firstName」という名前の変数宣言

	// "同じ型"であればコンマ(,)を用いて同一行に複数宣言可能
	var secondName, lastName string

	// 異なる方は異なる行に宣言
	var age int

	fmt.Printf("FN:%v SN:%v LN:%v AGE:%v", firstName, secondName, lastName, age)

}

func VariableBlock() {

	// ()を使用して『変数宣言用ブロック』の様に記述できる
	var(
		fName, sName, lName string
		ages int
	)

}

func VariableInit() {

	// ❐ 変数の初期化 ❐
	var (
		FName string = "John"
		LName string = "Doe"
		Age int = 32
	)

}

func VariableInitType() {

	// 変数を初期化する場合、値を指定して初期化するときにGoによって型が推定されるため、指定する必要がない
	var (
		FName = "John" // Goがstring型と推定
		LName = "Doe"  // Goがstring型と推定
		Age = 32			 // Goがint型と推定
	)

}