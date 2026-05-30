/*
 * ============================================================
 *   APOTEK-SMART
 *   Bagian: Mirza
 *   Spesifikasi A & B:
 *     - Struct dan Variabel Global
 *     - Fungsi Utilitas Input
 *     - CRUD Data Obat
 *     - CRUD Kategori Gejala
 *     - Pencatatan Stok Masuk
 * ============================================================
 *   Kode ekstra: eel | akun: @jebb_24
 * ============================================================
 */

package main

import "fmt"

// ============================================================
// KONSTANTA
// ============================================================

const (
	MaxObat     = 100
	MaxKategori = 20
	StokMinimum = 10
)

// ============================================================
// STRUCT / TIPE DATA
// ============================================================

type Tanggal struct {
	Hari  int
	Bulan int
	Tahun int
}

type Obat struct {
	ID         int
	Nama       string
	Kategori   string
	Indikasi   string
	Stok       int
	Harga      float64
	Kadaluarsa Tanggal
}

type KategoriGejala struct {
	ID        int
	Nama      string
	Deskripsi string
}

// ============================================================
// VARIABEL GLOBAL
// ============================================================

var daftarObat [MaxObat]Obat
var daftarKategori [MaxKategori]KategoriGejala
var jumlahObat int = 0
var jumlahKategori int = 0
var nextIDObat int = 1
var nextIDKategori int = 1

// ============================================================
// FUNGSI UTILITAS INPUT
// ============================================================

func bacaString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func bacaInt(prompt string) int {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err == nil {
			return input
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka!")
		var discard string
		fmt.Scanln(&discard)
	}
}

func bacaFloat(prompt string) float64 {
	var input float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err == nil {
			return input
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka!")
		var discard string
		fmt.Scanln(&discard)
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func getDaysInMonth(month int, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

func tanggalValid(t Tanggal) bool {
	if t.Tahun < 2024 || t.Tahun > 2100 {
		return false
	}
	if t.Bulan < 1 || t.Bulan > 12 {
		return false
	}
	if t.Hari < 1 || t.Hari > getDaysInMonth(t.Bulan, t.Tahun) {
		return false
	}
	return true
}

func bacaTanggal() Tanggal {
	var t Tanggal
	for {
		fmt.Println("    (Hari/Bulan/Tahun)")
		t.Hari = bacaInt("    Hari   : ")
		t.Bulan = bacaInt("    Bulan  : ")
		t.Tahun = bacaInt("    Tahun  : ")
		if tanggalValid(t) {
			break
		}
		fmt.Println("  [!] Tanggal tidak valid, coba lagi.")
	}
	return t
}

func formatTanggal(t Tanggal) string {
	return fmt.Sprintf("%02d/%02d/%04d", t.Hari, t.Bulan, t.Tahun)
}

// ============================================================
// CRUD KATEGORI GEJALA - SPESIFIKASI A
// ============================================================

func tampilkanKategori() {
	fmt.Println("\n  DAFTAR KATEGORI GEJALA PENYAKIT")
	fmt.Println("  ----------------------------------------")
	if jumlahKategori == 0 {
		fmt.Println("  [!] Belum ada kategori yang terdaftar.")
		return
	}
	fmt.Printf("  %-4s %-25s %-35s\n", "ID", "Nama Kategori", "Deskripsi")
	fmt.Println("  ----------------------------------------")
	for i := 0; i < jumlahKategori; i++ {
		k := daftarKategori[i]
		fmt.Printf("  %-4d %-25s %-35s\n", k.ID, k.Nama, k.Deskripsi)
	}
}

func tambahKategori() {
	if jumlahKategori >= MaxKategori {
		fmt.Printf("  [!] Kapasitas kategori sudah penuh (%d).\n", MaxKategori)
		return
	}
	var k KategoriGejala
	k.ID = nextIDKategori
	nextIDKategori++
	fmt.Print("  Nama Kategori : ")
	fmt.Scanln(&k.Nama)
	fmt.Print("  Deskripsi     : ")
	fmt.Scanln(&k.Deskripsi)
	daftarKategori[jumlahKategori] = k
	jumlahKategori++
	fmt.Printf("  [OK] Kategori '%s' berhasil ditambahkan.\n", k.Nama)
}

func ubahKategori() {
	tampilkanKategori()
	if jumlahKategori == 0 {
		return
	}
	id := bacaInt("\n  ID kategori yang ingin diubah: ")
	for i := 0; i < jumlahKategori; i++ {
		if daftarKategori[i].ID == id {
			var tmp string
			fmt.Print("  Nama baru      : ")
			fmt.Scanln(&tmp)
			if tmp != "" {
				daftarKategori[i].Nama = tmp
			}
			fmt.Print("  Deskripsi baru : ")
			fmt.Scanln(&tmp)
			if tmp != "" {
				daftarKategori[i].Deskripsi = tmp
			}
			fmt.Println("  [OK] Kategori berhasil diperbarui.")
			return
		}
	}
	fmt.Printf("  [!] Kategori dengan ID %d tidak ditemukan.\n", id)
}

func hapusKategori() {
	tampilkanKategori()
	if jumlahKategori == 0 {
		return
	}
	id := bacaInt("\n  ID kategori yang ingin dihapus: ")
	for i := 0; i < jumlahKategori; i++ {
		if daftarKategori[i].ID == id {
			var konfirmasi string
			fmt.Printf("  Yakin hapus '%s'? (y/n): ", daftarKategori[i].Nama)
			fmt.Scan(&konfirmasi)
			if konfirmasi != "y" && konfirmasi != "Y" {
				fmt.Println("  [!] Penghapusan dibatalkan.")
				return
			}
			for j := i; j < jumlahKategori-1; j++ {
				daftarKategori[j] = daftarKategori[j+1]
			}
			jumlahKategori--
			fmt.Println("  [OK] Kategori berhasil dihapus.")
			return
		}
	}
	fmt.Printf("  [!] Kategori ID %d tidak ditemukan.\n", id)
}

// ============================================================
// CRUD DATA OBAT - SPESIFIKASI A
// ============================================================

func tampilkanObat() {
	fmt.Println("\n  DAFTAR STOK OBAT")
	fmt.Println("  ---------------------------------------------------------------------------------------")
	if jumlahObat == 0 {
		fmt.Println("  [!] Belum ada data obat yang terdaftar.")
		return
	}
	fmt.Printf("  %-4s %-20s %-15s %-20s %-6s %-12s %-12s\n",
		"ID", "Nama Obat", "Kategori", "Indikasi", "Stok", "Harga", "Kadaluarsa")
	fmt.Println("  ---------------------------------------------------------------------------------------")
	for i := 0; i < jumlahObat; i++ {
		o := daftarObat[i]
		label := ""
		if o.Stok <= StokMinimum {
			label = " [STOK RENDAH]"
		}
		fmt.Printf("  %-4d %-20s %-15s %-20s %-6d Rp%8.0f %-12s%s\n",
			o.ID, o.Nama, o.Kategori, o.Indikasi,
			o.Stok, o.Harga, formatTanggal(o.Kadaluarsa), label)
	}
}

func tambahObat() {
	if jumlahObat >= MaxObat {
		fmt.Printf("  [!] Kapasitas data obat penuh (%d).\n", MaxObat)
		return
	}
	var o Obat
	o.ID = nextIDObat
	nextIDObat++
	fmt.Print("  Nama Obat    : ")
	fmt.Scanln(&o.Nama)
	fmt.Print("  Kategori     : ")
	fmt.Scanln(&o.Kategori)
	fmt.Print("  Indikasi     : ")
	fmt.Scanln(&o.Indikasi)
	o.Stok = bacaInt("  Stok Awal    : ")
	o.Harga = bacaFloat("  Harga Satuan : Rp ")
	fmt.Println("  Tanggal Kadaluarsa:")
	o.Kadaluarsa = bacaTanggal()
	daftarObat[jumlahObat] = o
	jumlahObat++
	fmt.Printf("  [OK] Obat '%s' berhasil ditambahkan (ID: %d).\n", o.Nama, o.ID)
}

func ubahObat() {
	tampilkanObat()
	if jumlahObat == 0 {
		return
	}
	id := bacaInt("\n  ID obat yang ingin diubah: ")
	for i := 0; i < jumlahObat; i++ {
		if daftarObat[i].ID == id {
			var tmp string
			fmt.Println("  (Kosongkan untuk melewati)")
			fmt.Print("  Nama Obat    : ")
			fmt.Scanln(&tmp)
			if tmp != "" {
				daftarObat[i].Nama = tmp
			}
			fmt.Print("  Kategori     : ")
			fmt.Scanln(&tmp)
			if tmp != "" {
				daftarObat[i].Kategori = tmp
			}
			fmt.Print("  Indikasi     : ")
			fmt.Scanln(&tmp)
			if tmp != "" {
				daftarObat[i].Indikasi = tmp
			}
			stokBaru := bacaInt("  Stok Baru (0=skip)  : ")
			if stokBaru > 0 {
				daftarObat[i].Stok = stokBaru
			}
			hargaBaru := bacaFloat("  Harga Baru (0=skip) : Rp ")
			if hargaBaru > 0 {
				daftarObat[i].Harga = hargaBaru
			}
			var pil string
			fmt.Print("  Update tanggal kadaluarsa? (y/n): ")
			fmt.Scan(&pil)
			if pil == "y" || pil == "Y" {
				fmt.Println("  Tanggal Kadaluarsa Baru:")
				daftarObat[i].Kadaluarsa = bacaTanggal()
			}
			fmt.Println("  [OK] Data obat berhasil diperbarui.")
			return
		}
	}
	fmt.Printf("  [!] Obat dengan ID %d tidak ditemukan.\n", id)
}

func hapusObat() {
	tampilkanObat()
	if jumlahObat == 0 {
		return
	}
	id := bacaInt("\n  ID obat yang ingin dihapus: ")
	for i := 0; i < jumlahObat; i++ {
		if daftarObat[i].ID == id {
			var konfirmasi string
			fmt.Printf("  Yakin hapus '%s'? (y/n): ", daftarObat[i].Nama)
			fmt.Scan(&konfirmasi)
			if konfirmasi != "y" && konfirmasi != "Y" {
				fmt.Println("  [!] Penghapusan dibatalkan.")
				return
			}
			for j := i; j < jumlahObat-1; j++ {
				daftarObat[j] = daftarObat[j+1]
			}
			jumlahObat--
			fmt.Println("  [OK] Obat berhasil dihapus.")
			return
		}
	}
	fmt.Printf("  [!] Obat dengan ID %d tidak ditemukan.\n", id)
}

// ============================================================
// PENCATATAN STOK MASUK - SPESIFIKASI B
// ============================================================

func tambahStokObat() {
	tampilkanObat()
	if jumlahObat == 0 {
		return
	}
	id := bacaInt("\n  ID obat yang akan ditambah stoknya: ")
	for i := 0; i < jumlahObat; i++ {
		if daftarObat[i].ID == id {
			fmt.Printf("  Stok saat ini : %d unit\n", daftarObat[i].Stok)
			tambah := bacaInt("  Jumlah stok masuk : ")
			if tambah <= 0 {
				fmt.Println("  [!] Jumlah harus lebih dari 0.")
				return
			}
			daftarObat[i].Stok += tambah
			fmt.Printf("  [OK] Stok '%s' sekarang menjadi %d unit.\n",
				daftarObat[i].Nama, daftarObat[i].Stok)
			var pil string
			fmt.Print("  Update tanggal kadaluarsa? (y/n): ")
			fmt.Scan(&pil)
			if pil == "y" || pil == "Y" {
				fmt.Println("  Tanggal Kadaluarsa Baru:")
				daftarObat[i].Kadaluarsa = bacaTanggal()
			}
			return
		}
	}
	fmt.Printf("  [!] Obat dengan ID %d tidak ditemukan.\n", id)
}

// ============================================================
// FUNGSI MAIN - MENU UTAMA
// ============================================================

func main() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("   APOTEK-SMART - Sistem Manajemen")
		fmt.Println("========================================")
		fmt.Println("1. Kelola Kategori Gejala")
		fmt.Println("2. Kelola Data Obat")
		fmt.Println("3. Tambah Stok Obat")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("0. Keluar")
		fmt.Println("========================================")

		pilihan := bacaInt("Pilihan: ")

		switch pilihan {
		case 1:
			menuKategori()
		case 2:
			menuObat()
		case 3:
			tambahStokObat()
		case 4:
			tampilkanSemuaData()
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan APOTEK-SMART!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuKategori() {
	for {
		fmt.Println("\n--- MENU KATEGORI GEJALA ---")
		fmt.Println("1. Tambah Kategori")
		fmt.Println("2. Lihat Kategori")
		fmt.Println("3. Ubah Kategori")
		fmt.Println("4. Hapus Kategori")
		fmt.Println("0. Kembali")
		pilihan := bacaInt("Pilihan: ")

		switch pilihan {
		case 1:
			tambahKategori()
		case 2:
			tampilkanKategori()
		case 3:
			ubahKategori()
		case 4:
			hapusKategori()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuObat() {
	for {
		fmt.Println("\n--- MENU DATA OBAT ---")
		fmt.Println("1. Tambah Obat")
		fmt.Println("2. Lihat Obat")
		fmt.Println("3. Ubah Obat")
		fmt.Println("4. Hapus Obat")
		fmt.Println("0. Kembali")
		pilihan := bacaInt("Pilihan: ")

		switch pilihan {
		case 1:
			tambahObat()
		case 2:
			tampilkanObat()
		case 3:
			ubahObat()
		case 4:
			hapusObat()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tampilkanSemuaData() {
	tampilkanKategori()
	tampilkanObat()
}
