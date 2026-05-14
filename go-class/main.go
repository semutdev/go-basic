package main

import (
	"fmt"
)

type Siswa struct {
	Nama string
	Alamat string
	Umur int
}

func (s Siswa) Perkenalan() {
	fmt.Printf("halo, nama saya %s dari %s.\n", s.Nama, s.Alamat)
}

func (s *Siswa) UlangTahun() {
	s.Umur++
}

func main() {

	siswa1 := Siswa{
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
