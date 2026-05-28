package main
import "fmt"

// Devia - tampilan menu
func menu() {
	fmt.Println("\n+++ APOTEK-SMART +++")
	fmt.Println("1. Tambah Data Obat")
	fmt.Println("2. Ubah Data Obat")
	fmt.Println("3. Hapus Data Obat")
	fmt.Println("4. Tampilkan Data")
	fmt.Println("5. Cari Obat")
	fmt.Println("6. Urutkan Data")
	fmt.Println("7. Statistik Obat")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

// Devia - Statistik Obat
func statistikObat() {

	fmt.Println("\n+++ APOTEK-SMART +++")

	fmt.Println("\nObat Hampir Habis:")

	for i := 0; i < len(dataObat); i++ {

		if dataObat[i].stok <= 5 {
			fmt.Println("-", dataObat[i].nama, "| Stok:", dataObat[i].stok)
		}
	}

	fmt.Println("\nObat Segera Kadaluarsa:")

	for i := 0; i < len(dataObat); i++ {
		fmt.Println("-", dataObat[i].nama, "| Exp:", dataObat[i].kadaluarsa)
	}
}

// Devia - main program
func main() {

	var pilih int
	var nama string

	for {

		menu()
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahObat()

		case 2:
			ubahObat()

		case 3:
			hapusObat()

		case 4:
			tampilData()

		case 5:
			fmt.Print("Masukkan nama obat: ")
			fmt.Scan(&nama)
			sequentialSearch(nama)

		case 6:
			selectionSort()

		case 7:
			statistikObat()

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}