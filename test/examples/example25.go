package main
import (
    "fmt"
)
type Human interface {
    Run()
    Eat()
    MakeLove()
}
type Man struct
{
    name string
    age int
}
func (man *Man) Run() {
    fmt.Printf("Hi, I am %s age is %d, I am running...!\n", man.name, man.age)
}
func main() {
    man := Man{"Billy", 18}
    man.Run()
    fmt.Println(man)
}

