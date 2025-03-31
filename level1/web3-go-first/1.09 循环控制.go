package main
import "fmt"

func main() {

	 // 方式1
	//  for i := 0; i < 10; i++ {
    //     fmt.Println("方式1，第", i + 1,"次循环")
    // }
    
    // 方式2
    // b := 1
    // for b < 10 {
    //     fmt.Println("方式2，第", b,"次循环")
    // }

	 // 遍历数组
	//  var a [10]string
	//  a[0] = "Hello"
	//  for i := range a {
	// 	 fmt.Println("当前下标：", i)
	//  }
	//  for i, e := range a {
	// 	 fmt.Println("a[", i, "] = ", e)
	//  }

	// 遍历切片
    // s := make([]string, 10)
    // s[0] = "Hello"
    // for i := range s {
    //     fmt.Println("当前下标：", i)
    // }
    // for i, e := range s {
    //     fmt.Println("s[", i, "] = ", e)
    // }

	// m := make(map[string]string)
    // m["b"] = "Hello, b"
    // m["a"] = "Hello, a"
    // m["c"] = "Hello, c"
    // for i := range m {
    //     fmt.Println("当前key：", i)
    // }
    // for k, v := range m {
    //     fmt.Println("m[", k, "] = ", v)
    // }

	//  // 中断for循环
	//  for i := 0; i < 5; i++ {
    //     if i == 3 {
    //         break
    //     }
    //     fmt.Println("第", i, "次循环")
    // }

	 // 中断switch
	//  switch i := 1; i {
    // case 1:
    //     fmt.Println("进入case 1")
    //     if i == 1 {
    //         break
    //     }
    //     fmt.Println("i等于1")
    // case 2:
    //     fmt.Println("i等于2")
    // default:
    //     fmt.Println("default case")
    // }

	// outter:
    // for i := 1; i <= 3; i++ {
    //     fmt.Printf("使用标记,外部循环, i = %d\n", i)
    //     for j := 5; j <= 10; j++ {
    //         fmt.Printf("使用标记,内部循环 j = %d\n", j)
    //         break outter
    //     }
    // }

	 // 中断for循环
	//  for i := 0; i < 5; i++ {
    //     if i == 3 {
    //         continue
    //     }
    //     fmt.Println("第", i, "次循环")
    // }

	// outter:
    // for i := 1; i <= 3; i++ {
    //     fmt.Printf("使用标记,外部循环, i = %d\n", i)
    //     for j := 5; j <= 10; j++ {
    //         fmt.Printf("使用标记,内部循环 j = %d\n", j)
    //         if j >= 7 {
    //             continue outter
    //         }
    //         fmt.Println("不使用标记，内部循环，在continue之后执行")
    //     }
    // }

}