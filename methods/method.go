package learn_methods

import (
	"fmt"
	"math"
)

// Methods
// Goにはクラス(class)がない。
// しかし、型にメソッド(Method)を定義できる。
type Oertex struct {
	X, Y float64
}

// メソッドは特別なレシーバ(receiver)引数を関数にとる。
// レシーバは、「func」キーワードとメソッド"名"の間に字志位の引数リストで表現する。
// ※下の例では「(o Oertex)」の"o"がレシーバで、その"o"の型が"Oertex"型。

// 「Abs()」メソッドは、"o"という名前の"Oertex"型のレシーバを持つことを意味する。
func (o Oertex) Abs() float64 {

	return math.Sqrt(o.X*o.X + o.Y*o.Y)

}

func MainMethod() {

	o := Oertex{3, 4}
	fmt.Println(o.Abs(), "Methods")

	fmt.Println(Asd(o), "method function")

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Cds())

	s := Pertex{3, 4}
	s.Scale(10)
	fmt.Println(s.Pbs(), "pointer receivers")
}



// Methods are functions
// メソッドは、『レシーバ引数を伴う関数』。
// Asdは上記の例を"関数"として記述している。
func Asd(o Oertex) float64 {

	return math.Sqrt(o.X*o.X + o.Y*o.Y)

}



// Methods continued
// structの型だけでなく、任意の型(type)にもメソッドを宣言できる。
type MyFloat float64

// Cbsメソッドを持つ、数値型の『MyFloat型』
// レシーバを伴うメソッドの宣言は、レシーバ型が『同じパッケージにある』必要がある。
// （ここでは『learn_methods』パッケージ。
func (f MyFloat) Cds() float64 {

	if f < 0 {
		return float64(-f)
	}

	return float64(f)

}

// 他のパッケージに定義している方に対して、レシーバを伴うメソッドを宣言できない。
// (組み込みのint等の型も同様。)



// Pointer Receviers
type Pertex struct {
	X, Y float64
}

func (p Pertex) Pbs() float64 {

	return math.Sqrt(p.X*p.X + p.Y*p.Y)

}

// ポインタレシーバでメソッドを宣言できる。
// *Pertexに"Scale"メソッドが定義されている。
func (p *Pertex) Scale(f float64) {

	// ポインタレシーバを持つメソッド(ここではScale)は、レシーバを指す変数を変更できる。
	// レシーバ自信を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的。
	p.X = p.X * f
	p.Y = p.Y * f

}

