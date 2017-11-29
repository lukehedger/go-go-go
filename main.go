// Executable commands must always use `package main`
package main

// 'factored' import statements can be used to group individual imports
import (
	"fmt"
	"github.com/lukehedger/golib"
)

func main() {
	fmt.Printf(":: Reverse ::\n")
	golib.Echo(golib.Reverse("\n!oG ,oG ,oG"))
	fmt.Printf("\n")

	fmt.Printf(":: Concat ::\n")
	a, b := "Hello", "World"
	golib.Echo(golib.Concat(a, b))
	fmt.Printf("\n")

	fmt.Printf(":: Conditioner ::\n")
	golib.Conditioner(9)
	fmt.Printf("\n")

	fmt.Printf(":: Looper ::\n")
	golib.Looper()
	fmt.Printf("\n")

	fmt.Printf(":: Pointers ::\n")
	golib.Pointers()
	fmt.Printf("\n")

	fmt.Printf(":: Structs ::\n")
	golib.Structs()
	fmt.Printf("\n")

	fmt.Printf(":: Swap ::\n")
	fmt.Println(golib.Swap(a, b))
	fmt.Printf("\n")

	fmt.Printf(":: Switcheroo ::\n")
	golib.Switcheroo()
	fmt.Printf("\n")

	fmt.Printf(":: Variables ::\n")
	golib.Variables()
	fmt.Printf("\n")
}
