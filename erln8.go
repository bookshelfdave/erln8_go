package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  fmt.Println("erln8")

  resp, err := http.Get("http://www.erlang.org/download/")
  if err != nil {
    fmt.Println("Error :-(")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(body)
}
