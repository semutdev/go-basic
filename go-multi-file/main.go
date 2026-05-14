package main

import (
	"fmt"
	"go-multi-file/models"
)

func main() {

	siswa1 := models.Siswa{
		Nama: "jujun jamaludin",
		Alamat: "Dayeuhluhur",
		Umur: 26,
	}

	siswa1.Perkenalan()

	fmt.Println("Umur awal:", siswa1.Umur)
	siswa1.UlangTahun()
	fmt.Println("Umur setelah ulang tahun:", siswa1.Umur)

	// fmt.Println(siswa1.Nama)
	// fmt.Println(siswa1.Alamat)
	// fmt.Println(siswa1.Umur)

}
