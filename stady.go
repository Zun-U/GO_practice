// go run <プログラムファイル名> でビルド実行

// GOはmainパッケージから始まる
// [所属パッケージ]
package main

//パッケージのインポート
import (

	// 利用パッケージ名
	"fmt"

)

// 関数の外でvarによる変数宣言可能
var outF string = "out of function"

//※省略宣言は関数外でできない
//※ただし、関数無いで省略宣言をした場合、別関数で呼び出せる。
func example() {
	ss := "example function";
	fmt.Println(ss)
}


// func 関数名
func main() {

	// 変数var
	var s string = "Golang Test"

	// 変数の一括宣言
	var (
		n string = "n"
		u int = 6
		m float64 = 2.1212
		t,f bool = true, false
	)

	// 変数の省略宣言
	var count int = 5
	count_  := 5

	fmt.Println(s)
	fmt.Println(n, u, m ,t ,f)
	fmt.Println("通常の変数宣言",count,"変数の省略",count_)
	fmt.Println(outF)
	example()
}
