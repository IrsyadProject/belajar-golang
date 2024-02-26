package main

import "fmt"

func main() {
	var tipestring string = "Halo Irsyad"
	const data2 = "Nama saya Irsyad"
	fmt.Println("=== Tipe Data String ===")
	fmt.Println()
	fmt.Println("String =", tipestring)
	fmt.Println("String 2 =", data2)
	fmt.Println()
	fmt.Println("=== Fungsi pad Tipe Data String ===")
	fmt.Println()
	// Menghitung jumlah karakter di String
	fmt.Println("Fungsi Len =", len(tipestring))
	// Mengambil karakter pada posisi yang ditentukan
	fmt.Println("Fungsi kedua =", data2[0])
	fmt.Println("Fungsi kedua =", data2[1])
	fmt.Println("Fungsi kedua =", data2[2])

}