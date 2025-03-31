package main
import "fmt"

type Other struct {}

type Person struct {
    Name  string             `json:"name" gorm:"column:<name>"`
    Age   int                `json:"age" gorm:"column:<name>"`
    Call  func()             `json:"-" gorm:"column:<name>"`
    Map   map[string]string  `json:"map" gorm:"column:<name>"`
    Ch    chan string        `json:"-" gorm:"column:<name>"`
    Arr   [32]uint8          `json:"arr" gorm:"column:<name>"`
    Slice []interface{}      `json:"slice" gorm:"column:<name>"`
    Ptr   *int               `json:"-"`
    O     Other              `json:"-"`
}

type A struct {
    a     string
    bytes [2]byte
}

func (a A) string() string {
    return a.a
}

func (a A) stringA() string {
    return a.a
}

func (a A) setA(v string) {
    a.a = v
}

func (a *A) stringPA() string {
    return a.a
}

func (a *A) setPA(v string) {
    a.a = v
}

func value(a A, value string) {
    a.a = value
}

func point(a *A, value string) {
    a.a = value
}


func main() {
	a := A{
        a: "a",
    }

    value(a, "any")

    point(&a, "any")

    pa := &a

    // a *A
    a.setPA("pa")

    // a A
    fmt.Println(a.string())
    // a A
    fmt.Println(a.stringA())
    // a *A
    fmt.Println(a.stringPA())

    // a A
    fmt.Println(pa.string())
    // a A
    fmt.Println(pa.stringA())
    // a *A
    fmt.Println(pa.stringPA())
}