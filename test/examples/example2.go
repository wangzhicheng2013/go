package main
import "fmt"
const (
    Unknown = 0
    Female = 10
    Male = 2
)
const (
    AA = iota
    BB
    CC = "CC"
    DD
    EE = iota
)
func main() {
    const PI float64 = 3.1213
    const STR = "hello world"
    fmt.Println(PI, STR, Female, AA, BB, CC, DD, EE)
}
