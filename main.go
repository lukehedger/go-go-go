// Executable commands must always use `package main`
package main

import (
	// "fmt"
	"github.com/lukehedger/golib"
)

func main() {
	// fmt.Printf("Go, Go, Go!\n")
	golib.Echo(golib.Reverse("\n!oG, oG, oG"))
}
