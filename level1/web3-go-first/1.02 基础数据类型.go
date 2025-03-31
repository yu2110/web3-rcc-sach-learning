package main
import "fmt"

func main() {
	var a uint8 = 0xF
	var b uint8 = 0xf
	
	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0O17
	
	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0B1111
	
	// 十进制
	// var h uint8 = 15

	fmt.Println(a==b);

	fmt.Println(c==d);
	fmt.Println(d==e);
	fmt.Println(f==g);

	var float1 float32 = 10
    float2 := 10.0
	fmt.Println(float1 == float32(float2))

	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1) // c2与c3是等价的
	fmt.Println(c1 == complex64(c2))
	fmt.Println(c2== c3)
	x := real(c2)
    y := imag(c2)
	fmt.Println(x)
	fmt.Println(y)

	// var s string = "Hello, world!"
    // var bytes []byte = []byte(s)
    // fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
    // fmt.Println(string(bytes) == s)


	var r1 rune = 'a'
    var r2 rune = '世'
	fmt.Println(r1)
	fmt.Println(r2)

	// var s string = "abc，你好，世界！"
    // var runes []rune = []rune(s)
	// fmt.Println(runes)
	// fmt.Println(len(runes))

// 	var s string = "Hello\nworld!\n"
    
// 	var s1 string = `Hello
// world!
// `
//     fmt.Println(s==s1)

	// var s string = "Go语言"
	// var bytes []byte = []byte(s)
	// var runes []rune = []rune(s)

	// fmt.Println("string length: ", len(s))
	// fmt.Println("bytes length: ", len(bytes))
	// fmt.Println("runes length: ", len(runes))


	// var s string = "Go语言"
    // var bytes []byte = []byte(s)
    // var runes []rune = []rune(s)

    // fmt.Println("string sub: ", s[0:7])
    // fmt.Println("bytes sub: ", string(bytes[0:7]))
    // fmt.Println("runes sub: ", string(runes[0:3]))

	var m int
	var m1 string
	var m2 bool

	fmt.Println(m)

	fmt.Println(m1)

	fmt.Println(m2)
}