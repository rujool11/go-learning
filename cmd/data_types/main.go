package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 5
	// int, int8, int16, int32, int64 exist, if undeclared default to 0
	// uint, unit8, uint16, uint32, uint64 also exist (unsigned int)
	// int default is machine dependent, usually int64
	fmt.Println(intNum)

	var floatNum float64 = 5.677
	// float32 and float64 exist (no default float)
	fmt.Println(floatNum)

	// typecasting ->
	var int32Num int32 = 45
	var float32Num float32 = 56.5454
	var result float32 = float32Num + float32(int32Num)
	fmt.Println(result)

	// int division results in int and rounds down
	var intNum1 int = 14
	var intNum2 int = 5
	fmt.Println(intNum1 / intNum2) // 14/5 = 2.8 = 2
	fmt.Println(intNum1 % intNum2) // modulus

	var stringData1 string = "Hello \nworld" // if using "", cannot continue string to next line
	var stringData2 string = `Hello
	world` // if using `` can continue to next line
	fmt.Println(stringData1)
	fmt.Println(stringData2)
	var myString string = "Hello" + " " + "world" // concatenation
	fmt.Println(myString)

	// len gives number of bytes in string, UTF-8
	// so len("A") = 1, but len("γ") = 2, because gamma symbol uses more bytes
	var a string = "A"
	var gamma string = "γ"
	fmt.Println(len(a))
	fmt.Println(len(gamma))
	// for complex strings, to get number of characters, use this->
	fmt.Println(utf8.RuneCountInString(gamma))

	// runes are a data type that represent chars
	// bit weird, and not used as much
	var rune_a rune = 'a'
	fmt.Println(rune_a) // this will print 97

	var myBool bool = true
	fmt.Println(myBool)

	// default values ->
	// all ints, units, floats, runes are 0 by default
	// strings are "" (empty string by default)
	// booleans are false by default

	// can omit type by setting variable right away (inferred type)
	var inferredStr = "apple"
	fmt.Println(inferredStr)

	// shorthand can be used to omit var keyword
	shorthandStr := "banana"
	fmt.Println(shorthandStr)

	// multi-variable initilization ->
	var x1, x2 int = 3, 4
	fmt.Println(x1, x2)

	y1, y2 := "orange", "mango"
	fmt.Println(y1, y2)

	const myConst string = "constant_string"
	fmt.Println(myConst)
	// same as vars, but will not change, and need to be instantly initialized
}
