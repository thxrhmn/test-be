package main

import (
	"fmt"
	"math/rand"
)

type Karyawan struct {
	ID   int
	Nama string
	Foto string
}

// ClockIn = proses absen masuk seorang karyawan.
func ClockIn(karyawan Karyawan, uploadFoto string) bool {
	// simulasi pengecekan kemiripan dengan hasil acak (0 hingga 100)
	kemiripanFoto := rand.Intn(100)

	if kemiripanFoto >= 80 {
		fmt.Printf("Karyawan %s berhasil melakukan absen masuk!\n", karyawan.Nama)
		return true
	} else {
		fmt.Println("absensi tidak valid. Silahkan mengambil foto ulang.")
		return false
	}

}

// ClockOut = simulai proses absen pulang seorang karyawan.
func ClockOut(karyawan Karyawan, uploadFoto string) bool {
	// simulasi pengecekan kemiripan dengan hasil acak (0 hingga 100).
	kemiripanFoto := rand.Intn(100)

	if kemiripanFoto >= 80 {
		fmt.Printf("Karyawan %s berhasil melakukan absen pulang!\n", karyawan.Nama)
		return true
	} else {
		fmt.Println("Absensi tidak valid. Silahkan mengambil foto ulang.")
		return false
	}

}

func main() {
	// simulasi mengambil data karyawan dan gambar yang disimpan dari database.
	// ini adalah fake data karyawan.
	karyawans := []Karyawan{
		{ID: 1, Nama: "Jhon", Foto: "jhon.jpg"},
		{ID: 2, Nama: "Jeni", Foto: "jeni.jpg"},
	}

	// simulasi absen masuk dan absen pulang untuk karyawan pertama.
	karyawan := karyawans[0]
	uploadFoto := "upload_foto_karyawan"

	fmt.Println("Sedang melakukan absen masuk...")
	ClockIn(karyawan, uploadFoto)

	fmt.Println("Sedang melakukan absen pulang...")
	ClockOut(karyawan, uploadFoto)
}
