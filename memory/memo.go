package memory_test

import (
	"fmt"
	"strings"
	"math"
)

func Main() {

	// ポインタ用変数
	n := 10

	fmt.Println("This package is pointer")
	PointerExample(10)
	PointerVar(&n)
	fmt.Println(Vertex{1, 2})
	StructPointer()
	AccessStruct()
	fmt.Println(v1, p, v2, v3)
	PremsArr()
	Slicer()
	SL()
	SlaR()
	SliDef()
	SLaC()
	NilS()
	CreateSlice()
	Slisli()
	ApndSli()
	PowErt()
	RanCont()
	Mapap()
	MapLit()
	litMapCont()
	Mapmut()
	MathFn()
	AdderCL()
	PrintZertexStruct()
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
func PointerTest() {

	i, j := 42, 2701

	// &オペレータは、そのオペランド(operand)へのポインタを引き出す。
	// 『オペランド』 -- 「1 + 2」の「+」が演算子、「1」、「2」がオペランド
	// ※式に登場する数値とか変数（値を入れる箱）がオペランド(被演算子)
	// ※つまり演算子じゃない方
	p := &i

	// ポインタpを通してiから値を読みだす。
	fmt.Println(*p, "『変数iのポインタ』")
	fmt.Println(p, "『変数iのアドレス』")
	// ポインタpを通してiへ代入する。
	*p = 21

	// これらを"dereferencing" または "indirecting"として知られている。

	fmt.Println(i, "『ポインタpを通して参照元の値が書き換わる』")

	p = &j
	*p = *p / 37
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
	fmt.Println(v.X, "『ドット(.)で構造体のフィールドにアクセスできる。Vertexの「X」にアクセス』")
}

// Pointer to struct
func StructPointer() {

	v := Vertex{1, 2}
	println(v.X, "『構造体Vartex』を変数vに格納。フィールドXにアクセス。』")

	p := &v
	println(p, "『構造体Vartexのフィールドへのメモリアクセス。』")

	p.X = 1e9
	fmt.Println(p.X, "『構造体VartextのフィールドXの値の書きかえ』")

	fmt.Println(v, "『構造体Vartextを呼び出したときに、中のXも書き換わっている。』")

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

func PrintZertexStruct() {
	fmt.Println(v1, "Zertex型")
	fmt.Println(v2, "ZertexのXだけ")
	fmt.Println(v3, "ZertexのX、Yが暗黙的に0と宣言されている")
	fmt.Println(p, "ポインタ")
}

// Arrays
func PremsArr() {

	// [n]T型は、型Tのn個の変数の配列(array)を表します。
	// ここでは、[2]string型は、string型の2個の変数の配列を表す。
	var a [2]string
	a[0] = "hello"
	a[1] = "World"

	fmt.Println(a[0], a[1], "『a[0]』=helloとa[1]のWorldを出力』")
	fmt.Println(a, "『配列a』に存在する[0][1]に格納されている文字列両方を出力。※その配列全体を出力。")

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[2], "『配列の長さは0から始まる。0,1,2,3...』")
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
	// ❐❐❐　a[low : high]　❐❐❐
	// 最初の要素は含み、最後の要素を除いた半開区間を選択する。
	var s []int = primes[1:4] // 3,5,7

	// インデックス番号1～4で、1は含み、4は含まない。
	// つまりは1，2，3のインデックス番号を選択している。

	fmt.Println(s, "『スライスsは配列primesのスライス。配列primesの[1][2][3]に格納されている値をスライスsが【参照】している。=> primes[1:4]』")

	// つまり、スライスは配列のポインタ変数。
}

// Slices are like reference to
func SlaR() {

	names := [4]string{
		"ringo",
		"mikan",
		"banana",
		"melon",
	}

	fmt.Println(names, "『配列names』")

	// スライスは配列の参照の様なもの。
	// スライスはどんなデータも『格納しておらず』、単に元の配列の『部分列を指し』ているだけ。
	a := names[0:2] //a[0] = "ringo", a[1] = "mikan"
	b := names[1:3] //b[0] = "mikan", b[2] = "banana"
	fmt.Println(a, b, "『スライスa => names[0:2]』、スライスb => names[1:3]』")

	// スライスの要素を変更すると、そのもととなる配列の対応する要素が変更される。
	b[0] = "XXX" // before change, b[0] = "mikan"

	fmt.Println(a, b, "スライスの要素が変更された >>> b[0]")

	// 同じ元となる配列を共有している他のスライスは、それらの変更が反映される。
	fmt.Println(names, "スライスの変更が元の配列に影響を与える")
}

// 　❐❐ Slice literals
func SL() {

	// これは配列のリテラル。
	exArr := [3]bool{true, true, true}

	// これは、上記と同様の配列を作成し、それを参照するスライスを作成する。
	// スライスのリテラルは『長さのない配列リテラル』の様なもの。
	exSli := []bool{false, false, false}

	fmt.Println(exArr,"<= [3]bool{true, true, true}", exSli, "<= []bool{false, false, false}")

	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q,"スライスqは配列q[0=>1, 1=>2, ...]の様な配列を作成し参照している。つまり自作自演")

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
	fmt.Println(s) //s[0] = 33, s[1] = 55, s[2] = 77  // 元の配列のindex番号1~3の値

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
	fmt.Println(c, ">>>スライスのゼロ値", len(c), ">>>len", cap(c), ">>>cap", "[スライスのゼロ値はnil。nilは0の長さと容量]")

	if c == nil {
		fmt.Println("nil")
	}

}

// Creating a slice with make
func CreateSlice() {

	// スライスは、組み込みの『make』関数を使用して作成可能。
	// 動的サイズの配列を作成する方法になる。
	// 『make』関数はゼロ化された配列を”割り当て”、その配列を指す『スライス』を”返す”。
	a := make([]int, 5)
	printSliceX("a", a)

	// makeの三番目の引数に、スライスの容量(capacity)を指定できる。
	// cap()で、スライスの容量を返す。
	b := make([]int, 0, 5)
	printSliceX("b", b)

	c := b[:5]
	printSliceX("c", c)

	d := c[2:5]
	printSliceX("d", d)

}

func printSliceX(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func Slisli() {

	// Create a tic-tic-toe board.
	// スライスは、『他のスライスを含む任意の型』を含むことができる。
	board := [][]string{
		// []string{"_", "_", "_"}, // sliceにおいて型が冗長化している。
		// []string{"_", "_", "_"},
		// []string{"_", "_", "_"},
		{"_", "_", "_"}, // 型宣言を省略してコンポジットリテラルを使用
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

// Appending to a slice
func ApndSli() {
	var s []int
	printSliceY(s)

	// append works on nil slices
	// Goの組み込みのappendは、要素をスライスの最後に追加する。
	// func append(slice [] Type , elems ... Type ) [] Type
	// appendの戻り値は、追加元のスライスと追加する変数群を併せたスライスになる。
	s = append(s, 0)
	printSliceY(s)

	// The slice grows as needed
	s = append(s, 1)
	printSliceY(s)

	//We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSliceY(s)
}

func printSliceY(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Range
var powop = []int{1, 2, 4, 8, 16, 32, 64, 128}

func PowErt() {

	// forループに利用するrangeは、スライスや、マップ(map)を"ひとつずつ反復処理する"ために使う。
	// スライスをrangeで繰り返す場合、rangeは反復ごとに２つの変数を返す。
	// 一つ目の変数はインデックス(index)で、二つ目はインデックスの場所の要素のコピー。
	for i, v := range powop {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// Range continued
func RanCont() {

	powo := make([]int, 10)

	for i := range powo {
		powo[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range powo {
		fmt.Printf("%d\n", value)
	}

}

// Maps

type Nertex struct {
	Lat, Long float64
}

// 『map』はキーと値を関連付ける。(マップする)
// mapのゼロ値はnil。
// nilマップはキーを持っておらず、またキーを追加することもできない。
var mmm map[string]Nertex

func Mapap() {

	// make関数は指定された型のマップを初期化し、使用可能な状態で返す。
	mmm = make(map[string]Nertex)
	mmm["Vell Labs"] = Nertex{
		40.68433, -74.39967,
	}
	fmt.Println(mmm["Vell Labs"])

}

// Map Literals
func MapLit() {

	type Varzexs struct {
		Lats, Longs float64
	}

	// mapリテラルは、structリテラルに似ていますが、 キー ( key )が必要。
	var oss = map[string]Varzexs{
		// "Bell Labs": Varzexs{ 40.688433, -74.39967,}, // MAP関数において型が冗長化している。Vartexsを省略。
		"Bell Lab": {40.688433, -74.39967},
		// "Google": Varzexs{37.42202, -122.08408,}, // MAP関数において型が冗長化している。Vartexsを省略。
		"Google": {37.42202, -122.08408},
	}

	fmt.Println(oss)
}



// Map literals continued
func litMapCont() {

	type Certex struct {
		Las, Lomd float64
	}


	// mapに渡すtop levelの型が単純な型名である場合、リテラルの要素から推定できるため
	// その型名を省略することができる。
	var m = map[string]Certex{
		"LabRatorY": {40.40404, -23.45678},
		"Golget": {1.23456789, 9.87654321},
	}

	fmt.Println(m)
}

// Mutating Maps
func Mapmut() {

	m := make(map[string]int)

	// map mの操作
	// m へ要素（elem）の挿入や更新
	// m[key] = elem
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// elem = m[key]
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の削除
	// delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかどうかは、２つ目の値で確認する。
	// elem, ok = m[key]
	v, ok := m["Answer"]

	// ※もしelemやokをまだ宣言していなければ、『:=』 で短く宣言できます。
	// elem, ok := m[key]

	// もし m にkeyがあれば、変数okはtrueとなり、存在しなければ変数okはfalseとなる。
	fmt.Println("The value:", v, "Present?", ok)

	// なお、mapにkeyが存在しない場合、elemはmapの要素の型のゼロ値となる。

}


// Function Values
func Compute(fn func(float64, float64) float64) float64{
	return fn(3, 4)
}

func MathFn() {

	// 関数も変数
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// 関数値(function value)は、関数の引数にとることもできる。
	// 戻り値としてもりようできる。
	fmt.Println(hypot(5, 12))

	fmt.Println(Compute(hypot), "変数hypotには関数値が格納されている")
	fmt.Println(Compute(math.Pow))

}



// Func closures
// ❐❐❐　Goの関数は クロージャ( closure ) です。　❐❐❐
// クロージャは、それ自身の外部から変数を参照する関数値です。
// この関数は、参照された変数へアクセスして変えることができ、その意味では、その関数は変数へ"バインド"( bind )されています。

// 例えば、 このadder 関数はクロージャを返しています。 各クロージャは、それ自身の sum 変数へバインドされます。
func adder() func(int) int {

	sum := 0

	// 無名関数(名前を持たない関数)
	// 関数を変数に代入するために使用されることが多い。
	return func(x int) int{
		sum += x
		return sum
	}

}

func AdderCL() {

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {

		fmt.Println(pos(i), neg(-2*i), i,"巡目")

	}

}