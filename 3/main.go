package main

import "fmt"

const maxDeret = 100

func cetakDeret(n int) {
	if n <= 0 || n >= 1000000 {
		fmt.Println("Masukkan bilangan positif yang lebih kecil dari 1.000.000")
		return
	}

	var deret [maxDeret]int
	index := 0

	for n != 1 {
		deret[index] = n
		index++

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	deret[index] = 1
	index++

	for i := 0; i < index; i++ {
		if i == index-1 {
			fmt.Print(deret[i])
		} else {
			fmt.Print(deret[i], " ")
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
