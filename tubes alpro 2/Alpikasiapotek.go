/*
 * ============================================================
 *   APOTEK-SMART
 *   Aplikasi Manajemen Stok dan Inventaris Apotek
 *   Tugas Besar Algoritma Pemrograman 2
 * ============================================================
 *   Anggota Kelompok:
 *     - Mirza  : Spesifikasi A & B (CRUD Data, Struktur Utama)
 *     - Zacky  : Spesifikasi C & D (Pencarian & Pengurutan)
 *     - Devia  : Spesifikasi E & Integrasi (Statistik & Main)
 * ------------------------------------------------------------
 *   Kode ekstra: eel | akun: @jebb_24
 * ============================================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

// ============================================================
// BAGIAN MIRZA - KONSTANTA, STRUCT, VARIABEL GLOBAL
// ============================================================

const (
	MaxObat     = 100
	MaxKategori = 20
	StokMinimum = 10
)

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

var daftarObat [MaxObat]Obat
var daftarKategori [MaxKategori]KategoriGejala
var jumlahObat int = 0
var jumlahKategori int = 0
var nextIDObat int = 1
var nextIDKategori int = 1

var reader = bufio.NewReader(os.Stdin)

// ============================================================
// BAGIAN MIRZA - FUNGSI UTILITAS INPUT
// ============================================================

func bacaString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	// hapus \r\n atau \n di akhir
	result := ""
	for i := 0; i < len(input); i++ {
		if input[i] != '\r' && input[i] != '\n' {
			result += string(input[i])
		}
	}
	return result
}

func bacaInt(prompt string) int {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		reader.ReadString('\n') // buang sisa newline
		if err == nil {
			return input
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka!")
	}
}

func bacaFloat(prompt string) float64 {
	var input float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		reader.ReadString('\n') // buang sisa newline
		if err == nil {
			return input
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka!")
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
// BAGIAN MIRZA - CRUD KATEGORI GEJALA (SPESIFIKASI A)
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
	k.Nama = bacaString("  Nama Kategori : ")
	k.Deskripsi = bacaString("  Deskripsi     : ")
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
			tmp := bacaString("  Nama baru      : ")
			if tmp != "" {
				daftarKategori[i].Nama = tmp
			}
			tmp = bacaString("  Deskripsi baru : ")
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
			konfirmasi := bacaString(fmt.Sprintf("  Yakin hapus '%s'? (y/n): ", daftarKategori[i].Nama))
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
// BAGIAN MIRZA - CRUD DATA OBAT (SPESIFIKASI A)
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
	o.Nama = bacaString("  Nama Obat    : ")
	o.Kategori = bacaString("  Kategori     : ")
	o.Indikasi = bacaString("  Indikasi     : ")
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
			fmt.Println("  (Kosongkan untuk melewati)")
			tmp := bacaString("  Nama Obat    : ")
			if tmp != "" {
				daftarObat[i].Nama = tmp
			}
			tmp = bacaString("  Kategori     : ")
			if tmp != "" {
				daftarObat[i].Kategori = tmp
			}
			tmp = bacaString("  Indikasi     : ")
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
			pil := bacaString("  Update tanggal kadaluarsa? (y/n): ")
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
			konfirmasi := bacaString(fmt.Sprintf("  Yakin hapus '%s'? (y/n): ", daftarObat[i].Nama))
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
// BAGIAN MIRZA - PENCATATAN STOK MASUK (SPESIFIKASI B)
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
			pil := bacaString("  Update tanggal kadaluarsa? (y/n): ")
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
// BAGIAN ZACKY - FUNGSI BANTU
// ============================================================

func toLowerChar(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func toLower(s string) string {
	hasil := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		hasil[i] = toLowerChar(s[i])
	}
	return string(hasil)
}

func mengandung(s, sub string) bool {
	ls, lsub := len(s), len(sub)
	if lsub == 0 {
		return true
	}
	for i := 0; i <= ls-lsub; i++ {
		cocok := true
		for j := 0; j < lsub; j++ {
			if s[i+j] != sub[j] {
				cocok = false
				break
			}
		}
		if cocok {
			return true
		}
	}
	return false
}

func tanggalLebihAwal(t1, t2 Tanggal) bool {
	if t1.Tahun != t2.Tahun {
		return t1.Tahun < t2.Tahun
	}
	if t1.Bulan != t2.Bulan {
		return t1.Bulan < t2.Bulan
	}
	return t1.Hari < t2.Hari
}

// ============================================================
// BAGIAN ZACKY - PENCARIAN OBAT (SPESIFIKASI C)
// ============================================================

func sequentialSearchNama(keyword string) {
	fmt.Println("\n  Hasil pencarian nama:", keyword)
	ketemu := false
	for i := 0; i < jumlahObat; i++ {
		if mengandung(toLower(daftarObat[i].Nama), toLower(keyword)) {
			o := daftarObat[i]
			fmt.Printf("  %d | %s | %s | %s | Stok: %d | Exp: %s\n",
				o.ID, o.Nama, o.Kategori, o.Indikasi, o.Stok, formatTanggal(o.Kadaluarsa))
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Obat tidak ditemukan.")
	}
}

func sequentialSearchGejala(keyword string) {
	fmt.Println("\n  Hasil pencarian gejala:", keyword)
	ketemu := false
	for i := 0; i < jumlahObat; i++ {
		if mengandung(toLower(daftarObat[i].Indikasi), toLower(keyword)) {
			o := daftarObat[i]
			fmt.Printf("  %d | %s | %s | %s | Stok: %d | Exp: %s\n",
				o.ID, o.Nama, o.Kategori, o.Indikasi, o.Stok, formatTanggal(o.Kadaluarsa))
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Obat tidak ditemukan.")
	}
}

func binarySearchNama(keyword string) {
	fmt.Println("\n  Hasil binary search nama:", keyword)
	if jumlahObat == 0 {
		fmt.Println("  Belum ada data obat.")
		return
	}
	var temp [MaxObat]Obat
	for i := 0; i < jumlahObat; i++ {
		temp[i] = daftarObat[i]
	}
	for i := 1; i < jumlahObat; i++ {
		kunci := temp[i]
		j := i - 1
		for j >= 0 && toLower(temp[j].Nama) > toLower(kunci.Nama) {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = kunci
	}
	kw := toLower(keyword)
	kiri := 0
	kanan := jumlahObat - 1
	ketemu := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		nama := toLower(temp[tengah].Nama)
		if nama == kw {
			o := temp[tengah]
			fmt.Printf("  %d | %s | %s | %s | Stok: %d | Exp: %s\n",
				o.ID, o.Nama, o.Kategori, o.Indikasi, o.Stok, formatTanggal(o.Kadaluarsa))
			ketemu = true
			break
		} else if nama < kw {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	if !ketemu {
		fmt.Println("  Obat tidak ditemukan.")
	}
}

// ============================================================
// BAGIAN ZACKY - PENGURUTAN OBAT (SPESIFIKASI D)
// ============================================================

func selectionSortKadaluarsa() {
	for i := 0; i < jumlahObat-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahObat; j++ {
			if tanggalLebihAwal(daftarObat[j].Kadaluarsa, daftarObat[idxMin].Kadaluarsa) {
				idxMin = j
			}
		}
		daftarObat[i], daftarObat[idxMin] = daftarObat[idxMin], daftarObat[i]
	}
	fmt.Println("  Data diurutkan dengan Selection Sort.")
	tampilkanObat()
}

func insertionSortKadaluarsa() {
	for i := 1; i < jumlahObat; i++ {
		kunci := daftarObat[i]
		j := i - 1
		for j >= 0 && tanggalLebihAwal(kunci.Kadaluarsa, daftarObat[j].Kadaluarsa) {
			daftarObat[j+1] = daftarObat[j]
			j--
		}
		daftarObat[j+1] = kunci
	}
	fmt.Println("  Data diurutkan dengan Insertion Sort.")
	tampilkanObat()
}

func menuCariObat() {
	for {
		fmt.Println("\n--- MENU PENCARIAN OBAT ---")
		fmt.Println("1. Sequential Search - Nama")
		fmt.Println("2. Sequential Search - Gejala")
		fmt.Println("3. Binary Search     - Nama (exact)")
		fmt.Println("0. Kembali")
		pilihan := bacaInt("Pilihan: ")
		switch pilihan {
		case 1:
			sequentialSearchNama(bacaString("Keyword nama  : "))
		case 2:
			sequentialSearchGejala(bacaString("Keyword gejala: "))
		case 3:
			binarySearchNama(bacaString("Nama obat     : "))
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUrutkanObat() {
	for {
		fmt.Println("\n--- MENU PENGURUTAN OBAT ---")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Println("0. Kembali")
		pilihan := bacaInt("Pilihan: ")
		switch pilihan {
		case 1:
			selectionSortKadaluarsa()
		case 2:
			insertionSortKadaluarsa()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ============================================================
// BAGIAN DEVIA - STATISTIK OBAT (SPESIFIKASI E)
// ============================================================

func statistikObat() {
	fmt.Println("\n+++ APOTEK-SMART +++")
	fmt.Println("\nObat Hampir Habis (stok <= 5):")
	adaHampirHabis := false
	for i := 0; i < jumlahObat; i++ {
		if daftarObat[i].Stok <= 5 {
			fmt.Println("-", daftarObat[i].Nama, "| Stok:", daftarObat[i].Stok)
			adaHampirHabis = true
		}
	}
	if !adaHampirHabis {
		fmt.Println("  Semua stok aman.")
	}
	fmt.Println("\nObat Segera Kadaluarsa:")
	if jumlahObat == 0 {
		fmt.Println("  Belum ada data obat.")
		return
	}
	for i := 0; i < jumlahObat; i++ {
		fmt.Println("-", daftarObat[i].Nama, "| Exp:", formatTanggal(daftarObat[i].Kadaluarsa))
	}
}

// ============================================================
// BAGIAN DEVIA - MENU & MAIN PROGRAM
// ============================================================

func tampilkanSemuaData() {
	tampilkanKategori()
	tampilkanObat()
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

func main() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("+++ APOTEK-SMART +++")
		fmt.Println("Sistem Manajemen Stok & Inventaris Apotek")
		fmt.Println("[eel | @jebb_24]")
		fmt.Println("========================================")
		fmt.Println("1. Kelola Data Obat")
		fmt.Println("2. Kelola Kategori Gejala")
		fmt.Println("3. Tambah Stok Obat")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("5. Cari Obat")
		fmt.Println("6. Urutkan Data Obat")
		fmt.Println("7. Statistik Obat")
		fmt.Println("0. Keluar")
		fmt.Println("========================================")

		pilihan := bacaInt("Pilih menu: ")

		switch pilihan {
		case 1:
			menuObat()
		case 2:
			menuKategori()
		case 3:
			tambahStokObat()
		case 4:
			tampilkanSemuaData()
		case 5:
			menuCariObat()
		case 6:
			menuUrutkanObat()
		case 7:
			statistikObat()
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan APOTEK-SMART!")
			fmt.Println("[eel | @jebb_24]")
			return
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
