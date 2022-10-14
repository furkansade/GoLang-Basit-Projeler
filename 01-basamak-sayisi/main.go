package main

import "fmt"

func main() {
	var num, count int

	fmt.Print("Pozitif tam sayı giriniz: ")
	fmt.Scanln(&num)

	for num > 0 {
		num /= 10 // num = num / 10
		count++   // count = count + 1
	}

	fmt.Printf("Girdiğiniz sayının basamak sayısı: %d\n", count)
}
