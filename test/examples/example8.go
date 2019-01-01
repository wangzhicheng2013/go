package main
import "fmt"
func main() {
numbers := [6] int {1, 2, 3, 4, 5}
    for i := 0;i < 6;i++ {
        fmt.Println(numbers[i])
    }
    for i, x := range numbers {
        fmt.Println(i, x)
    }
}
