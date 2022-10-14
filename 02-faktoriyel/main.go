package main

import "fmt"

func main() {
	var factVal uint64 = 1
	var number int

	fmt.Println("Lütfen 0-65 arasında bir tam sayı giriniz:")
	fmt.Scanf("%d\n", &number)

	if number < 0 {
		fmt.Println("Negatif sayının faktöriyeli olmaz! ")
	} else {
		for i := 1; i <= number; i++ {
			factVal *= uint64(i) // factVal = factVal * uint64(i)
		}
		fmt.Printf("%d sayısının faktöriyeli: %d", number, factVal)
	}

}
