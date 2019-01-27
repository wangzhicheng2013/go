package main
import "fmt"
func main() {
    var numbers = make([]int, 3, 10)
    PrintSlice(numbers)
    var numbers1[] int
    PrintSlice(numbers1)
    numbers1 = append(numbers1, 0)
    numbers1 = append(numbers1, 1)

    PrintSlice(numbers1)
    numbers2 := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    PrintSlice(numbers2)
    fmt.Println(numbers2[:3])
}
func PrintSlice(x [] int) {
    fmt.Printf("len = %d, cap = %d, slice = %v\n",  len(x), cap(x), x)
}
