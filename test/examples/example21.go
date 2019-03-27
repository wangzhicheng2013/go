package main
import (
    "fmt"
    "runtime"
)
var info string
var mapper = map[int]string {1 : "A", 2 : "B"}
func init() {
    fmt.Printf("Map:%v\n", mapper)
    info = fmt.Sprintf("OS:%s, Arch:%s\n", runtime.GOOS, runtime.GOARCH)
}
func main() {
    fmt.Printf(info)
}
