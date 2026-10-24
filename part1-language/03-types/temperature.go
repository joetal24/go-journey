
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var s string
	var i int
	var f float64
	var b bool
	var sl []int
	var m map[string]int
	var p *int
	var person Person

	fmt.Printf("%-10T %q\n", s, s)
	fmt.Printf("%-10T %d\n", i, i)
	fmt.Printf("%-10T %g\n", f, f)
	fmt.Printf("%-10T %t\n", b, b)
	fmt.Printf("%-10T %v\n", sl, sl)
	fmt.Printf("%-10T %v\n", m, m)
	fmt.Printf("%-10T %v\n", p, p)
	fmt.Printf("%-10T %+v\n", person, person)

	// Usable immediately:
	// string, int, float64, bool, and a nil slice.
	// A nil slice can be appended to.

	// Not usable for writing without initialization:
	// A nil map panics when assigning a key.
	// A nil pointer panics when dereferenced.
	// A struct's zero value is usable, but its fields
	// contain their respective zero values.
}
