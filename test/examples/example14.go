package main
import "fmt"
type Book struct {
    title string
    price int
    author string
}
func main() {
    a := 10
    b := 20
    swap(&a, &b)
    fmt.Printf("a = %d, b = %d\n", a, b)
    book := Book{"123", 12, "456"}
    ShowBook(&book)
}
func swap(a *int, b *int) {
    *a, *b = *b, *a
}
func ShowBook(book *Book) {
    if (nil == book) {
        return
    }
    fmt.Printf("title = %s, price = %d, author = %s\n", book.title, book.price, book.author)
}
