package models

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
