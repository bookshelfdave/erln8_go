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
  "strconv"
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
  var rawconfig interface{}
  err := json.Unmarshal(file, &rawconfig)
  var _ = err
  config := rawconfig.(map[string]interface{})
  erln8Config := config["erln8"].(map[string]interface{})
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
  c := make(chan bool)
  var localfile = filepath.Join(erlangs_dir,"erlangs", filename)
  out, err := os.Create(localfile)
  if err != nil {
    fmt.Println("Can't create file")
    return
  }
  defer out.Close()
  go spinner(c, localfile)
  fmt.Println("Downloading", filename)
  resp, err := http.Get("http://www.erlang.org/download/" + filename)
  defer resp.Body.Close()
  n, err := io.Copy(out, resp.Body)
  close(c)
  fmt.Println("Downloaded", n, "bytes")
}

func spinner(ch chan bool, f string) {
  time.Sleep(1000 * time.Millisecond)
    lastlen := 0
  for {
    select {
      case <- ch:
        return
      default:
          s, _ := os.Stat(f)
          sizeStr := strconv.FormatInt(s.Size(), 10)
          lastlen = len(sizeStr)
          fmt.Printf("%v",s.Size())
          for i := 0; i < int(lastlen); i++ {
            fmt.Printf("\b")
          }
          time.Sleep(250 * time.Millisecond)
   }
  }
  fmt.Printf("Done!")
}

func main() {
  fmt.Printf(Greeting)
  checkConfig()
  var cfg = readConfig()
  var home = cfg["erln8_dir"].(string)
  createErln8DirIfMissing(home)
  downloadErl(home, "otp_src_R15B03.tar.gz")
}


