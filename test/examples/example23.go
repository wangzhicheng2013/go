package main
import (
    "fmt"
)
type Person struct {
    id string
    name string
    age int
}
type Stu struct {
    Person
    sex string
}
type Interval struct {
    start int
    end int
}
func (p Person) Show() {
    fmt.Println(p)
}
func (p *Person) Update() {
    p.age = 22
    p.id = "12345"
    p.name = "6778"
}
func main() {
    var person Person
    person.id = "id00001"
    person.name = "zhangsan"
    person.age = 33
    fmt.Println(person)
    interval0 := Interval{0, 10}
    interval1 := Interval{start:100, end:110}
    interval2 := Interval{end:1110}
    fmt.Println(interval0)
    fmt.Println(interval1)
    fmt.Println(interval2)
    person.Show()
    stu := Stu{sex:"female"}
    stu.id = "1111"
    stu.name = "2222"
    stu.age = 11
    fmt.Println(stu)
    person.Update()
    fmt.Println(person)
}
