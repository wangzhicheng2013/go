package main
import "fmt"
func main() {
    countryCapitalMap := make(map[string]string)
    countryCapitalMap["China"] = "beijing"
    countryCapitalMap["France"] = "Paris"
    for country := range countryCapitalMap {
        fmt.Println(country, "capital = ", countryCapitalMap[country])
    }
    capital, ok := countryCapitalMap["American"]
    if (true == ok) {
        fmt.Println(capital)
    } else {
        fmt.Println("Not Existed...!");
    }
    MM := map[string]int {"A" : 0, "B" : 1, "C" : 2}
    fmt.Print(MM["A"])
}
