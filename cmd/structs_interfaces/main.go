package main

import "fmt"

type vehicle struct {
	mileage int
	fuel    int
}

type X1 struct {
	ex1 int
	yEl Y
} // can get name as x1Var.yEl.name

type X2 struct {
	ex2 int
	Y
} // can get name directly as x2Var.name

type Y struct {
	name int
}

// method for struct
// defined like this => assigned to vehicle struct and can be called
// directly as vehicleVar.distCapacity()
func (e vehicle) distCapacity() int {
	return e.mileage * e.fuel
}

// can also write normal functions
func canMakeDistance(e vehicle, miles int) bool {

	// return e.distCapacity() >= miles
	if e.distCapacity() >= miles {
		return true
	} else {
		return false
	}
}

// canMakeDistance can only take in var of type vehicle
// to make it more general, we can use interface
type engine interface {
	distCapacity() int
}

// this can now be added to a function, and it will take in any var which
// has a distCapacity() method, thus making it more general

func canMakeDistance2(e engine, miles int) bool {
	return e.distCapacity() >= miles
}

func main() {

	var myVehicle vehicle                          // values will be default
	fmt.Println(myVehicle.mileage, myVehicle.fuel) // 0 0

	var a vehicle = vehicle{mileage: 10, fuel: 15}
	var b vehicle = vehicle{25, 30} // assigned in order
	fmt.Println(a.fuel, a.mileage)
	fmt.Println(b.fuel, b.mileage)

	var x1 X1 = X1{10, Y{name: 5}}
	var x2 X2 = X2{14, Y{3}}

	fmt.Println(x1.ex1, x1.yEl.name)
	fmt.Println(x2.ex2, x2.name)

	fmt.Println(a.distCapacity())
	fmt.Println(canMakeDistance(a, 149))

	fmt.Println(canMakeDistance2(b, 1000)) // using interface for generality
}
