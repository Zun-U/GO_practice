// GOが普及した背景
// マルチコアプロセッサの普及
// クラウドインフラストラクチャの使用による何百、何千ものサーバー
// これらにより、複数のプロセッサを導入するアプリケーションは普遍的になった為、インフラストラクチャの拡張性が大幅に向上
// ダイナミックでより多くの容量がある

// つまり、ハードの機能が良くなったので、複数のタスクを『同時並行』でこなせればより効率が上がる

// ❐❐　マルチスレッドの概念　❐❐
// スレッドは基本的に一つのタスクを処理する

// ☆☆☆☆☆☆☆☆☆☆☆☆
// GOはマルチスレッドの同時書き込みを可能に
// それが『go routin』

// また、GOは、保守しやすく、読みやすく、拡張しやすい
// GOのアプリケーションは非常に高速に構築できる
// 起動も非常に速い
// 実行するCPUやRAM等の使用するリソースが少ない（効率的である）
// ☆☆☆☆☆☆☆☆☆☆☆☆




// GOの導入
// 1.goコンパイラが必要
// 2.IDE等の統合環境


// GOのインストール
// https://go.dev/doc/install


// ::::::::::::::::::::::::::
// GOのアプリケーションにする
// Initialize our project or "module"!
//
// [command]
// go mod ini <module path>
//
// 例) go mod init booking-app
// ::::::::::::::::::::::::::


// ❐❐❐　GOで記述されたすべてのコードは、パッケージに属さなければいけない　❐❐❐
// ※パッケージで整理される
//
// パッケージは2種類ある
// 1.アプリケーション全体にわたるパッケージ
// 2.独自のアプリケーションを作成するときのパッケージ



// GOアプリケーションの『最初の行』にパッケージを定義する
// アプリケーションはそのパッケージの一部となる
// GOのパッケージは必ずmainから始まる（初めにmainを通る）
package main






// 「Print」関数は、GOの組み込みパッケージ「fmt」、または「format」に属している。
// ❐❐❐  使用するのには、元のパッケージをインポートする必要がある  ❐❐❐
import (
	"fmt"
)


// GOはどこで実行するか
// コンパイラの開始点の最初の行

// アプリケーションのエントリーポイント
// Goアプリケーションを実行するとき、出発点(エントリーポイント)が必要



// 『main』functionはGO言語におけるentry point
// main関数（エントリーポイント）以外の関数は、呼び出すのに関数の名前を再度指定する必要がある
func main() {


	// ::::  Variable  ::::
	//
	// 「var」キーワードで変数を宣言することができる
	// 以下では、変数「name」から値を参照できるようになった
	// 変数名の命名は具体的で、明確にする（ここではキャメルケース）
	//
	// ☆☆☆　慣例として、変数は関数の始まりに定義する　☆☆☆
	//
	// 変数宣言の代替構文
	conferenceName := "Go Conference"


	// ::::  Constnats  ::::
	//
	// 「const」キーワードで定数を宣言することができる
	// 一度定義したら、値は変わらない(変更できない)
	// ※「定数は」値の変更が"許可"されていないと言える
	// ※ 定数は暗黙的な宣言が出来ない
	const conferenceTickets int = 50


	// ❐❐❐　GOは値に基づくデータ型を判断できる　❐❐❐
	// ☆ 暗示的変数宣言、定数宣言
	//
	// *** 型を明示する場合は、「var」による変数宣言を行う ***
	var remaningTickets uint = 50


	//
	// ☆☆☆☆☆　型には変数の値を保護する役割がある　☆☆☆☆☆
	// 「int」　正・負の整数
	// 「uint」 符号なし整数"のみ"
	// なので、数値型だけでもたくさん種類がある(int8 int32、float32、float64...)
	//
	// float型は統計データや比較、金銭的な値などによく使用される　>>> 『浮動小数点数は高精密』



	// 「%T」 - 任意のタイプを出力する
	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remaningTickets, conferenceName)

	// Print関数を使用するのには、「Print」どこから来たかGOに教える必要がある
	// ここでは「fmtパッケージのPrint関数」
	// ※カンマ区切りで変数を挿入することができる。その際、変数の前後に「スペースが」挿入される
	// ☆☆☆　そしてPrintlnは自動で改行する（新しい行に出力する）　☆☆☆
	// fmt.Println("Welcome to", conferenceName, "booking application")


	// ❐　printf　❐
	// フォーマットされたデータを出力する
	// 「%v」 - プレースホルダー。のちに渡す引数の値が格納される。「変数参照」
	// ※変数のフォーマットを行っている。%vは対応する引数から参照する
	// 「\n」 - 改行を表す
	fmt.Printf("Welcome to %v booking application\n", conferenceName)


	// fmt.Println("We have total of", conferenceTickets, "tickets and", remaningTickets, "are still available.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")


	// GO特有の文化
	// 変数を定義するとき、または特定の値を持つ変数を作成した際に、その『変数が使用されていない』というエラーが発生する。
	// ❐❐❐ パッケージも同様に、使用されていないパッケージが宣言されているとエラーが発生する　❐❐❐


	// [[[[[[[[[  for loop  ]]]]]]]]]
	for i := 0; i < count; i++ {
		
	}



	// Data Types
	var firstName string
	var lastName string
	var email string
	// ask user for the name
	var userTicket uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask user for their name
	// fmtパッケージは、入出力機能が搭載されている。「Scan」はその入力機能の内の一つ


	// >>>>>>>>>>>>>    pointer     >>>>>>>>>>>>>>>>>
	// 変数に値を代入するときに、実際にその値は『メモリ』に保存される
	//
	// その値を参照する『たび』に、GOのコンパイラがメモリ内でその値を見つけに行く必要がある為、
	// 『メモリ内の正確な場所を知る必要がある』　=>  "メモリアドレス"
	//
	// pointerはその値のメモリアドレスを指す『変数』
	// pointerはgolangでは特殊変数と呼ばれる
	//
	// &userName -> 変数「userName」のpointer(userNameにメモリアドレスであるハッシュが格納されている変数)
	//
	//
	// 「var userName string」の変数宣言では、"値が入っていない"
	//
	// 「fmt.Scan(&userName)」には、メモリアドレスを渡す。
	// ユーザーが入力したものはすべて読み取られ、その値をメモリ内の変数に"割り当てる"。
	//
	// つまり、Scan関数は、渡された引数のメモリアドレスに値を割り当てる関数である
	//
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<


	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)


	fmt.Println("Enter your email: ")
	fmt.Scan(&email)


	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	remaningTickets = remaningTickets - userTicket
	// 「uint」型と「int」型と型が異なる為、計算することができない

	// ++++++++++++++ 配列 ++++++++++++++++
	// ❐ 配列の初期化に必要なのは『サイズ』❐
	// 角括弧([])でサイズを指定する
	// 指定したサイズ分の『要素』を格納できる
	//
	// ❐ そしてもう一つ定義する必要があるのは、『型』
	//
	// ❐ 配列は初期化して使う ❐
	var bookings [50]string

	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value:%v\n", bookings[0])
	fmt.Printf("Array type:%T\n", bookings[0])     // 「%T」　型を出力する
	fmt.Printf("Array length:%v\n", len(bookings)) // 「len」関数　配列の長さを出力する　※Goの組み込み関数



	// ************* スライス *****************
	// ❐ スライスの定義は、『長さ』を指定しない配列 ❐
	bookingsSlice := []string{}
	// ※以下の様な宣言方法もある
	// var bookingsSlice = []string{}
	// var bookingsSlice []string

	// 「append」関数　スライスに要素を追加する　※Goの組み込み関数
	bookingsSlice = append(bookingsSlice, firstName + " " + lastName) //　☆☆　『長さ』は自動的に拡張される(Auto Increment)　☆☆


	fmt.Printf("The whole slice: %v\n", bookingsSlice)
	fmt.Printf("The first value:%v\n", bookingsSlice[0])
	fmt.Printf("Slice type:%T\n", bookingsSlice[0])     // 「%T」　型を出力する
	fmt.Printf("Slice length:%v\n", len(bookingsSlice)) // 「len」関数　配列の長さを出力する　※Goの組み込み関数


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remaningTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookingsSlice)

}


// アプリケーションが属するポイントパッケージにはエントリーポイントがある為、GOはどこから実行を開始するかを知っている



// *************************************
// GOのプログラムは、パッケージでまとめられている
//
// また、GOは基本的なライブラリ、組み込みパッケージが含まれている
// 例)
// ❐fmtパッケージ
//  -- 〇〇.go
//  -- △△.go
//  -- ◇◇.go
//
// GOが提供している基本ライブラリ
// https://pkg.go.dev/std
// *************************************



// ☆☆☆　GOの実行　☆☆☆
// go run ファイル名
//
// 例) go run tutorial.go
// ☆☆☆☆☆☆☆☆☆☆☆☆






// 配列とスライスと呼ばれるデータ型


