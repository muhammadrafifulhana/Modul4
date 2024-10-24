package main

import "fmt"

const jumlahSoal = 8

func hitungSkor(nama string, soal [jumlahSoal]int, totalSoal *int, totalSkor *int) {
	*totalSoal = 0
	*totalSkor = 0
	for i := 0; i < jumlahSoal; i++ {
		if soal[i] <= 300 {
			*totalSoal++
			*totalSkor += soal[i]
		}
	}
}

func main() {
	var peserta string
	var waktu [jumlahSoal]int
	var pemenang string
	maxSoal := 0
	minSkor := 301
	totalSkorPemenang := 0

	for {
		fmt.Scan(&peserta)
		if peserta == "Selesai" {
			break
		}

		for i := 0; i < jumlahSoal; i++ {
			fmt.Scan(&waktu[i])
		}

		var totalSoal, totalSkor int
		hitungSkor(peserta, waktu, &totalSoal, &totalSkor)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			maxSoal = totalSoal
			minSkor = totalSkor
			pemenang = peserta
			totalSkorPemenang = totalSkor
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, totalSkorPemenang)
}
