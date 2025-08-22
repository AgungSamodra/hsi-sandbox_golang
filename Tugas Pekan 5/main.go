package main

import (
	"fmt"
	"project/pegawai"
)

func main() {
	// Membuat data dummy
	karyawan := pegawai.Pegawai{
		Nama:        "Agung Samodra",
		Posisi:      "Software Engineer",
		GajiBulanan: 100000000,
	}

	// Memanggil method HitungGajiTahunan()
	fmt.Printf("Gaji Tahunan %s: Rp %.2f\n\n", karyawan.Nama, karyawan.HitungGajiTahunan())

	// Menggunakan interface
	var info pegawai.InformasiPegawai = karyawan
	info.TampilkanInformasi()
}
