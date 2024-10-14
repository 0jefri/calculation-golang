package main

import (
	"fmt"
)

func main() {
	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	// Harga jual setelah diskon
	priceAfterDiscount := hargaJual * (1 - diskon/100)

	// Total pendapatan
	totalPendapatan := priceAfterDiscount * float64(jumlahTerjual)

	// Total biaya
	totalBiaya := (hargaBeli + biayaOperasional) * float64(jumlahTerjual)

	// Total keuntungan
	totalKeuntungan := totalPendapatan - totalBiaya

	// Menampilkan hasil menggunakan fmt.Printf
	fmt.Printf("Harga Jual Setelah Diskon: Rp%.2f\n", priceAfterDiscount)
	fmt.Printf("Total Pendapatan: Rp%.2f\n", totalPendapatan)
	fmt.Printf("Total Biaya: Rp%.2f\n", totalBiaya)
	fmt.Printf("Total Keuntungan: Rp%.2f\n", totalKeuntungan)
}
