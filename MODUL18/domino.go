package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1.a) Tipe Data Kartu Domino
type Domino struct {
	KiriKartu  int  //sisi kiri kartu
	KananKartu int  //sisi kanan kartu
	TotalKartu int  //nilai total kartu
	IsBalak    bool //apakah kartu ini balak (kembar) atau bukan
}

type Dominoes struct {
	kartu     [28]Domino
	sisaKartu int
}

func inisiasiDomino() Dominoes {
	var d Dominoes
	var idx int = 0
	var i, j int

	for i = 0; i <= 6; i++ {
		for j = i; j <= 6; j++ {
			d.kartu[idx] = Domino{
				KiriKartu:  i,
				KananKartu: j,
				TotalKartu: i + j,
				IsBalak:    i == j,
			}
			idx++
		}
	}
	d.sisaKartu = 28
	return d
}

// 1.b) Prosedur kocok kartu
func kocokKartu(d *Dominoes) {
	var i, j int

	for i = d.sisaKartu - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i] //menukar posisi kartu
	}
}

// 1.c) Fungsi ambil kartu
func ambilKartu(d *Dominoes) Domino {
	if d.sisaKartu <= 0 {
		fmt.Println("Tumpukan kartu habis")
		return Domino{}
	}

	var kartuDiambil Domino
	kartuDiambil = d.kartu[d.sisaKartu-1]
	d.sisaKartu--

	return kartuDiambil
}

// 1.d) Fungsi gambar kartu
func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.KiriKartu
	} else if suit == 2 {
		return d.KananKartu
	}
	return -1
}

// 1.e) Fungsi nilai kartu
func nilaiKartu(d Domino) int {
	return d.TotalKartu
}

// 2.a) Prosedur gali kartu
func galiKartu(d *Dominoes, target Domino, kartuHasil *Domino, ditemukan *bool) {
	var kartuSekarang Domino
	*ditemukan = false

	for d.sisaKartu > 0 && !*ditemukan {
		kartuSekarang = ambilKartu(d)

		//cek suit yang cocok dengan target
		if kartuSekarang.KiriKartu == target.KiriKartu || kartuSekarang.KiriKartu == target.KananKartu ||
			kartuSekarang.KananKartu == target.KiriKartu || kartuSekarang.KananKartu == target.KananKartu {
			*kartuHasil = kartuSekarang
			*ditemukan = true
		}
	}
}

// 2.b) Fungsi sepasang kartu
func sepasangKartu(d1, d2 Domino) bool {
	var total int
	total = nilaiKartu(d1) + nilaiKartu(d2)
	return total == 12
}

// 3. Implementasi permainan Gapleh
func main() {
	rand.Seed(time.Now().UnixNano())
	var skorTotal int = 0
	var totalRonde int = 0
	var i, skorRonde, keputusan, indeksKartu, ujungKanan, ujungKiri int
	var kartuDipakai [7]bool
	var rondeSelesai, keluarGame bool
	var tumpukanKartu Dominoes
	var kartuTanganPemain [7]Domino
	var kartuMulai, kartuPilihan Domino
	var rantaiKartu []Domino
	var tingkatKemenangan float64

	for {
		fmt.Printf("Your score is %d/%d \t\t At the beginning scores=%d, round=%d\n", skorTotal, totalRonde, skorTotal, totalRonde)
		fmt.Println("Dealing... \t\t\t Shuffle the tiles for a new game")

		//1. Kocok kartu
		tumpukanKartu = inisiasiDomino()
		kocokKartu(&tumpukanKartu)

		//2. Bagi 7 kartu ke pemain
		for i = 0; i < 7; i++ {
			kartuTanganPemain[i] = ambilKartu(&tumpukanKartu)
		}

		//3.Buka satu kartu dari tumpukkan untuk memulai
		kartuMulai = ambilKartu(&tumpukanKartu)
		rantaiKartu = []Domino{kartuMulai}

		for i = 0; i < 7; i++ {
			kartuDipakai[i] = false
		}

		rondeSelesai = false
		keluarGame = false
		skorRonde = 0

		for !rondeSelesai {
			fmt.Print("Your tiles: ")
			for i = 0; i < 7; i++ {
				if !kartuDipakai[i] {
					fmt.Printf("(%d,%d) ", kartuTanganPemain[i].KiriKartu, kartuTanganPemain[i].KananKartu)
				} else {
					fmt.Print("     ")
				}
			}
			fmt.Println()

			fmt.Printf("Chain is (%d,%d) \t\t Now chain is ", rantaiKartu[0].KiriKartu, rantaiKartu[len(rantaiKartu)-1].KananKartu)
			for _, c := range rantaiKartu {
				fmt.Printf("(%d,%d)", c.KiriKartu, c.KananKartu)
			}
			fmt.Println()

			fmt.Print("Decision: ")
			fmt.Scan(&keputusan)

			if keputusan == 9 {
				keluarGame = true
				rondeSelesai = true
				break
			} else if keputusan == 0 {
				fmt.Println("No more legal move")
				rondeSelesai = true
			} else {
				indeksKartu = keputusan
				if indeksKartu < 0 {
					indeksKartu = -indeksKartu
				}
				indeksKartu -= 1

				if indeksKartu < 0 || indeksKartu >= 7 || kartuDipakai[indeksKartu] {
					fmt.Println("Invalid card choice! Try again.")
					continue
				}

				kartuPilihan = kartuTanganPemain[indeksKartu]

				if keputusan > 0 {
					ujungKanan = rantaiKartu[len(rantaiKartu)-1].KananKartu

					if kartuPilihan.KiriKartu == ujungKanan {
						rantaiKartu = append(rantaiKartu, kartuPilihan)
						kartuDipakai[indeksKartu] = true
						skorRonde++
					} else if kartuPilihan.KananKartu == ujungKanan {
						kartuPilihan.KiriKartu, kartuPilihan.KananKartu = kartuPilihan.KananKartu, kartuPilihan.KiriKartu
						rantaiKartu = append(rantaiKartu, kartuPilihan)
						kartuDipakai[indeksKartu] = true
						skorRonde++
					} else {
						fmt.Println("Card doesn't match the right end!")
					}
				} else {
					ujungKiri = rantaiKartu[0].KiriKartu

					if kartuPilihan.KananKartu == ujungKiri {
						rantaiKartu = append([]Domino{kartuPilihan}, rantaiKartu...)
						kartuDipakai[indeksKartu] = true
						skorRonde++
					} else if kartuPilihan.KiriKartu == ujungKiri {
						kartuPilihan.KiriKartu, kartuPilihan.KananKartu = kartuPilihan.KananKartu, kartuPilihan.KiriKartu
						rantaiKartu = append([]Domino{kartuPilihan}, rantaiKartu...)
						kartuDipakai[indeksKartu] = true
						skorRonde++
					} else {
						fmt.Println("Card doesn't match the left end!")
					}
				}
			}
			if skorRonde == 7 {
				rondeSelesai = true
			}
		}

		if keluarGame {
			break
		}

		skorTotal += skorRonde
		totalRonde++
		fmt.Printf("You got %d points\n\n", skorRonde)
	}
	fmt.Printf("Your last score is %d/%d\n", skorTotal, totalRonde)
	fmt.Println("Thank you for playing with us.")

	if totalRonde > 0 {
		tingkatKemenangan = (float64(skorTotal) / float64(totalRonde*7)) * 100
	} else {
		tingkatKemenangan = 0.0
	}
	fmt.Printf("Your winning rate is %.2f%%\t\t which is total scores/(#rounds*7)\n", tingkatKemenangan)
}
