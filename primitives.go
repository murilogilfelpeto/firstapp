package main

import "fmt"

func main() {
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	test1 := 1 == 1
	test2 := 1 == 2
	fmt.Printf("%v, %T\n", test1, test1)
	fmt.Printf("%v, %T\n", test2, test2)

	//Numeric types
	var a int = 42                    // int is 32 or 64 bit depending on the OS
	var b int8 = 127                  // int8 is 8 bit -128 to 127
	var c int16 = 32767               // int16 is 16 bit -32768 to 32767
	var d int32 = 2147483647          // int32 is 32 bit -2147483648 to 2147483647
	var e int64 = 9223372036854775807 // int64 is 64 bit -9223372036854775808 to 9223372036854775807
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	// We have the uint type too, that starts from 0
	var f uint = 42                     // uint is 32 or 64 bit depending on the OS
	var g uint8 = 255                   // uint8 is 8 bit 0 to 255
	var h uint16 = 65535                // uint16 is 16 bit 0 to 65535
	var i uint32 = 4294967295           // uint32 is 32 bit 0 to 4294967295
	var j uint64 = 18446744073709551615 // uint64 is 64 bit 0 to 18446744073709551615
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	//Math operations
	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	//Go does not allow us to mix types in operations,
	//so if we divide two ints the result will be an int
	//If we want to divide two ints and get a float we need to convert both of them
	fmt.Println(float32(x) / float32(y))

	//Bit operations
	//Bitwise AND
	fmt.Println(x & y)
	//Bitwise OR
	fmt.Println(x | y)
	//Bitwise XOR
	fmt.Println(x ^ y)
	//Bitwise AND NOT
	fmt.Println(x &^ y)

	//Bit shifting
	//Left shift
	fmt.Println(x << 1)
	//Right shift
	fmt.Println(x >> 1)

	//Complex numbers
	var k complex64 = 1 + 2i
	var l complex64 = 2 + 5.2i
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", real(k), real(k))
	fmt.Printf("%v, %T\n", imag(k), imag(k))
	fmt.Printf("%v, %T\n", k+l, k+l)
	fmt.Printf("%v, %T\n", k-l, k-l)
	fmt.Printf("%v, %T\n", k*l, k*l)
	fmt.Printf("%v, %T\n", k/l, k/l)

}
