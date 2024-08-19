package main 

import "fmt"

func SwapBits(octet byte) byte {
	return (octet & 0xFE) | ((octet & 0x01) << 1)  // Swap the first and last bits
	 
}

func main() {
	
	fmt.Println(SwapBits(1))
}