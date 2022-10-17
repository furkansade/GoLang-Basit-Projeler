package main

import "fmt"

func main() {

	var primeNum, count int

	fmt.Println("Kontrol etmek istediğiniz sayı:")
	fmt.Scanln(&primeNum)

	if primeNum <= 0 {
		fmt.Println("Pozitif tam sayı giriniz!")
		return
	}

	for i := 2; i < primeNum/2; i++ {
		if primeNum%i == 0 {
			count++
			break
		}
	}

	if count == 0 && primeNum != 1 {
		fmt.Printf("%d -> Asal!", primeNum)
	} else {
		fmt.Printf("%d -> Asal değil!", primeNum)
	}

}
