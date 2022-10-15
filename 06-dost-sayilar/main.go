package main

import "fmt"

func main() {
	var num1, num2 int
	var num1Sum, num2Sum int

	fmt.Println("Birinci sayı:")
	fmt.Scan(&num1)

	fmt.Println("İkinci sayı:")
	fmt.Scan(&num2)

	for i := 1; i < num1; i++ {
		if num1%i == 0 {
			num1Sum += i
		}
	}

	for j := 1; j < num1; j++ {
		if num2%j == 0 {
			num2Sum += j
		}
	}

	if num1Sum == num2 && num2Sum == num1 {
		fmt.Printf("%d - %d -> Dost Sayılar", num1, num2)
	} else {
		fmt.Printf("%d - %d -> Dost Sayı Değiller", num1, num2)
	}
}
