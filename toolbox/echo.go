// I'm echo
//create by Tian,Hua  14:18 20170705
// read argument

package main
import "os"
import "fmt"

func main() {
    sdf := os.Args[1:]
    for i:=0; i<len(sdf); i++{
        fmt.Print(sdf[i])
    }
    fmt.Println()
}
