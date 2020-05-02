package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(1000) + 1
	fmt.Println("=====================================================")
	fmt.Println(" Aplikasi ini digunakan sebagai bahan pertimbangan   ")
	fmt.Println("  penerimaan karyawan baru pada Pelindo III Group.   ")
	fmt.Println("        Silahkan gunakan alat bantu apapun,          ")
	fmt.Println("               tidak ada batas waktu.                ")
	fmt.Println("=====================================================")
	fmt.Println("")
	fmt.Println("Game : ")
	fmt.Println("Diantara angka 1 sampai 1000, sistem sudah menentukan")
	fmt.Println("               sebuah angka rahasia.                 ")
	fmt.Println("  Kamu memiliki 10 kali kesempatan untuk menebaknya. ")
	fmt.Println("  Sistem akan memberi petunjuk apakah angkamu terlalu")
	fmt.Println("                rendah atau tinggi.                  ")
	fmt.Println("                       SIAP ?.                       ")
	fmt.Println("=====================================================")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin) //Membaca keyboard
	tebakanSukses := false

	nilaiLogika := 100
	jumlahTebak := 10
	kisaranTerendah := 0
	kisaranTertinggi := 1000

	for jumlahTebak > 0 {

		fmt.Println("Kamu memiliki", jumlahTebak, "sisa tebakan.")
		fmt.Println("Masukkan angkamu : ")

		input, err := reader.ReadString('\n') //membaca input sampai dengan tombol enter
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		tebakan, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//Jika tebakan tidak logis dikurangi 10
		if !tebakanLogis(tebakan, kisaranTerendah, kisaranTertinggi, jumlahTebak) {
			//jika tebakan tidak logis ditambah lagi tidak masuk akal dikurangi 15
			if tebakan > kisaranTertinggi || tebakan < kisaranTerendah {
				nilaiLogika = nilaiLogika - 15
			} else {
				nilaiLogika = nilaiLogika - 10
			}
		}

		if tebakan > target {
			fmt.Printf("                                  %v terlalu tinggi.\n", tebakan)
			fmt.Println("                           -------------------------")
			if kisaranTertinggi > tebakan {
				kisaranTertinggi = tebakan
			}

		} else if tebakan < target {
			fmt.Printf("                                  %v terlalu rendah.\n", tebakan)
			fmt.Println("                           -------------------------")
			if kisaranTerendah < tebakan {
				kisaranTerendah = tebakan
			}

		} else {
			tebakanSukses = true
			fmt.Println("=====================================================")
			fmt.Printf("Selamat, Kamu berhasil menebak pada tebakan ke- %v\n", 11-jumlahTebak)
			break
		}

		jumlahTebak--
	}

	if !tebakanSukses {
		fmt.Println("=====================================================")
		fmt.Println("Maaf kamu tidak bisa menebaknya hingga akhir,\nangkanya adalah", target)
	}

	if jumlahTebak > 7 && nilaiLogika < 90 {
		fmt.Printf("Kamu sangat beruntung, namun gamenya harus di ulang karena\nperbandingan penilaiannya belum memenuhi syarat.")
	}
	fmt.Printf("Nilai logika kamu dalam menebak adalah %v dari 100\n", nilaiLogika)
	fmt.Println("=====================================================")

	reader.ReadString('\n')
	reader.ReadString('\n')
	reader.ReadString('\n')

}

func tebakanLogis(tebakan int, terendah int, tertinggi int, toleransi int) bool {

	if toleransi < 2 {
		toleransi = 2
	}

	nilaiTengah := (tertinggi + terendah) / 2
	areaBenarAtas := int(nilaiTengah) + toleransi
	areaBenarBawah := int(nilaiTengah) - toleransi

	if tebakan <= areaBenarAtas && tebakan >= areaBenarBawah {
		return true
	}
	return false
}
