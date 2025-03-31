package main
import "fmt"


var s1 string = "Hello"
var zero int
var b1 = true

var (
    i int = 123
    b2 bool
    s2 = "test"
)

var (
    group = 2
)

func main() {
   
	// fmt.Println(s1)
	// fmt.Println(zero)
	// fmt.Println(b1)
	// fmt.Println(i)
	// fmt.Println(b2)
	// fmt.Println(s2)
	// fmt.Println(group)

	method1()
	a2, b2:= method2()
	a3, b3:= method3()
	a4, b4:= method4()

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(a3)
	fmt.Println(b3)
	fmt.Println(a4)
	fmt.Println(b4)
}

func method1() {
    // 方式1，类型推导，用得最多
    a := 1
    // 方式2，完整的变量声明写法
    var b int = 2
    // 方式3，仅声明变量，但是不赋值，
    var c int
    fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func method2() (a int, b string) {
    // 这种方式必须声明return关键字
    // 并且同样不需要使用，并且也不用必须给这种变量赋值
    return 1, "test"
}

func method3() (a int, b string) {
    a = 1
    b = "test"
    return
}

func method4() (a int, b string) {
    return
}