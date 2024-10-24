# Latihan Modul 4

## Soal Latihan 4.6 (No. 1)
### Program Menghitung Permutasi & Kombinasi
Program difungsikan untuk menghitung dan menampilkan jumlah permutasi dan kombinasi dari dua pasang bilangan yang diberikan user menggunakan prosedur.

1. **Deklarasi variabel dan input user**:
   ```go
   var a, b, c, d, hasil int

   fmt.Scan(&a, &b, &c, &d)
   ```

   Di sini, variabel `a`, `b`, `c`, `d`, dan `hasil` dideklarasikan sebagai `int`, dan program menerima input dari user menggunakan `fmt.Scan()`.

2. **Fungsi `factorial` untuk menghitung faktorial**:
   ```go
   func factorial(n int, hasil *int) {
       var factorial int = 1
       for i := n; i >= 1; i-- {
           factorial *= i
       }
       *hasil = factorial
   }
   ```

   Fungsi diatas menghitung faktorial dari nilai `n` dengan cara mengalikan nilai dari `n` sampai 1, kemudian hasilnya disimpan dalam pass by reference `hasil`.

3. **Fungsi `permutation` untuk menghitung permutasi**:
   ```go
   func permutation(n, r int, hasil *int) {
       var nfactorial, nrfactorial int
       factorial(n, &nfactorial)
       factorial(n-r, &nrfactorial)

       *hasil = nfactorial / nrfactorial
   }
   ```

   Fungsi diatas menghitung permutasi dengan memanggil fungsi `factorial` untuk menghitung `n!` dan `(n-r)!`, lalu membagi hasilnya untuk mendapatkan permutasi.

4. **Fungsi `combination` untuk menghitung kombinasi**:
   ```go
   func combination(n, r int, hasil *int) {
       var nfactorial, rnrfactorial, rnrfactorial2 int
       factorial(n, &nfactorial)
       factorial(r, &rnrfactorial)
       factorial(n-r, &rnrfactorial2)

       *hasil = nfactorial / (rnrfactorial * rnrfactorial2)
   }
   ```

   Fungsi diatas menghitung kombinasi dengan memanggil `factorial` untuk `n!`, `r!`, dan `(n-r)!`, lalu menggunakan rumus kombinasi \( C(n, r) = \frac{n!}{r!(n-r)!} \).

---

## Soal Latihan 4.6 (No. 2)
### Program Penentuan Pemenang Berdasarkan Skor Waktu Pengerjaan Soal
Program difungsikan untuk mencari pemenang berdasarkan jumlah soal yang dijawab benar dalam waktu kurang dari atau sama dengan 300 detik. Jika dua peserta menjawab jumlah soal yang sama, peserta dengan waktu total lebih sedikit yang menang.

### 1. **Deklarasi Konstanta dan Fungsi `hitungSkor`**:
   ```go
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
   ```
   - **Konstanta `jumlahSoal`**: Menentukan jumlah soal tetap, yaitu 8.
   - **Fungsi `hitungSkor`**: Menghitung total soal yang dijawab benar (waktu <= 300 detik) dan menjumlahkan skor total (jumlah waktu yang dihabiskan). Hasilnya disimpan dalam pass by reference `totalSoal` dan `totalSkor`.
   - **Logika dalam `for`**: Jika waktu pengerjaan soal (`soal[i]`) kurang dari atau sama dengan 300, maka soal dihitung sebagai benar dan nilainya ditambahkan ke skor total.

### 2. **Bagian `main`**:
   ```go
   var peserta string
   var waktu [jumlahSoal]int
   var pemenang string
   maxSoal := 0
   minSkor := 301
   totalSkorPemenang := 0
   ```

   - **Variabel**:
     - `peserta`: Menyimpan nama peserta.
     - `waktu`: Array yang menyimpan waktu yang dihabiskan untuk setiap soal.
     - `pemenang`, `maxSoal`, `minSkor`, `totalSkorPemenang`: Variabel untuk melacak siapa peserta terbaik, soal terbanyak yang dijawab, dan skor terendah (total waktu) untuk menentukan pemenang.

### 3. **Pengulangan untuk Input Peserta**:
   ```go
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
   ```

   - **Pengulangan Input**: Program terus meminta input nama peserta dan waktu pengerjaan soal hingga user memasukkan "Selesai".
   - **Mengumpulkan Waktu**: Setiap peserta memberikan waktu pengerjaan untuk 8 soal, yang disimpan dalam array `waktu`.
   - **Memanggil `hitungSkor`**: Menghitung jumlah soal yang dijawab dan skor total untuk setiap peserta.
   - **Logika Pemenang**: Jika peserta saat ini menjawab lebih banyak soal benar dibanding pemenang sebelumnya, atau jika jumlah soal benar sama tapi waktu lebih sedikit, peserta ini dianggap pemenang sementara.

### 4. **Mencetak Hasil Pemenang**:
   ```go
   fmt.Printf("%s %d %d\n", pemenang, maxSoal, totalSkorPemenang)
   ```
   - Setelah pengulangan berakhir, program mencetak nama pemenang, jumlah soal yang dijawab benar, dan skor total (jumlah waktu yang digunakan).

---

## Soal Latihan 4.6 (No. 3)
### Program Sklena
Program bertujuan untuk mencetak deret angka sesuai dengan aturan Sklena, yaitu transformasi bilangan dengan aturan tertentu hingga mencapai angka 1.

### 1. **Konstanta dan Fungsi `cetakDeret`**:
   ```go
   const maxDeret = 100
   ```
   - **`maxDeret`**: Menetapkan batas maksimum elemen dalam array yang akan menyimpan deret angka, yaitu 100 angka.

   ```go
   func cetakDeret(n int) {
       if n <= 0 || n >= 1000000 {
           fmt.Println("Masukkan bilangan positif yang lebih kecil dari 1.000.000")
           return
       }

       var deret [maxDeret]int
       index := 0
   ```
   - **Validasi input**: Jika input `n` kurang dari atau sama dengan 0 atau lebih besar dari 1 juta, program akan menampilkan pesan kesalahan dan keluar dari fungsi.
   - **Deklarasi array `deret`**: Array berukuran 100 ini menyimpan hasil deret Collatz.
   - **`index`**: Variabel ini melacak posisi dalam array saat deret angka dibentuk.

### 2. **Penghitungan Deret**:
   ```go
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
   ```
   - **Perulangan**: Program terus mengubah nilai `n` sampai mencapai 1, sesuai aturan:
     - Jika `n` genap, maka `n` dibagi 2.
     - Jika `n` ganjil, maka `n` dihitung dengan rumus \( n = 3n + 1 \).
   - Hasil setiap perubahan nilai `n` disimpan di dalam array `deret`.
   - Perulangan berhenti saat `n` menjadi 1, lalu angka 1 ditambahkan sebagai elemen terakhir.

### 3. **Mencetak Deret**:
   ```go
   for i := 0; i < index; i++ {
       if i == index-1 {
           fmt.Print(deret[i])
       } else {
           fmt.Print(deret[i], " ")
       }
   }
   fmt.Println()
   ```
   - Setelah deret lengkap terbentuk, program mencetak deret angka yang tersimpan di array `deret` dalam satu baris, dipisahkan oleh spasi.
   - **Pengecekan `i == index-1`**: Ini memastikan angka terakhir dicetak tanpa spasi setelahnya.

### 4. **Bagian `main`**:
   ```go
   func main() {
       var n int
       fmt.Scan(&n)
       cetakDeret(n)
   }
   ```
   - Program meminta input user (`n`), lalu memanggil fungsi `cetakDeret` untuk menghitung dan mencetak deret berdasarkan nilai input.

