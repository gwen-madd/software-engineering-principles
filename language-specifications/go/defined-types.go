package main

//see - https://go.dev/ref/spec#Types
import "fmt"

func main() {
	var boolVar bool = false

	//print the type and value, as seen here - https://go.dev/tour/basics/11
	fmt.Printf("This is a bool (short for boolean) of type: %T and value: %v\n", boolVar, boolVar)
	fmt.Println("bools can have one of two values [true, false]")

	var aByte uint8 = 255
	var twoBytes uint16 = 65530

	fmt.Printf("This is a uint8 (short for unsigned integer of 8-bit capacity) of type: %T and value: %v\n", aByte, aByte)
	fmt.Println("uint8 can represent the following set of non-negative integers [0,255]")

	fmt.Printf("This is a uint16 (short for unsigned integer of 16-bit capacity) of type: %T and value: %v\n", twoBytes, twoBytes)
	fmt.Println("uint16 can represent the following set of non-negative integers [0,65535]")
	fmt.Println("This pattern continues for uint32 and uint64 where the range can be calculated by converting from binomial representation of the byters allocated.")
	fmt.Println("For example, uint32 can store values in range [0, (2^32)-1] and uint64 - [0, (2^64)-1]")
	fmt.Println("Signed integers follow similar convention, however, the leftmost digit holds value of negative 2^n  for the conversion from binary to decimal representation")
	fmt.Println("Signed integers follow similar convention, have a different range, for example int8 captures values [-128,127] and you may calculate these ranges in a similar fashion [-(2^(n-1)), 2^(n-1) - 1]")
	fmt.Println("This is commonly referred to as Two's complement as detailed here: https://en.wikipedia.org/wiki/Two's_complement")

	var aRationalNumber float32 = 3.8
	var aComplexNumber complex64 = 3.8 + 0.1i

	fmt.Printf("This is a float32 (short for signed floating point number of 32-bit capacity) of type: %T and value: %v\n", aRationalNumber, aRationalNumber)
	fmt.Printf("This is a complex64 (short for signed complex number of 64-bit capacity) of type: %T and value: %v\n", aComplexNumber, aComplexNumber)

	fmt.Println("Floating point numbers use the IEEE-754 2008 Revision specification for 32 bit and 64 bit numbers respectively")
	fmt.Println("Complex numbers follow the same convention as floating point numbers but are broken into two parts, real and imaginary. For example, 32 bits in a complex 64 are used for the real part while the remaining 32 bits are used for the imaginary part of the number")
}
