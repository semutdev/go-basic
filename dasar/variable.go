package main

import "fmt"

func main() {
	// var name string = "jujun jamaludin"
	// var age int = 20
	// fmt.Println("halo saya " + name)
	// fmt.Printf("halo saya %s umur saya %d tahun\n", name, age)

	// siswas := []string{"joko", "budi", "ani"}
	// fmt.Println("halo saya " + siswas[0])
	// for i, siswa := range siswas {
	// 	fmt.Printf("siswa ke-%d adalah %s\n", i+1, siswa)
	// }

	// const DB_HOST = "localhost"
	// fmt.Println("Database host is " + DB_HOST)

	// pointer
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
}