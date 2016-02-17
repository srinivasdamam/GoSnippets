package main
import "fmt"

func main() {
  name:="Golang"
  ptr:=&name
  fmt.Println("Address of name is", ptr, "and value is ", *ptr)
}
