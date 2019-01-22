package main
import "fmt"
func main() {
    var a int = 10
    var ptr1 * int = &a
    var ptr2 ** int = &ptr1
    fmt.Printf("a = %d\n", **ptr2);
}
