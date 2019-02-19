package main
import "fmt"
type Phone interface {
    Call()
}
type MobilePhone struct {
}
func (phone MobilePhone) Call () {
    fmt.Println("I am mobile phone...!")
}
type Car interface {
    MakeCar() string
}
type FengTianCar struct {
}
func (car FengTianCar) MakeCar() string {
    return "I am a fengtian car"
}
func main() {
    phone := new(MobilePhone)
    phone.Call()
    car := new(FengTianCar)
    fmt.Println(car.MakeCar())
}

