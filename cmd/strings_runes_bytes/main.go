package main

import (
	"fmt"
	"strings"
)

func main() {
	var mystr string = "résumé"
	var indexed = mystr[0]                  // can index string like an array
	fmt.Printf("%v %T\n", indexed, indexed) // %v is for value, %t is for type
	// will give 114(ascii of r) and uint8

	for i, v := range mystr {
		fmt.Printf("%v %v\n", i, v)
		// skips index2, and prints ascii values for each
	}

	// this is because accented e is 233, which uses 2 bytes for representation in utf8
	// 0s padded to left to fit in utf8 representation
	// range keyword gives correct value and skips index 2
	// if we just did mystr[1], it would give 195 (first half i.e first byte of accented e)
	fmt.Printf("%v\n", mystr[1])

	// takeaway -> easier way is just casting a string into an array of runes
	// instead of dealing with underlying byte array (which is what string is)

	var runestr []rune = []rune("résumé")
	fmt.Println(runestr, runestr[1])
	for i, v := range runestr {
		fmt.Printf("%v %v\n", i, v)
		// gives correct values, doesnt skip any index
	}

	// runes are just an alias for int32
	var myrune rune = 'a' // single quotes can also be used to declare runes
	fmt.Println(myrune)

	// to print as characters, use %c format specifier
	fmt.Printf("%v %c %v %c\n", myrune, myrune, runestr[1], runestr[1])

	// string building
	var strSlice = []string{"a", "p", "p", "l", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}

	fmt.Println(strSlice, catStr)
	// gives [a p p l e] apple
	// strings in go are immutable
	// when we concatenate a string and assign it to a new variable, we are creating a new string
	// everytime, which is inefficient

	// just use the strings package and stringbuilder
	// this will allocate internal array, and append values
	// new string created only at the end
	// this is much faster and efficinet
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStrr = strBuilder.String()
	fmt.Println(catStrr)

}
