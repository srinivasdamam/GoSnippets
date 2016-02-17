package main
import (
  "fmt"
  "os"
)

func main() {
  for _, val := range os.Environ() {
    fmt.Println(val)
  }
}
