// created on 13:21 20170712 
// https://golang.org/src/net/ip.go
sdfsdfsd
package main
import (
    "net"
    "fmt"
)
func main() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }
    for _, addr := range addrs {
        fmt.Println(addr.String())
    }
}

