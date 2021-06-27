package main

import "fmt"

func main () {
	type ktp int
	type kota string

	var noktp ktp = 101020
	var asal kota = "surabaya"
	fmt.Println(noktp)
	fmt.Println(asal)
}