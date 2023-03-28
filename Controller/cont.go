package controller_test

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func Controllers_yes() {

	fmt.Println(ForCont(10))
	fmt.Println(ForContinu(10))
	fmt.Println(ForToWhile(10))
	fmt.Println(MathIf(2), MathIf(-3))
	fmt.Println(Pow(3, 2, 10), Pow(3, 3, 10))
	fmt.Println(PowPow(3, 2, 10), PowPow(3, 3, 10))
	fmt.Println(LoopFunc(2))
	SwitchFuncs()
	SwitchOrder()
	SwitchTrue()
	Defdif()
	stadef()

}

func ForCont(x int) int {

	sum := 0

	// GoのFor文は３つのパーツから成り立つ
	// 1.初期化ステートメント　i := 0;
	// 2.条件式 i < 10;
	// 3.後処理ステートメント i++
	for i := 0; i < 10; i++ {
		sum += i * x
	}

	// 初期化ステートメント：最初のイテレーション（繰り返し）の前に初期化が実行されます。
	// 条件式：イテレーション（繰り返し）毎に評価されます。
	// 後処理ステートメント：イテレーション（繰り返し）毎の最後に実行されます。

	// 初期化ステートメントは。短い変数宣言によく利用します。
	// その変数はforステートメントのスコープ内で『のみ』有効です。

	// ループは、条件式の評価がfalseとなった場合にイテレーションを停止します。

	// Goのforステートメントは括弧()は必要なく、中括弧{}は常に必要。


	return sum

}


// for文の初期化と後処理ステートメントの記述は任意。
func ForContinu(x int) int {

	sum := 1

	for ; sum < 1000; {
		sum += sum * x
		fmt.Println(sum)
	}

	return sum

}


// Goのwhile文
func ForToWhile(x int) int {

	sum := 1

	// セミコロンを省略することができる。
	// Goではforだけ表現する。
	for sum < 1000 {
		sum = sum * x + x
	}

	return sum

}


// Forever
func ForForever() {
	// 条件式を省略したfor文は無限ループになる。
	// for {
	//
	// }
}


// if statement
func MathIf(x float64) string {

	if x < 0 {
		return MathIf(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))

}



// Ifの簡単なステートメント
func Pow(x, n, lim float64) float64 {


	// for文のように、条件式の『前に』、評価するためのステートメントを書くことができる。
	// ここで宣言された変数（v := mathPow(x, n)）は、ifのスコープ内だけで有効。
	if v:= math.Pow(x, n); v < lim {
		return v
	}

	// スコープの外なので、ifステートメントで宣言された変数は使えない。（ここでは変数vのこと）
	return lim

}



// if else
func PowPow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {

		// ifステートメントで宣言された変数はelseでも使うことができる。
		fmt.Printf("%g >= %g\n", v, lim)

	}

	// can't use v here, though
	return lim

}



// ❐❐ Loops and Functions ❐❐
// ※ 要勉強 ※
func LoopFunc(x float64) float64 {

	z := 1.0

	// ループを使った平方根の計算
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z

}


// Switch --if elseステートメントのシーケンスを短く書く方法
func SwitchFuncs() {

	fmt.Println("Go runs on ")

	switch os := runtime.GOOS; os{

		// Goのswitchは条件に合うものだけ選択され、実行される。他のcaseは実行されない。
		// ※Goでは自動でbreakが行われる。
		// また、caseは定数である必要がなく、関係する値は整数である必要もない。
	case  "darwin":
		fmt.Println("OS X.")
	case  "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

}


// Switch evaluation order
func SwitchOrder() {

	fmt.Println("When's Manday?")

	today := time.Now().Weekday()


  // switch caseは上から下へcaseを評価する。
	// caseの条件が一致すれば、そこで停止（自動的にbreak）する。
	// ※break以下のcaseは呼び出されない。
	switch time.Monday {

	case today + 0:
		fmt.Println("today!")
	case today + 2:
		fmt.Println("In twe days")
	default:
		fmt.Println("Too far away")

	}

}



// Switch with no condition
func SwitchTrue() {

	t := time.Now()

	// 条件の無いswitchは、「switch true」と書くことと同じ。
	// 長くなりがちな「if-then-else」のつながりをシンプルにできる。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evning")
	}

}


// Defer
func Defdif() {

	// deferステートメントは、deferに渡した関数の実行を、呼び出し元の関数の終わり（returnする）まで
	// 『遅延させる』もの。
	defer fmt.Println("world")

	// deferに渡した関数の引数は、『すぐに評価される』が、その関数『自体』は呼び出し元の関数がreturnされるまで、
	// 実行されない。

	fmt.Println("hell0")

}


// Stacking defers
func stadef() {

	fmt.Println("counting")


	// deferへ渡した関数が複数あると、その呼び出しはスタック(stack)される。
	// 呼び出し元の関数がreturnするとき、deferへ渡した関数はLIFO（Last In, First Out, 後入れ先出し）の順番で実行される。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	//※ LIFO 古いデータほど、長く残るデータ格納方式。

	fmt.Println("done")

}