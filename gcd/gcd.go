package main

import "fmt"

func SteinGcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b uint) uint {
	return (a * b) / SteinGcd(a, b)
}

// // Stein's algorithm
func Gcd(a, b uint) uint {
	if a == b {
		return a
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a&1 == 0 && b&1 == 0 {
		return Gcd(a>>1, b>>1) << 1 // GCD is 2 times the GCD of a/2 and b/2
	}
	if a&1 == 0 {
		return Gcd(a>>1, b)
	}
	if b&1 == 0 {
		return Gcd(a, b>>1)
	}
	if a > b {
		return Gcd((a-b)>>1, b)
	}
	return Gcd((b-a)>>1, a)
}

// // Binary GCD
func binaryGcd(a, b uint64) uint64 {
	if a == 0 || b == 0 {
		return a | b
	}

	shift := 0
	for (a&1) == 0 && (b&1) == 0 { // if both numbers are even
		a >>= 1 // divide a by 2
		b >>= 1 // divide b by 2
		shift++
	}
	for a != b {
		if a&1 == 0 { // if a is even, divide a by 2
			a >>= 1
		} else if b&1 == 0 { // if b is even, divide b by 2
			b >>= 1
		} else if a > b {
			a -= b
			// a = (a - b) >> 1      // subtract  b from a
		} else {
			b -= a
			// b = (b - a) >> 1   // subtract a from b
		}
	}
	return a << shift // GCD is non zero number multiplied by two
}

// // Euclidean algorithm
// func Gcd(a, b uint) uint {
// 	if a == 0 || b == 0 {
//         return 0
//     }
// 	for a != b {
// 		if a > b {
// 			a -= b
// 		} else {
// 			b -= a
// 		}
// 	}
// 	return a
// }

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
	fmt.Println(SteinGcd(42, 12))
	fmt.Println(binaryGcd(14, 77))
	fmt.Println(binaryGcd(17, 3))

	fmt.Println(Lcm(42, 10))
	fmt.Println(Lcm(14, 77))
	fmt.Println(Lcm(17, 3))
	fmt.Println(SteinGcd(42, 12))
	fmt.Println(binaryGcd(14, 77))
	fmt.Println(Lcm(17, 3))
}
