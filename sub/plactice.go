package sub

// 括弧でまとめて宣言を行うことをfactoredステートメントという（因数分解宣言？）
import (
	"fmt"
	"math"
	"math/cmplx"
)

func HelloGo() {
	fmt.Println("this Folder is sub")
	fmt.Println(math.Pi)
	fmt.Println(add(5, 6))
	fmt.Println(addPlus(11, 12))
	fmt.Println(maltFunctions(3,"maltiple strings"))
	fmt.Println(namedReturn(10))
	fmt.Println(namedNotReturn(5))
	fmt.Println(res, tes)
	fmt.Println(ko, jo, io, to)
	fmt.Println(i, j)
	fmt.Println(t, u)
	fmt.Println(sd ,ed ,fd)
	fmt.Println(shortVarDec(1000, true))
	fmt.Printf("Type: %T Value :%v\n", boolin, boolin)
	fmt.Printf("Type: %T Value :%v\n", b1, b1)
	fmt.Printf("Type: %T Value :%v\n", f1, f1)
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}


// 変数名の『後ろ』に型を指定する。
// 理由は右から左に読めるように
func add(xa int, ya int) int {
	return xa + ya
}


// 複数の引数において、型がどれも同じだった場合、最後の型を残して他を省略できる。
func addPlus(xz, yz int) int {
	return xz + yz
}


// 関数の『戻り値に』複数の型を指定できる。
// ここでは、２個目の括弧(int, string)が関数の戻り値を複数指定している。
func maltFunctions(s int, d string) (int, string) {
	return s, d
}


// named return value
// (q, w int)のように、戻り値に名前を付けることができる。
// そうすると、関数で最初に定義した変数名として『扱われる』。
func namedReturn(sum int) (q, w int){

	// この関数で、最初に定義した変数
	// これが、戻り値として扱われる。
	q = sum * 2
	w = sum + q

	// naked return
	// 『名前を付けた戻り値』を使うと、returnに何も指定しなくても良い
	return

	// ※※※ ただし、この関数の様な短い関数でのみ使うべき。 ※※※

}

// ❐❐　上記のコードの通常バージョン　❐❐
func namedNotReturn(sum int) (int, int) {


	// 変数の宣言は２種類ある
	var qa int = sum * 2
	// ※省略宣言は、関数のスコープ内限定で宣言可能
	wa := sum + qa

	return qa, wa

}


// 関数外で『var』を使った変数の宣言が可能

var (
	res string = "関数外で変数宣言"
	tes string = "複数の変数宣言を一括"
)

// いろいろな宣言方法
// 一括宣言、最後にまとめて型指定
var ko, jo, io, to bool


// 初期化子
// ここでは『1』と『2』が初期化子
var i, j int = 1, 2

// 初期化子が与えられている変数は、型省略できる。
// 変数『t』、『u』の型は初期化子の型になる。ここでは、『3』、『4』がint型なので、自動でint型になる。
var t, u = 3, 4


// 型が異なる初期化子を付与できる。
// この場合は、順番にint、string、bool型となっている。
var sd, ed, fd = 7, "三原色", true


// 暗黙的な変数宣言
func shortVarDec(sdg int, vdg bool) (int, bool) {

	// varを省略して『:=』で変数宣言する。
	// 関数【内】でのみ暗黙的な変数宣言ができる
	// 初期化子が『1』なので、型省略されて変数「gos」は自動でint型となる。
	gos := 1

	// 複数の暗黙的な変数宣言ができる。
	ges, gus := 10, 100

	ris := gos + ges + gus +sdg

	return ris, vdg

}

// Basic Types （Goの基本的な型）

var (

	boolin bool = true

	stringin string = "文字列"

	// int型はint8、int16、int32、int64と細かく指定できる
	a1 int = 1

	// 符号なし整数型。uint8、16、32、64
	b1 uint = 1<<64 - 1

	// uint8の別名
	c1 byte = 3

	// int32の別名。Unicodeのコードポイントを表す。
	d1 rune = 4

	// float32、64
	e1 float32 = 1.2345

	// complex64、128
	f1 complex128 = cmplx.Sqrt(-5 + 12i)
)


// Zero values
func ZV() {

	// 初期値を与えずに変数宣言すると『ゼロ値』が与えられる。

	// 数値型のゼロ値は『0』
	var i int
	var f float64

	// 真偽値型のゼロ値は『false』
	var b bool

	// 文字列型のゼロ値は『』（空文字：empty string）
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}


// ★☆★☆★☆★☆★☆★☆
//  ❐❐❐　型変換　❐❐❐
// ★☆★☆★☆★☆★☆★☆
func TypeConversion() {

	var first int = 42
	// Goでは明示的な型変換を行う。(float64(first)の「float64」の部分が明示的に行われている部分)
	var second float64 = float64(first)
	var third uint = uint(second)


	// Goでは通常は文字列から数値型、数値型から文字列に変換できない。

	fmt.Printf("Type: %T Value :%v\n", first, first)
	fmt.Printf("Type: %T Value :%v\n", second,second)
	fmt.Printf("Type: %T Value :%v\n", third, third)
}


// 型変換のシンプルな記述
func TypeConversionSimple() {

	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Println(x, y, z)

}



// 暗黙的な変数宣言の型指定
func TypeInterfaceJob() {

	// 数字ならint型
	v := 123

	// 小数ならfloat型
	s := 1.23

	// 複素数complex128型 ※float64の実数部と、虚数を持つ複素数
	c := 0.867 + 0.5i

	// ダブルクォテーション("")で囲まれた値ならstring型
	d := "文字列型だよ"

	fmt.Printf("Type: %T Value :%v\n", v, v)
	fmt.Printf("Type: %T Value :%v\n", s, s)
	fmt.Printf("Type: %T Value :%v\n", c, c)
	fmt.Printf("Type: %T Value :%v\n", d, d)

}



// 定数(constant)
const Pis = 3.14

func ClCons() {

	// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使える。
	const World = "世界"

	const Truth = true

	fmt.Println("This is ", Pis)
	fmt.Println("ここは異", World, "です")
	fmt.Println("ああ、その通り。", Truth)
}



// Numeric Constants
// 数値の定数は、高精度な値(values)

const (
	big = 1 << 100
	small =big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

