package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	word := computer(a)
	fmt.Println(a, word)
}

func computer(a int) (b string) {
	switch {
	case a%10 == 1 && a != 11:
		b = " компьютер"
		return
	case a > 1 && a < 5:
		b = " компьютера"
	case a%10 > 1 && a%10 < 5 && a > 20:
		b = " компьютера"
		return
	case a > 4:
		b = " компьютеров"
		return
	}
	return
}
