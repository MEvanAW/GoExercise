package main

import (
	"fmt"
	"os"
	"strconv"
)

type teman struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("There is no argument passed. Valid argument range: 0-49")
		return
	}
	arg, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if arg < 0 {
		fmt.Println("Argument may not be negative. Valid argument range: 0-49")
		return
	}
	if arg > 49 {
		fmt.Println("Argument out of range. Valid argument range: 0-49")
		return
	}
	temans := []teman{
		{
			"Muhammad Evan Anindya Wahyuaji",
			"Sugiyapranata 8, Bareng Lor, Klaten Utara, Klaten, Jawa Tengah",
			"Sedang mencari kerja",
			"Untuk memperluas ilmu Teknologi Informasi saya",
		}, {
			"Faqih Yugos",
			"Alamat 2",
			"Pekerjaan 2",
			"Alasan 2",
		}, {
			"Ilham Nopi Hendri",
			"Alamat 3",
			"Pekerjaan 3",
			"Alasan 3",
		}, {
			"Ivan Muhammad Nizar",
			"Alamat 4",
			"Pekerjaan 4",
			"Alasan 4",
		}, {
			"Laba",
			"Alamat 5",
			"Pekerjaan 5",
			"Alasan 5",
		}, {
			"Michael Brian Bagaskara",
			"Alamat 6",
			"Pekerjaan 6",
			"Alasan 6",
		}, {
			"Syahrul Apriansyah",
			"Alamat 7",
			"Pekerjaan 7",
			"Alasan 7",
		}, {
			"Wira Ardi Kesuma Adji",
			"Alamat 8",
			"Pekerjaan 8",
			"Alasan 8",
		}, {
			"Adi Ahmad Wijaya",
			"Alamat 9",
			"Pekerjaan 9",
			"Alasan 9",
		}, {
			"Ahmad Aditya Gocendra Daud",
			"Alamat 10",
			"Pekerjaan 10",
			"Alasan 10",
		},
	}
	if arg > len(temans)-1 {
		fmt.Println("Nama", arg)
		fmt.Println("Alamat", arg)
		fmt.Println("Pekerjaan", arg)
		fmt.Println("Alasan", arg)
	} else {
		temans[arg].println()
	}
}

func (t teman) println() {
	fmt.Println(t.nama)
	fmt.Println(t.alamat)
	fmt.Println(t.pekerjaan)
	fmt.Println(t.alasan)
}
