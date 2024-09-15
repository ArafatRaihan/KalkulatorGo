package main

import (
	"fmt"
)

func main() {
	var operator string
	var angka1, angka2 float64

	fmt.Println("Masukkan operasi (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Println("Masukkan bilangan pertama:")
	fmt.Scanln(&angka1)
	fmt.Println("Masukkan bilangan kedua:")
	fmt.Scanln(&angka2)

	switch operator {

	case "+":
		fmt.Printf("Hasil: %.2f\n", angka1+angka2)

	case "-":
		fmt.Printf("Hasil: %.2f\n", angka1-angka2)

	case "*":
		fmt.Printf("Hasil: %.2f\n", angka1*angka2)

	case "/":
		if angka2 != 0 {
			fmt.Printf("Hasil: %.2f\n", angka1/angka2)
		} else {
			fmt.Println("Error: Tidak bisa membagi dengan 0")
		}
	default:
		fmt.Println("Kalkulator tidak mengerti input user")
	}
}
