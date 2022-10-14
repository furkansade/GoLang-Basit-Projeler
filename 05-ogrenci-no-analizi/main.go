package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/* Ornek Cikti!
201905030821
Giriş Yılı -> 2019
Fakülte No -> 05
Bölüm No -> 03
Öğrenci No -> 08
Giriş Sırası -> 21 */

func main() {
	var stdNum int
	fmt.Println("Lütfen 11 haneli okul numaranızı girin:")
	fmt.Scanf("%d\n", &stdNum)

	strNum := strconv.Itoa(stdNum) //strconv paketi int -> string icin kullandik.

	// utf8.RuneCountInString() karakter sayısı paketi.
	if utf8.RuneCountInString(strNum) != 12 {
		fmt.Println("Lütfen 12 haneli öğrenci numaranızı giriniz!")
		return
	}

	var studentInfo map[string]string // tek tek degiskene atayip yapabilirdik.

	studentInfo = map[string]string{
		"Giriş Yılı":   strNum[:4],
		"Fakülte No":   strNum[4:6],
		"Bölüm No":     strNum[6:8],
		"Öğrenci No":   strNum[8:10],
		"Giriş Sırası": strNum[10:],
	}

	for i, v := range studentInfo {
		fmt.Printf("%v -> %s\n", i, v)
	}

}
