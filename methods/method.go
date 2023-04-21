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

	ossa := Oertex{3, 4}
	fmt.Println(ossa.Abs(), "Methods")
}