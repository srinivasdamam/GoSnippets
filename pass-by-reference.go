package main
import (
  "fmt"
  "reflect"
)

func main() {
  course:="golang"
  category:="programing"

  fmt.Println("Type of course is ", reflect.TypeOf(course))
  fmt.Println("Course name", course, "category is", category)

  //passing by reference
  changeCourse(&course)

  fmt.Println(course)
}

func changeCourse(course *string) string {
  *course = "Node"
  fmt.Println("Course is", *course)
  return *course
}
