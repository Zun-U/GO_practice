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
// 呼び出すのには、関数の名前を再度指定する必要がある
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

	// ☆☆☆　型には変数の値を保護する役割がある　☆☆☆
	// 「int」　正・負の整数
	// 「uint」 正の整数"のみ"
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




	// Data Types
	var userName string
	// ask user for the name
	var userTicket int

	userName = "Tom"
	userTicket = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTicket)

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









