package main
import (
    "fmt"
)
func foo() * int {
    a := 10
    return &a
}
func main() {
    x := foo()
    fmt.Println(*x)
}

