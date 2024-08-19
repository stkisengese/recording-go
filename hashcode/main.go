package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

//If the resulting character is unprintable add 33 to it.

func HashCode(dec string) string {
	var hash int
	result := ""
    for _, char := range dec {
        hash = (int(char) + len(dec)) % 127
        if hash < 32 {
            hash += 33
        }
		// result += fmt.Sprintf("%c", hash)
		result += string(rune(hash))
    }
    return result
}