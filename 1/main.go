package main

import "fmt"

func main() {

	var a, b, c, d, hasil int

	fmt.Scan(&a, &b, &c, &d)
	permutation(a, c, &hasil)
	fmt.Print(hasil, " ")

	combination(a, c, &hasil)
	fmt.Println(hasil)

	permutation(b, d, &hasil)
	fmt.Print(hasil, " ")

	combination(b, d, &hasil)
	fmt.Println(hasil)
}

func factorial(n int, hasil *int) {

	var factorial int = 1
	for i := n; i >= 1; i-- {
		factorial *= i
	}
	*hasil = factorial
}

func permutation(n, r int, hasil *int) {

	var nfactorial, nrfactorial int
	factorial(n, &nfactorial)
	factorial(n-r, &nrfactorial)

	*hasil = nfactorial / nrfactorial

}

func combination(n, r int, hasil *int) {

	var nfactorial, rnrfactorial, rnrfactorial2 int
	factorial(n, &nfactorial)
	factorial(r, &rnrfactorial)
	factorial(n-r, &rnrfactorial2)

	*hasil = nfactorial / (rnrfactorial * rnrfactorial2)
}
