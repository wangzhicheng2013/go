package main
import (
    "fmt"
)
func say(s string) {
    for i := 0;i < 100;i++ {
        fmt.Println(s)
    }
}
func main() {
    go say("go go go...!")
    say("hey")
}
