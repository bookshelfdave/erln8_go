package main

import (
  "fmt"
  "net/http"
  //"io/ioutil"
  //"regexp"
  "os"
  "io"
  "erln8/config"
)

func download() {
  out, err := os.Create("otp_src_R9C-2.tar.gz")
  if err != nil {
    fmt.Println("Can't create file")
    return
  }
  defer out.Close()
  fmt.Println("Downloading...")
  resp, err := http.Get("http://www.erlang.org/download/otp_src_R9C-2.tar.gz")
  defer resp.Body.Close()
  n, err := io.Copy(out, resp.Body)
  fmt.Println("Downloaded %i bytes", n)
}

type Erln8Config struct {
  LinkDir string
}

type OTPDownloadSource struct {
  Name string
  URL string
}

type OTPVersion struct {
  Name string
  Major string
  Minor string
}


type OTPCompilerFlagsSource struct {
  Name string
  URL string
}

type OTPCompileFlags struct {
  Platform string
  Tag string // release vs debug
  Version string
  Flags string
}

func read_config() {
  s := OTPDownloadSource{"erlang.org", "http://www.erlang.org/download"}
  fmt.Println(s)
}

func main() {
  fmt.Println("erln8")
  e := Erln8Config{"/Users/dparfitt/erl"}
  fmt.Println(e)
  read_config()
/*
  resp, err := http.Get("http://www.erlang.org/download/")
  if err != nil {
    fmt.Println("Error :-(")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  s := string(body)
  r, err := regexp.Compile(`/download/otp_src_R[0-9]+[A-D]+.[0-9]+\.tar\.gz`)
  res := r.FindAllString(s, -1)
  for i := 0; i < len(res); i++ {
    fmt.Println("http://www.erlang.org/download" + res[i])
  }
  download()
  */
}
