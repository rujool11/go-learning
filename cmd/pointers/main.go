package main

import "fmt"

func main() {
	// var p1 *int // uninitialized pointer, value nil
	var p *int = new(int) // gives a free memory location and assigns it to p2
	var i int

	fmt.Printf("%v %v\n", *p, i)
	*p = 10
	fmt.Printf("%v %v\n", *p, i)

	//*p1 = 5 // gives exception because p1 is nil

	p = &i // p points to address of i now
	i = 9
	fmt.Printf("%v %v\n", *p, i)
}
