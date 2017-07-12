// I'm hostname
//create by Tian,Hua  14:18 20170712
// read argument

package main
 
import (
  "fmt"
  "os"
)
 
func main() {
  hostname,err := os.Hostname()
  if err == nil {
    fmt.Println(hostname)
  }
}
