package memory_test

import (

	"fmt"

)

func Main() {

	// ポインタ用変数
	n := 10

	fmt.Println("This package is pointer")
	PointerExample(10)
	PointerVar(&n)
}


// Pointer
func PointerExample(z int) {

	z = z + 1

	fmt.Println(z, "変数zに格納されている数字")

	// 変数zのアドレスを取得するには、(&)を使う。
	// pはzのアドレスを指す『ポインタ変数』
	// zのアドレスをを表示するには、pを出力すればよい。
	p := &z

	fmt.Println(p, "変数zのアドレス（メモリの格納場所）")

	fmt.Println(*p, "ポインタ変数の中身")

}


// Pointer Var
func PointerVar(y *int) {

	// この時点では変数
	fmt.Println(*y)
	*y = *y + 1

	fmt.Println(y)
	fmt.Println(*y)
}



// Goはポインタを扱う。
// ポインタの値はメモリアドレスを指す。
// 変数「T」のポインタは、*T型で、ゼロ値はnil。
func PointerTest(){

	i, j := 42, 2701

	// &オペレータは、そのオペランド(operand)へのポインタを引き出す。
	// 『オペランド』 -- 「1 + 2」の「+」が演算子、「1」、「2」がオペランド
	// ※式に登場する数値とか変数（値を入れる箱）がオペランド(被演算子)
	// ※つまり演算子じゃない方
	p := &i

	fmt.Println(*p)

	*p = 21

	fmt.Println(i)



	p = &j
	*p = *p /37
	fmt.Println(j)


}