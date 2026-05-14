package main

import (
	"fmt"
)

func main() {
	// Belajar perulangan
	for i := 1; i <= 3; i++ {
		fmt.Println("putaran ke: ", i)
	}

	fmt.Println("-----------------------------------")
	// ini foreach biasanya
	daftarSiswa := []string{"jujun", "jamal", "udin"}

	for _, siswa := range daftarSiswa {
		fmt.Println(siswa)
	} 

}
