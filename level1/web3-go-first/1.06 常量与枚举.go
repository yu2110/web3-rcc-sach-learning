package main
import "fmt"

// // 方式1
// const a int = 1

// // 方式2
// const b = "test"

// // 方式3
// const c, d = 2, "hello"

// // 方式4
// const e, f bool = true, false

// // 方式5
// const (
//     h    byte = 3
//     i         = "value"
//     j, k      = "v", 4
//     l, m      = 5, false
// )

// const (
//     n = 6
// )

type Gender string

const (
   Male   Gender = "Male"
   Female Gender = "Female"
)

func (g *Gender) String() string {
    switch *g {
    case Male:
        return "Male"
    case Female:
        return "Female"
    default:
        return "Unknown"
    }
}

func (g *Gender) IsMale() bool {
    return *g == Male
}

const pre int = 1
const a int = iota
const (
    b int = iota
    c
    d
    e
)
const (
    f = 2
    g = iota
    h
    i
)
func main() {
     fmt.Println(f)
}