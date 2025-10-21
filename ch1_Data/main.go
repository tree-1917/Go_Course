package main

import "fmt"

func main() {
	// ============================================================
	// Section 1: Integers
	// ============================================================

	// -------------------------------
	// Section 1-A: Signed Integers
	// -------------------------------
	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)

	// Assigning maximum values
	i = int(^uint(0) >> 1)    // Max int (depends on system architecture)
	i8 = 127                  // Max int8  =  2^7  - 1
	i16 = 32767               // Max int16 =  2^15 - 1
	i32 = 2147483647          // Max int32 =  2^31 - 1
	i64 = 9223372036854775807 // Max int64 =  2^63 - 1

	fmt.Println("=== Signed Integers ===")
	fmt.Printf("int   : %d\n", i)
	fmt.Printf("int8  : %d (Min: %d)\n", i8, -128)
	fmt.Printf("int16 : %d (Min: %d)\n", i16, -32768)
	fmt.Printf("int32 : %d (Min: %d)\n", i32, -2147483648)
	fmt.Printf("int64 : %d (Min: %d)\n", i64, -9223372036854775808)

	// -------------------------------
	// Section 1-B: Unsigned Integers
	// -------------------------------
	var (
		u   uint
		u8  uint8
		u16 uint16
		u32 uint32
		u64 uint64
	)

	// Assigning maximum values
	u = ^uint(0)               // Max uint (depends on system architecture)
	u8 = 255                   // Max uint8  =  2^8  - 1
	u16 = 65535                // Max uint16 =  2^16 - 1
	u32 = 4294967295           // Max uint32 =  2^32 - 1
	u64 = 18446744073709551615 // Max uint64 =  2^64 - 1

	fmt.Println("\n=== Unsigned Integers ===")
	fmt.Printf("uint   : %d\n", u)
	fmt.Printf("uint8  : %d\n", u8)
	fmt.Printf("uint16 : %d\n", u16)
	fmt.Printf("uint32 : %d\n", u32)
	fmt.Printf("uint64 : %d\n", u64)

	// ============================================================
	// Section 1: Floating Point
	// ============================================================
	var f32 float32 = 10.5
	var f64 float64 = 10.000123123

	fmt.Println("\n=== Floating Point ===")
	fmt.Printf("float32   : %f\n", f32)
	fmt.Printf("float64  : %f\n", f64)

	// ============================================================
	// Section 1: Boolean
	// ============================================================
	var isActive bool = true
	var isOn bool = false

	fmt.Println("\n=== Boolean ===")
	fmt.Printf("isActive   : %t\n", isActive)
	fmt.Printf("isOn 	   : %t\n", isOn)

	// ============================================================
	// Section 1: Complex
	// ============================================================
	var cn64 complex64 = complex(5, 10)
	var cn128 complex128 = complex(10, 5)

	fmt.Println("\n===  Point ===")
	fmt.Printf("complex64  : %v\n", cn64)
	fmt.Printf("complex128 	   : %v\n", cn128)
}
