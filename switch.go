package main

import "fmt"

func main() {
  i:="Monday"
  switch i {
  case "Monday": fmt.Println("Today is Monday")
  default: fmt.Println("Today is not Monday") //can be optional
  }
}
