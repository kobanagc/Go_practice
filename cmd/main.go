// Goのファイルはなんらかのpackageに属している必要があり、そのうち1つはmainに属している必要がある
package main
// fmtはよく使う関数をまとめたもの
import (
	"fmt"
	"reflect"
)


// Goのプログラムはmainパッケージのmain関数が最初に実行される（エントリーポイント）
func main() {
	fmt.Println("Hello Go!")

	var num int = 10 // intは省略可能
	fmt.Println(num)

	num2 := 20 // :=の代入方法はよく使われる（varは不要）
	fmt.Println(num2)

	Num := 30 // 変数の1文字目が小文字の場合はそのパッケージのみ、大文字の場合は全てのパッケージで使用可能
	fmt.Println(Num) // 関数名も同様。Printlnはmainパッケージの関数ではなく、fmtパッケージの関数
	                 // しかし大文字なのでmainパッケージでも使用可能
	num01 := 123
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))

	var string_a string = "Hello"
	string_b := "World"

	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

	a := 10
	b := 20
	num_true := a < b
	num_false := a > b
	fmt.Println(num_true)
	fmt.Println(reflect.TypeOf(num_true))
	fmt.Println(num_false)
	fmt.Println(reflect.TypeOf(num_false))
}