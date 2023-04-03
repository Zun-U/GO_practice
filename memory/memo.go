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
	PremsArr()
	Slicer()
	SL()
	SlaR()
	SliDef()
	SLaC()
	NilS()
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

// Arrays
func PremsArr() {

	// [n]T型は、型Tのn個の変数の配列(array)を表します。
	// ここでは、[2]string型は、string型の2個の変数の配列を表す。
	var a [2]string
	a[0] = "hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)


	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[2])
	fmt.Println(primes)


	// 配列の長さは、型の一部のため、配列のサイズを変えることは出来ない。

}

// Slice（スライス）
// 配列は『固定長』。
// スライスは『可変長』。
// 型[]TはT型のスライスを表す。
func Slicer() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// コロンで区切られた二つのインデックスlowとhighの境界を指定することによってスライスが形成される。
	// a[low : high]
	// 最初の要素は含み、最後の要素を除いた半開区間を選択する。
	var s []int = primes[1:4] // 3,5,7

	// インデックス番号1～4で、1は含み、4は含まない。
	// つまりは1，2，3のインデックス番号を選択している。

	fmt.Println(s)

}

// Slices are like reference to
func SlaR() {

	names := [4]string{
		"ringo",
		"mikan",
		"banana",
		"melon",
	}

	fmt.Println(names)

	// スライスは配列の参照の様なもの。
	// スライスはどんなデータも『格納しておらず』、単に元の配列の『部分列を指し』ているだけ。
	a := names[0:2] //a[0] = "ringo", a[1] = "mikan"
	b := names[1:3] //b[0] = "mikan", b[2] = "banana"
	fmt.Println(a, b)

	// スライスの要素を変更すると、そのもととなる配列の対応する要素が変更される。
	b[0] = "XXX" // before change, b[0] = "mikan"


	fmt.Println(a, b)

	// 同じ元となる配列を共有している他のスライスは、それらの変更が反映される。
	fmt.Println(names)
}



//　❐❐ Slice literals
func SL() {

	// これは配列のリテラル。
	exArr := [3]bool{true, true, true}

	// これは、上記と同様の配列を作成し、それを参照するスライスを作成する。
	// スライスのリテラルは『長さのない配列リテラル』の様なもの。
	exSli := []bool{false, false, false}

	fmt.Println(exArr, exSli)

	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)


	s := []struct {
		n int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

}



// Slice default
func SliDef() {

	s := []int{22, 33, 55, 77, 111, 131}

	s = s[1:4]
	fmt.Println(s) //s[0] = 33, s[1] = 55, s[2] = 77

	s = s[:2]
	fmt.Println(s) //s[0] = 33, s[1] = 55

	s = s[1:]
  fmt.Println(s) //s[1] = 55


	// スライスの既定値は、下限が0、上限はスライスの長さ。
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// この場合は、下限が「0」、上限が「10」
	fmt.Println(a[0:10])
	fmt.Println(a[:10])
	fmt.Println(a[0:])
	fmt.Println(a[:])

}


// Slice lenghth and capacity
func SLaC() {

	// スライスは長さ(length)と容量(capacity)の両方を持っている。
	// スライスの長さは、持っている要素の『数』
	// スライスの容量は、『スライスの最初から数えて』、元となる配列の要素数。
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) //len = 6, cap = 6

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s) // len = 0, cap = 6

	// Extends its length
	s = s[:4]
	printSlice(s) // len = 4, cap = 6

	// Drop its first two values
	s = s[2:]
	printSlice(s) //len = 2, cap = 4


}

func printSlice(s []int) {

	// スライス"s"の長さと容量はlen(s)とcap(s)という式を利用して得ることができる。
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}




// Nil slice
func NilS() {

	// スライスのゼロ値は『nil』。
	// nilスライスは『0』の長さと容量を持ち、何の元となる配列を持っていない。
	var c []int
	fmt.Println(c, len(c), cap(c))


	if c == nil {
		fmt.Println("nil")
	}

}