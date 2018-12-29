package main
import "fmt"
func main() {
    grade := 100
    score := "A"
    switch (grade) {
        case 100:
            score = "A"
        case 90:
            score = "B"
        case 80, 75:
            score = "C"
        case 70, 60:
            score = "D"
        default:
            score = "E"
    }
    fmt.Println(score)
}
