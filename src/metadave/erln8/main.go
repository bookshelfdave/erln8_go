package main
import (
//  "erln8/config"
  "io/ioutil"
  "fmt"
  "os"
  "os/user"
  "log"
  "json"
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

func readConfig() {
  file, e := ioutil.ReadFile(configFile())
  if e != nil {
    fmt.Printf("Error reading .erln8: %v\n", e)
    os.Exit(1)
  }
  fmt.Printf("%s\n", string(file))
  var jsontype jsonobject
  json.Unmarshal(file, &jsontype)
  fmt.Printf("Results: %v\n", jsontype)
}

func main() {
  fmt.Printf(Greeting)
  checkConfig()
  readConfig()
}
