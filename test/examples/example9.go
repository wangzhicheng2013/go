package main
import (
        "fmt"
        "math")
func Max(a, b int) int {
    max := a
    if (max < b) {
        max = b
    }
    return max
}
func SwapString(a, b string) (string, string) {
    return b, a
}
func SwapInt(a, b *int) {
    tmp := *a
    *a = *b
    *b = tmp
}
type Circle struct {
    radius float64
}
func (c Circle)GetArea() float64 {
    return 3.14 * c.radius * c.radius
}
func main() {
    fmt.Println(Max(10, 100))
    a := "world"
    b := "Hello"
    fmt.Println(SwapString(a, b))
    x := 10
    y := 20
    SwapInt(&x, &y)
    fmt.Println(x, y)
    GetSqrt := func(x float64) float64 {
        return math.Sqrt(x)
    }
    fmt.Println(GetSqrt(100))
    var c Circle
    c.radius = 10
    fmt.Println(c.GetArea())

}
