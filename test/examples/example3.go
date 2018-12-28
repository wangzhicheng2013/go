package main
import "fmt"
func main() {
    a := 10
    b := 20
    c := a + b
    fmt.Println(c)
    c--
    fmt.Println(c)
    if (a >= b) {
        fmt.Println("a >= b")
    } else {
        fmt.Println("a < b")
    }
    c |= 2
    fmt.Println(c)
    ptr := &a
    fmt.Println(*ptr)
}
