package main

import (
	"fmt"
	"strconv"
)

// Numero is a custom type definition uses int.
type Numero int

// String implements Stringer interface.
func (n Numero) String() string {
	return strconv.Itoa(int(n))
}

// Format implements Formatter interface.
func (n Numero) Format(f fmt.State, verb rune) {
	val := n.String() // get string version
	if verb == 81 {   // check if Q passed
		val = "*" + val + "*" // add quotes
	}
	fmt.Fprint(f, val)
}

func main() {
	a := Numero(1)

	fmt.Printf("number is: %Q\n", a) // number is: "1"
}
