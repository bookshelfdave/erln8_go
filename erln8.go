package main
import (
//  "erln8/config"
  "io/ioutil"
  "fmt"
  "os"
  "os/user"
  "log"
  "encoding/json"
  "path/filepath"
  "net/http"
  //"io/ioutil"
  //"regexp"
  "io"
  "time"
  )



const DotErln8 = "~/.erln8"
const Greeting = "erln8 v0.1\n"


func configFile() string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal("No home directory detected")
  }
  return usr.HomeDir + string(os.PathSeparator) + ".erln8"
}

// TODO: Windows support
// http://stackoverflow.com/questions/7922270/obtain-users-home-directory
func checkConfig() {
 dotConfig := configFile()
  file, err := os.Open(dotConfig)
  var _ = file
  if err != nil {
    log.Fatal("~/.erln8 configuration file not found")
  }
}

func readConfig() map[string]interface{} {
  file, e := ioutil.ReadFile(configFile())
  if e != nil {
    fmt.Printf("Error reading .erln8: %v\n", e)
    os.Exit(1)
  }
  //fmt.Printf("%s\n", string(file))
  var rawconfig interface{}
  err := json.Unmarshal(file, &rawconfig)
  var _ = err
  //fmt.Printf("Results: %v\n", rawconfig)
  config := rawconfig.(map[string]interface{})
  erln8Config := config["erln8"].(map[string]interface{})
  //fmt.Printf("Erlang directory %v\n", erln8Config["erlang_dir"])
  return erln8Config
}

// http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func createErln8DirIfMissing(dir string) {
  var dirExists, err = exists(dir)
    if err != nil {
      log.Fatal("Error opening %v", dir)
    }
  if !dirExists {
    log.Println(dir + " does not exist, creating")
      os.MkdirAll(filepath.Join(dir,"erlangs"), 0700)
      os.MkdirAll(filepath.Join(dir,"settings"), 0700)
      os.MkdirAll(filepath.Join(dir,"cache"), 0700)
      os.MkdirAll(filepath.Join(dir,"current_erl"), 0700)
  }
}


func downloadErl(erlangs_dir string, filename string) {
  var localfile = filepath.Join(erlangs_dir,"erlangs", filename)
  out, err := os.Create(localfile)
  if err != nil {
    fmt.Println("Can't create file")
    return
  }
  defer out.Close()
  fmt.Println("Downloading", filename)
  resp, err := http.Get("http://www.erlang.org/download/" + filename)
  //resp, err := http.Get("http://www.erlang.org/download/otp_src_R15B03.tar.gz")
  defer resp.Body.Close()
  n, err := io.Copy(out, resp.Body)
  fmt.Println("Downloaded %i bytes", n)
}

func spinner() {
  time.Sleep(1000 * time.Millisecond)
  fmt.Printf("*")
  chars := []string{"\b|","\b/","\b-","\b\\"}
  for i := 0; i < 100; i++ {
    for c := range chars {
      fmt.Printf(chars[c])
      time.Sleep(250 * time.Millisecond)
    }
  }
}

func main() {
  fmt.Printf(Greeting)
  checkConfig()
  var cfg = readConfig()
  var home = cfg["erln8_dir"].(string)
  createErln8DirIfMissing(home)
  go spinner()
  downloadErl(home, "otp_src_R15B03.tar.gz")
  spinner()
}


