package main
import "fmt"
func main() {
    array := [] int {1, 2, 3, 4}
    for i, num := range array {
        fmt.Println("seq = ", i, "num = ", num)
    }
    sum := 0
    for _, num := range array {
        sum += num
    }
    fmt.Println("sum = ", sum)
    kvs := map[string]string {"a":"0", "b":"1"}
    for k, v := range(kvs) {
        fmt.Println("key = ", k, " value = ", v)
    }
}
