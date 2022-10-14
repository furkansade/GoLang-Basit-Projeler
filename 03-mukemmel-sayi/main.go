package main

import "fmt"

func main() {

	var num, sum int
	fmt.Println("Lütfen pozitif tam sayı giriniz: ")
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	if sum == (num * 2) {
		fmt.Printf("%d mükemmeldir.", num)
	} else {
		fmt.Printf("%d mükemmel değildir.", num)
	}

}
