package main
import "fmt"

func main() {

// 	var p1 *int
// 	var p2 *string

// 	i := 1
// 	s := "Hello"
// 	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
//    // 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
// 	p1 = &i
// 	p2 = &s

// 	//var p3 **sring = &p2 指针的指针  指针变量的内存地址
// 	p3 := &p2
// 	fmt.Println(p1)
// 	fmt.Println(p2)
// 	fmt.Println(p3)


	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}