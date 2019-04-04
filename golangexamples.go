//Package golangexamples that show the go examples
package golangexamples

import "github.com/ehteshamz/greetings"

// ConcatSlice is a function contents of the slice concatenated together
func ConcatSlice(sliceToConcat []byte) string {

	var concatenatedString = ""
	for i := 0; i < len(sliceToConcat); i++ {
		concatenatedString += string(sliceToConcat[i])
		if i != len(sliceToConcat)-1 {
			concatenatedString += "-"
		}
	}
	return concatenatedString
}

//Encrypt is a function that encrypts the slice using ceaser cypher
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {

	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] += byte((ceaserCount))
	}
	for i := 0; i < len(sliceToEncrypt); i++ {
		print(string(sliceToEncrypt[i]))
	}
}

//EZGreetings is a function that return the greetings
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
