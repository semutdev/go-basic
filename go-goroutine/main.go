package main

import (
	"fmt"
	"time"
)

// fungsi cek server misal
func CekServer(namaServer string, jalurLaporan chan string) {
	fmt.Printf("Mulai mengecek %s ...\n", namaServer)
	// simulasi loading 2 detik
	time.Sleep(2 * time.Second)
	pesan := fmt.Sprintf("%s beroperasi normal", namaServer) 

	// kirim pesan tersebut ke channel
	jalurLaporan <- pesan
}

func main() {
	daftarServer := []string{"Singapore", "Indonesia", "USA"}

	// ini pakai channel
	pipaLaporan := make(chan string)

	fmt.Println("=== Mulai Pengecekan====")

	for _, server := range daftarServer {
		go CekServer(server, pipaLaporan)
	}

	for i := 0; i < len(daftarServer); i++ {
		laporanMasuk := <-pipaLaporan

		fmt.Println("Laporan diterima:", laporanMasuk)
	}

	fmt.Println("==== semua server selesai dicek ====")
}