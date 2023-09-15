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

	arr := [3]string {"a", "b", "c"} // []の中は要素数
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	arr[2] = "d"
	fmt.Println(arr[2])
	fmt.Println(arr)

	arr2 := [...]string{"e", "f", "g"} // 要素数は「...」で省略可能だが、「...」も省略できる
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr3 := [][]string{{"h", "i"}, {"j", "k"}} //　多次元配列の場合は「...」は使えない
	fmt.Println(arr3[0][0])
	fmt.Println(arr3[0][1])
	fmt.Println(arr3[1][0])
	fmt.Println(arr3[1][1])


	c := 10
	d := 2
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)
	fmt.Println(c % d)

	fmt.Println(c > d)
	fmt.Println(c < d)
	fmt.Println(c == d)
	fmt.Println(c != d)
	fmt.Println(c > 8 && c < 20)
	fmt.Println(c > 8 || c < 20)

	c += 10
	d += b
	fmt.Println(c)
	fmt.Println(d)

	c ++
	d --
	fmt.Println(c)
	fmt.Println(d)

	if c > d {
		fmt.Println("c > d")
	} else if c < d {
		fmt.Println("c < d")
	} else {
		fmt.Println("c == d")
	}

	if age := 10; age < c { // ifの後に変数を宣言することもできる。しかしスコープはif文の中。
		fmt.Println("age < c")
	} else {
		fmt.Println("age > c")
	}
}