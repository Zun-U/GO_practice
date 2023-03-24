package sub

	import (
		"fmt"
		"math"
	)

func HelloGo() {
	fmt.Println("this Folder is sub")
	fmt.Println(math.Pi)
	fmt.Println(add(5, 6))
	fmt.Println(addPlus(11, 12))
}


// 変数名の『後ろ』に型を指定する。
// 理由は右から左に読めるように
func add(xa int, ya int) int {
	return xa + ya
}

// 複数の引数において、型がどれも同じだった場合、最後の型を残して他を省略できる。
func addPlus (xz, yz int) int {
	return xz + yz
}