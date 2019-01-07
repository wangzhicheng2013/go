package main
import "fmt"
func getAverage(arr [10]int) float64 {
    size := len(arr)
    sum := 0
    for i:= 0;i < size;i++ {
        sum += arr[i]
    }
    avg := float64((1.0 * sum / size))
    fmt.Println(avg)
    return 1.0;
}
func main() {
    array := [10] int {1, 2, 3}
    a := [2][2] int {
           {1, 2},
           {3, 4}}
    fmt.Println(array[8])
    fmt.Println(a[1][1])
    getAverage(array)
}
