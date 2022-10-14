package main

import "fmt"

func main() {
	var num, piece int

	fmt.Println("Sayı giriniz:")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			piece += 1
			fmt.Printf("%d | ", i)

		}
	}
	fmt.Println("Pozitif tam bölen sayısı:", piece)
}
