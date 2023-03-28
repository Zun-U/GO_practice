package memory_test

import (

	"fmt"

)

func Main() {
	fmt.Println("This package is pointer")
}


// Goはポインタを扱う。
// ポインタの値はメモリアドレスを指す。
// 変数「T」のポインタは、*T型で、ゼロ値はnil。
func PointerTest(){

	i, j := 42, 2701

	// &オペレータは、そのオペランド(operand)へのポインタを引き出す。
	p := &i

	fmt.Println(*p)

	*p = 21

	fmt.Println(i)



	p = &j
	*p = *p /37
	fmt.Println(j)


}