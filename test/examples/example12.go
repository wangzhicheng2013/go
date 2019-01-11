package main
import "fmt"
func main() {
    a := 10
//    var p *int = &a 
    p := &a 
    fmt.Println(p)
    fmt.Println(&a)
    if (p != nil) {
        fmt.Println("1")
    }
    array := [] int {1, 2, 10}
    q := &array
    i := 0
    for ;i < 3;i++ {
        fmt.Println(*q);
    }
    var ff [3] * int

}
