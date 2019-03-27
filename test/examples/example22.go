package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("please input your name:")
    input, err := inputReader.ReadString('\n')
    if err != nil {
        fmt.Printf("error = %s\n", err)
        return
    }
    input = input[:len(input) - 1]
    fmt.Printf("Hello %s\n", input)
}
