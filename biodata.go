package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	data := []Teman{

		{Nama: "Nabila Cahyani", Alamat: "Jakarta", Pekerjaan: "Freelancer", Alasan: "untuk menambah pengetahuan tentang backend"},
		{Nama: "Latifah", Alamat: "Jogja", Pekerjaan: "Freelancer", Alasan: "menambah ilmu dan untuk mencari kerja"},
		{Nama: "Vita Tri Utami", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "menambah pengetahuan dalam service suatu website"},
		{Nama: "Kulsum Khoiriah", Alamat: "Bogor", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Komang Andira Trisna P.", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambah pengetahuan dan pemahaman bahasa Go"},
		{Nama: "Amelia Sari", Alamat: "Pamulang", Pekerjaan: "Freelance", Alasan: "untuk menambah pengetahuan backend"},
		{Nama: "Feti Lutviyani ", Alamat: "Bandung", Pekerjaan: "Freelance", Alasan: "Menambahskill di bidang IT"},
		{Nama: "Jesica", Alamat: "Padang", Pekerjaan: "Freelance", Alasan: "ingin mengenal pemrograman golang "},
		{Nama: "Vini Jumatul Fitri", Alamat: "Medan", Pekerjaan: "Freelance", Alasan: "menambah bahasa pemrograman"},
		{Nama: "Natasya Sitorus", Alamat: "Semarang", Pekerjaan: "Freelance", Alasan: "Untuk menambah pengetahuan lebih mengenai go"},
	}

	input := os.Args
	inputAngka, err := strconv.Atoi(input[1])
	if err != nil {
		panic("bukan angka kk")
	}

	fmt.Println("Nama: ", data[inputAngka].Nama)
	fmt.Println("Alamat: ", data[inputAngka].Alamat)
	fmt.Println("Pekerjaan: ", data[inputAngka].Pekerjaan)
	fmt.Println("Alasan Mengikuti DTS: ", data[inputAngka].Alasan)

}
