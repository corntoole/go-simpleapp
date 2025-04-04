package main
import (
  "fmt"
  "github.com/corntoole/go-libtest"
  golibtestv2 "github.com/corntoole/go-libtest/v2"
)
func main(){
  fmt.Printf("Using version %s and %s\n", golibtest.Version(), golibtestv2.Version())
}
