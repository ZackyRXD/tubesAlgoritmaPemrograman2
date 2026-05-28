/*
 * ============================================================
 *   APOTEK-SMART
 *   Bagian: Zacky
 *   Spesifikasi C & D:
 *     - Pencarian Obat (Sequential Search & Binary Search)
 *     - Pengurutan Obat (Selection Sort & Insertion Sort)
 * ============================================================
 */

package main

import (
	"fmt"
	"strings"
)

// Membandingkan dua tanggal, true jika t1 lebih awal dari t2
func tanggalLebihAwal(t1, t2 Tanggal) bool {
	if t1.Tahun != t2.Tahun {
		return t1.Tahun < t2.Tahun
	}
	if t1.Bulan != t2.Bulan {
		return t1.Bulan < t2.Bulan
	}
	return t1.Hari < t2.Hari
}

// Mencari obat berdasarkan sebagian nama, menampilkan semua hasil cocok
func sequentialSearchNama(keyword string) {
	fmt.Println("\n  Hasil pencarian nama:", keyword)
	ketemu := false
	for i := 0; i < jumlahObat; i++ {
		if strings.Contains(strings.ToLower(daftarObat[i].Nama), strings.ToLower(keyword)) {
			o := daftarObat[i]
			fmt.Println(" ", o.ID, o.Nama, "|", o.Kategori, "|", o.Indikasi, "| Stok:", o.Stok, "| Exp:", formatTanggal(o.Kadaluarsa))
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Obat tidak ditemukan.")
	}
}

// Mencari obat berdasarkan sebagian kata indikasi/gejala, menampilkan semua hasil cocok
func sequentialSearchGejala(keyword string) {
	fmt.Println("\n  Hasil pencarian gejala:", keyword)
	ketemu := false
	for i := 0; i < jumlahObat; i++ {
		if strings.Contains(strings.ToLower(daftarObat[i].Indikasi), strings.ToLower(keyword)) {
			o := daftarObat[i]
			fmt.Println(" ", o.ID, o.Nama, "|", o.Kategori, "|", o.Indikasi, "| Stok:", o.Stok, "| Exp:", formatTanggal(o.Kadaluarsa))
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("  Obat tidak ditemukan.")
	}
}

// Mencari obat berdasarkan nama persis, data disalin dan diurutkan sementara sebelum dicari
func binarySearchNama(keyword string) {
	fmt.Println("\n  Hasil binary search nama:", keyword)
	if jumlahObat == 0 {
		fmt.Println("  Belum ada data obat.")
		return
	}

	// Salin data agar urutan asli tidak berubah
	var temp [MaxObat]Obat
	for i := 0; i < jumlahObat; i++ {
		temp[i] = daftarObat[i]
	}

	// Urutkan salinan berdasarkan nama A-Z
	for i := 1; i < jumlahObat; i++ {
		kunci := temp[i]
		j := i - 1
		for j >= 0 && strings.ToLower(temp[j].Nama) > strings.ToLower(kunci.Nama) {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = kunci
	}

	// Binary search pada salinan yang sudah terurut
	kw := strings.ToLower(keyword)
	kiri, kanan := 0, jumlahObat-1
	ketemu := false
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		nama := strings.ToLower(temp[tengah].Nama)
		if nama == kw {
			o := temp[tengah]
			fmt.Println(" ", o.ID, o.Nama, "|", o.Kategori, "|", o.Indikasi, "| Stok:", o.Stok, "| Exp:", formatTanggal(o.Kadaluarsa))
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

// Mengurutkan daftarObat dari kadaluarsa terdekat ke terjauh menggunakan Selection Sort
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

// Mengurutkan daftarObat dari kadaluarsa terdekat ke terjauh menggunakan Insertion Sort
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

// Sub-menu pencarian obat
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
			sequentialSearchNama(bacaString("Keyword nama: "))
		case 2:
			sequentialSearchGejala(bacaString("Keyword gejala: "))
		case 3:
			binarySearchNama(bacaString("Nama obat: "))
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Sub-menu pengurutan obat berdasarkan kadaluarsa
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