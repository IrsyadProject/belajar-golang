package main

import "fmt"

func main() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 21474836488
	const uinteger8 uint8 = 255
	const uinteger16 uint16 = 65535
	const uinteger32 uint32 = 4294967295
	const uinteger64 uint64 = 429496729555
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
	fmt.Println()
	fmt.Println("=== Tipe Data Integer ===")
	fmt.Println()
	fmt.Println("Int8 = ", integer8, "- Uint8 =", uinteger8)
	fmt.Println("Int16 = ", integer16, "- Uint16 =", uinteger16)
	fmt.Println("Int32 = ", integer32, "- Uint32 =", uinteger32)
	fmt.Println("Int64 = ", integer64, "- Uint64 =", uinteger64)

	const tipefloat32 float32 = 3.40928
	const tipefloat64 float64 = 3.409280909099989
	fmt.Println()
	fmt.Println("=== Tipe Data Float ===")
	fmt.Println()
	fmt.Println("Float32 = ", tipefloat32, "- Float64 =", tipefloat64)
}