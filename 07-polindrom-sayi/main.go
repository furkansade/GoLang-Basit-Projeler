package main

import "fmt"

func main() {
	var num, rem, temp int // number, remainder, temp
	var rev int = 0        // reverse

	fmt.Print("Pozitif tam sayı giriniz: ")
	fmt.Scan(&num)

	temp = num

	for {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10

		if num == 0 {
			break
		}
	}

	if temp == rev {
		fmt.Printf("%d -> Polindrom!", temp)
	} else {
		fmt.Printf("%d -> Polindrom Değil!", temp)
	}
}
