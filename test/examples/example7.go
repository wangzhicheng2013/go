package main
import "fmt"
func main() {
    a := 1
    var b chan int
    fmt.Println(a)
    select {
        case a = <-b:
            fmt.Println("a receive from b")
        default:
            fmt.Println("no select")
    }
}
