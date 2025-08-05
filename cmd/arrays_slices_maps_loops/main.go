package main

import "fmt"

func main() {

	var array [3]int

	array[1] = 123

	// arrays stored in contiguous memory locations
	fmt.Println(array[0])
	fmt.Println(array[1:3]) // [1 2] indexes

	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3} // size inferred from spread operator
	fmt.Println(arr1, arr2, arr3)

	// slices are just wrappers around arrays
	// just omit the length field to get a slice
	var slice1 []int = []int{1, 2, 5}
	fmt.Println(slice1)
	slice1 = append(slice1, 7)
	// when 7 appended, the capacity of slice went from 3 to 6, but we cannot access indexes 4 and 5
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(slice1)

	var slice2 []int = []int{6, 7, 8}
	slice1 = append(slice1, slice2...) // spread operator to append multiple values
	fmt.Println(slice2, slice1)

	// can use make function to create slice, and optionally also provide capacity
	// default capacity is just length of slice
	var slice3 []int = make([]int, 3)
	fmt.Println(slice3, len(slice3), cap(slice3))

	var slice4 []int = make([]int, 3, 8)
	fmt.Println(slice4, len(slice4), cap(slice4))

	//maps -> (key: value) pairs
	var map1 map[string]int = make(map[string]int)
	fmt.Println(map1)

	var map2 map[string]int = map[string]int{"Rujool": 22, "Jason": 21}
	fmt.Println(map2)

	fmt.Println(map2["Rujool"])
	fmt.Println(map2["anon"]) // maps return default value for value if key not present

	val1, doesExist1 := map2["anon"] // also return an optional bool value which is true if key present
	fmt.Println(val1, doesExist1)

	val2, doesExist2 := map2["Jason"]
	delete(map2, "Jason") // deletes given key from map
	// deletes by reference, so no return value given
	val3, doesExist3 := map2["Jason"]
	fmt.Println(val2, doesExist2) // 21 true
	fmt.Println(val3, doesExist3) // 0 false

	map2["Tony"] = 45
	// for loop over map
	for name := range map2 {
		fmt.Println(name, map2[name])
	} // no order preserved when iterating over a map

	// for loop over map other way
	for name, value := range map2 {
		fmt.Println(name, value)
	}

	fmt.Println(slice2)
	// for loop over array/slice
	for idx, val := range slice2 {
		fmt.Println(idx, val)
	}

	for _, val := range slice2 {
		fmt.Println(val)
	}

	for i := range slice2 {
		fmt.Println(i, slice2[i])
	}

	// go doesnt have a while loop
	// but this is how you do it using the for keyword
	var i int = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	var k int = 0
	for {
		if k > 5 {
			break
		}
		fmt.Println(k)
		k++
	}

}
