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

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, "-", j)
		}
	}

	num_arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range num_arr { // _はブランク演算子。通常iなどのindexが入るがforの中でiを使わない場合に_を使う
		if num % 2 == 0 {
			fmt.Println(num)
		}
	}

	hello()
	greeting("John")
	fmt.Println(calc(10, 20))
	fmt.Println(calc2(10, 20))
	fmt.Println(calc3(10, 20))
	fmt.Println(calc4(10, 20))

	// 無名関数の場合は関数の中に書ける（呼び出しなしで即実行）
	func () {
		fmt.Println("no name function!!")
	}()

	// 関数を変数に代入することも可能
	f := func () {
		fmt.Println("関数を変数に代入!!")
	}
	f()

	var student Student
	student.name = "John"
	student.math = 80
	student.english = 70
	fmt.Println(student)

	student2 := Student{"Ken", 90, 80} // フィールド(nameなど)が順番通りだったら引数にフィールド名を書かなくてよい
	fmt.Println(student2)

	student3 := Student{
		english: 90, // フィールドを引数に書いていれば、フィールド名の順番は不問
		name: "Mike",
		math: 100,
	}

	student3.avg()
	student3.avg2(100, 100)
	fmt.Println(student3.avg3(60, 50))
	fmt.Println(student3.avg4(60, 50))
}

// 関数は必ずトップレベルに書く（関数内関数はできない）
// 関数の定義は呼び出しよりも下に書いてOK
func hello() {
	fmt.Println("Hello World!")
}

func greeting(name string) {
	fmt.Println("Hello " + name + "!")
}

func calc(num1 int, num2 int) int {
	return num1 + num2
}

func calc2 (num1, num2 int) (int, int) {
	return num1 * num2, num1 + num2
}

func calc3 (num1, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum , sub
}

func calc4 (num1, num2 int) (sum, sub int) { // 戻り値でsum, subを宣言しているので:=（宣言+代入）ではなく=（代入）で書く
	sum = num1 + num2
	sub = num1 - num2
	return sum, sub // 戻り値でsum, subを宣言しているためreturnの後は省略可能
}

// 構造体
type Student struct {
	name string
	math, english float64
}

// メソッド（= 構造体に紐づけられた関数）
func (s Student) avg() {
	fmt.Println(s.name, (s.math + s.english) / 2)
}

func (s Student) avg2(math, english float64) { // フィールドを引数に書いていれば、メソッドのなかでs.を書かなくて良い
	fmt.Println(s.name, (math + english) / 2)
}

func (s Student) avg3(math, english float64) float64 {
	return (math + english) / 2
}

func (s Student) avg4(math, english float64) (average_score float64) {
	average_score = (math + english) / 2
	return
}