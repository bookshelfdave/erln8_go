package main
import (
  "erln8/config"
  "fmt"
)

func main() {
  x := config.DoStuff()
  fmt.Printf("Hello world %i", x)
}
