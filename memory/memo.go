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
	fmt.Println(Vertex{1, 2})
	AccessStruct()
	fmt.Println(v1, p, v2, v3)
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

	// ポインタpを通してiから値を読みだす。
	fmt.Println(*p)

	// ポインタpを通してiへ代入する。
	*p = 21

	// これらを"dereferencing" または "indirecting"として知られている。

	fmt.Println(i)



	p = &j
	*p = *p /37
	fmt.Println(j)


}


// Structs（構造体）
// Struct（構造体）はfield(フィールドの集まり)
type Vertex struct {
	X int
	Y int
}

// structフィールドは、ドット(.)を用いてアクセスします。
func AccessStruct() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// Struct Literals
type Zertex struct {
	X, Y int
}


// リテラル(literal)
// ソースコードに直接記載する値
// ・整数リテラル
// ・浮動小数点リテラル
// ・イマジナリリテラル
// ・ルーンリテラル
// ・文字列リテラル

// structの初期値の『割り当て』を示している。
var (

	// Zertex型
	v1 = Zertex{1, 2}

	// 「X: 1」は、Xだけを初期化している。
	// 「Y: 0」 が暗黙的に宣言されている。
	v2 = Zertex{X: 1}

	// 「X: 0」と「Y: 0」が暗黙的に宣言されている。
	v3 = Zertex{}

	// *Zertex型。
	// structへのポインタを戻す。
	p = &Zertex{1, 2}

)