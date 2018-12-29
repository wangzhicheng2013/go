package main
import "fmt"
func main() {
    total := 0
    isPrime := true
    for number := 2;number < 10000;number++ {
        isPrime = true
        for i := 2;i <= number / 2;i++ {
            if (0 == number % i) {
                isPrime = false
                break
            }
        }
        if (true == isPrime) {
            total++
        }
    }
    fmt.Println("toal prime number = ", total);
}
