package main
import "fmt"

func main() {
	// var a int = 10
    // if b := 1; a > 10 {
    //     b = 2
    //     // c = 2
    //     fmt.Println("a > 10")
    // } else if c := 3; b > 1 {
    //     b = 3
    //     fmt.Println("b > 1")
    // } else {
    //     fmt.Println("其他")
    //     if c == 3 {
    //         fmt.Println("c == 3")
    //     }
    //     fmt.Println(b)
    //     fmt.Println(c)
    // }

	// a := "test string"

    // // 1. 基本用法
    // switch a {
    // case "test":
    //     fmt.Println("a = ", a)
    // case "s":
    //     fmt.Println("a = ", a)
    // case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
    //     fmt.Println("catch in a test, a = ", a)
    // case "n":
    //     fmt.Println("a = not")
    // default:
    //     fmt.Println("default case")
    // }


	var d interface{}
    // var e byte = 1
    d = 1
    switch t := d.(type) {
    case byte:
        fmt.Println("d is byte type, ", t)
    case *byte:
        fmt.Println("d is byte point type, ", t)
    case *int:
        fmt.Println("d is int type, ", t)
    case *string:
        fmt.Println("d is string type, ", t)
    // case *CustomType:
    //     fmt.Println("d is CustomType pointer type, ", t)
    // case CustomType:
    //     fmt.Println("d is CustomType type, ", t)
    default:
        fmt.Println("d is unknown type, ", t)
    }
}