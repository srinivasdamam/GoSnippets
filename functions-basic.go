package main
import (
  "fmt"
  "strings"
)

func main() {
  course:="go lang"
  category:="programing language"
  fmt.Println(convert(course, category))
}

func convert(course string, category string) (s1, s2 string) {
  course = strings.ToUpper(course)
  category = strings.Title(category)
  return course, category
}
