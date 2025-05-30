package main

import (
	"fmt"
	"math"
)

const NMAX int = 1000

type pinjaman struct {
	namaPeminjam   string
	jumlahPinjaman int
	lamaPelunasan  int
	sisaPelunasan  int
	sudahBayar     bool
	keterlambatan  int
}
type dataPinjaman [NMAX]pinjaman

func login(nama *string) {
	//tampilan awal aplikasi, user dapat masuk dan keluar aplikasi lewat tampilan ini
	//I.S Variabel nama belum terisi
	//F.S Jika user login, nama akan diisi oleh input. Jika user keluar, nama akan direset

	clearScreen()
	var pilihan int
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("|%24s%s%24s|\n", "", "Aplikasi Pinjaman BoWo Milkita", "")
	fmt.Printf("|%30s%s%30s|\n", "", "by Arendra and Rai", "")
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "Apa yang bisa kami bantu?")
	fmt.Printf("| %-76s |\n", "1. Login")
	fmt.Printf("| %-76s |\n", "2. Exit")
	fmt.Println("+------------------------------------------------------------------------------+")
	inputPilihan(&pilihan, 1, 2)
	if pilihan == 1 {
		fmt.Println("Silahkan masukkan nama anda untuk login.")
		fmt.Print("Nama saya adalah: ")
		fmt.Scan(nama)
	} else {
		*nama = ""
	}
}

func menu(nama string) {
	//menampilkan menu utama dari aplikasi, dimana user diberikan beberapa pilihan fitur
	//memanfaatkan variabel nama dan printf untuk mengeluarkan kalimat selamat datang sesuai nama user
	clearScreen()
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| Selamat datang, %-60s |\n", nama)
	fmt.Printf("| %-76s |\n", "Apa yang ingin anda lakukan hari ini?")
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "1. Ajukan pinjaman")
	fmt.Printf("| %-76s |\n", "2. Lihat pinjaman anda saat ini")
	fmt.Printf("| %-76s |\n", "3. Daftar semua pinjaman saat ini")
	fmt.Printf("| %-76s |\n", "4. Cari data seseorang")
	fmt.Printf("| %-76s |\n", "5. Bayar tagihan")
	fmt.Printf("| %-76s |\n", "6. Simulasi menuju 30 hari kedepan")
	fmt.Printf("| %-76s |\n", "7. Logout")
	fmt.Println("+------------------------------------------------------------------------------+")
}

func main() {
	var data dataPinjaman
	var nData int
	var nama, nama2 string
	var pilihan int
	var yesno string
	var i int
	login(&nama)
	for nama != "" {
		menu(nama)
		inputPilihan(&pilihan, 1, 7)
		switch pilihan {
		case 1:
			clearScreen()
			inputPinjaman(&data, &nData, nama)
			nData++
		case 2:
			clearScreen()
			cariDataSeq(&data, &nData, nama)
		case 3:
			clearScreen()
			if nData == 0 {
				fmt.Println("Belum ada data yang tercatat saat ini.")
				kembaliKeMenu()
			} else {
				fmt.Println("+------------------------------------------------------------------------------+")
				fmt.Printf("| %-76s |\n", "Berikut adalah data seluruh pinjaman saat ini.")
				for i = 0; i < nData; i++ {
					fmt.Println("+------------------------------------------------------------------------------+")
					fmt.Printf("| %s %-61d |\n", "Pinjaman nomor", i+1)
					tampilkanData(data, i)
					fmt.Println("+------------------------------------------------------------------------------+")
				}
				fmt.Println("Apakah anda ingin mengurutkan data ini?")
				inputYesNo(&yesno)
				if yesno == "Y" {
					fmt.Println("1. Urutkan berdasarkan jumlah pinjaman terbesar")
					fmt.Println("2. Urutkan berdasarkan jumlah pinjaman terkecil")
					fmt.Println("3. Urutkan berdasarkan nama ascending")
					fmt.Println("4. Urutkan berdasarkan nama descending")
					inputPilihan(&pilihan, 1, 4)
					switch pilihan {
					case 1:
						sortDataSelect(&data, nData, "desc")
					case 2:
						sortDataSelect(&data, nData, "asc")
					case 3:
						sortDataInsert(&data, nData, "asc")
					case 4:
						sortDataInsert(&data, nData, "desc")
					}
					pilihan = 0
					clearScreen()
					fmt.Println("Data telah berhasil diurutkan!")
					fmt.Println("Berikut adalah data pinjaman dengan pengurutan yang dipilih")
					for i = 0; i < nData; i++ {
						fmt.Println("+------------------------------------------------------------------------------+")
						fmt.Printf("| %s %-61d |\n", "Pinjaman nomor", i+1)
						tampilkanData(data, i)
						fmt.Println("+------------------------------------------------------------------------------+")
					}
				}
				kembaliKeMenu()
			}
		case 4:
			fmt.Print("Masukkan nama dari pengguna yang ingin anda lihat: ")
			fmt.Scan(&nama2)
			clearScreen()
			sortDataInsert(&data, nData, "min")
			index := cariDataBin(data, nData, nama2)
			if index == -1 {
				fmt.Printf("Data atas nama %s tidak ditemukan,\n", nama2)
			} else {
				fmt.Println("+------------------------------------------------------------------------------+")
				tampilkanData(data, index)
				fmt.Println("+------------------------------------------------------------------------------+")
			}
			kembaliKeMenu()
		case 5:
			clearScreen()
			bayarTagihan(&data, nData, nama)
		case 6:
			simulasi30hari(&data, nData)
			isiDummy(&data, &nData)
		case 7:
			login(&nama)
		}
	}
	fmt.Print("Terima kasih telah menggunakan aplikasi kami!")
}

func inputPinjaman(A *dataPinjaman, n *int, nama string) {
	//I.S Data pada index n berisi sembarang
	//F.S Jika user setuju dengan bunga, data pada index n akan diisi sesuai dengan input user.
	//    Jika user tidak setuju, data pada index n tidak mengalami perubahan.

	var input dataPinjaman
	var yesno string
	input[0].namaPeminjam = nama

	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "Pilih jumlah pinjaman yang anda inginkan: ")
	fmt.Println("+------------------------------------------------------------------------------+")
	listPinjaman(&input, 0)

	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "Pilih lama pelunasan yang diinginkan: ")
	fmt.Println("+------------------------------------------------------------------------------+")
	listPelunasan(&input, 0)

	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "Sesuai kebijakan dari kami, pinjaman anda akan dikenakan bunga sebesar 5%")
	fmt.Printf("| %-76s |\n", "yang dihitung dari jumlah pinjaman awal.")
	fmt.Printf("| %-76s |\n", "Keterlambatan dalam pembayaran akan dikenakan denda tambahan sebesar 10%")
	fmt.Printf("| %-76s |\n", "dari jumlah pinjaman awal, sekaligus pembayaran dari bulan yang terlewat.")
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Print("Apakah anda bersedia untuk menerima kebijakan kami?")
	inputYesNo(&yesno)
	if yesno == "Y" {
		fmt.Println("Terima kasih telah menggunakan layanan kami! Pinjamanmu sedang kami proses!")
		A[*n] = input[0]
	} else {
		fmt.Println("Ajuan pinjaman anda telah dibatalkan!")
	}
	kembaliKeMenu()
}

func listPinjaman(A *dataPinjaman, n int) {
	//I.S field jumlahPinjaman array A terisi sembarang
	//F.S field jumlahPinjaman array A terisi sesuai input user

	var pilihan int
	fmt.Println("1. Rp. 1,000,000,-")
	fmt.Println("2. Rp. 3,000,000,-")
	fmt.Println("3. Rp. 5,000,000,-")
	fmt.Println("4. Rp. 10,000,000,-")
	fmt.Println("5. Rp. 15,000,000,-")
	inputPilihan(&pilihan, 1, 5)
	switch pilihan {
	case 1:
		A[n].jumlahPinjaman = 1000000
	case 2:
		A[n].jumlahPinjaman = 3000000
	case 3:
		A[n].jumlahPinjaman = 5000000
	case 4:
		A[n].jumlahPinjaman = 10000000
	case 5:
		A[n].jumlahPinjaman = 15000000
	}
	clearScreen()
}

func listPelunasan(A *dataPinjaman, n int) {
	//I.S field lamaPelunasan dan sisaPelunasan array A terisi sembarang
	//F.S field lamaPelunasan dan sisaPelunasan array A terisi sesuai input user

	var pilihan int
	fmt.Println("1. 12 Bulan")
	fmt.Println("2. 18 Bulan")
	fmt.Println("3. 24 Bulan")
	fmt.Println("4. 30 Bulan")
	fmt.Println("5. 36 Bulan")
	inputPilihan(&pilihan, 1, 5)
	switch pilihan {
	case 1:
		A[n].lamaPelunasan = 12
	case 2:
		A[n].lamaPelunasan = 18
	case 3:
		A[n].lamaPelunasan = 24
	case 4:
		A[n].lamaPelunasan = 30
	case 5:
		A[n].lamaPelunasan = 36
	}
	A[n].sisaPelunasan = A[n].lamaPelunasan
	clearScreen()
}

func cariDataBin(A dataPinjaman, n int, nama string) int {
	//pencarian data metode binary search
	//mengembalikan indeks data dengan field namaPeminjam yang sama dengan isi variabel nama

	var index int
	var left, mid, right int
	left = 0
	right = n - 1
	index = -1
	for left <= right && index == -1 {
		mid = (left + right) / 2
		if nama < A[mid].namaPeminjam {
			right = mid - 1
		} else if nama > A[mid].namaPeminjam {
			left = mid + 1
		} else if nama == A[mid].namaPeminjam {
			index = mid
		}
	}
	return index
}

func cariDataSeq(A *dataPinjaman, n *int, nama string) {
	//pencarian data metode sequential search
	//mengeluarkan data dengan field namaPeminjam yang sama dengan isi variabel nama,
	//indeks data tersebut disimpan pada array listIndex untuk digunakan pada fitur edit dan hapus data

	var listIndex [100]int
	var i, j int
	var yesno string
	j = 1
	for i = 0; i < *n; i++ {
		if A[i].namaPeminjam == nama {
			listIndex[j-1] = i
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %s %-61d |\n", "Pinjaman nomor", j)
			tampilkanData(*A, i)
			fmt.Println("+------------------------------------------------------------------------------+")
			j++
		}
	}
	if j == 1 {
		fmt.Println("Anda belum memiliki pinjaman")
		kembaliKeMenu()
	} else {
		fmt.Println("Apakah ada data yang ingin diubah?")
		inputYesNo(&yesno)
		if yesno == "Y" {
			j--
			editHapusData(A, n, &listIndex, &j, nama)
		}
	}
}

func tampilkanData(A dataPinjaman, i int) {
	//menampilkan dataPinjaman pada indeks ke i

	fmt.Printf("| %-23s: %-51s |\n", "Nama peminjam", A[i].namaPeminjam)
	fmt.Printf("| %-23s: %-51s |\n", "Jumlah pinjaman", formatUang(A[i].jumlahPinjaman))
	fmt.Printf("| %-23s: %d %-48s |\n", "Lama pelunasan", A[i].lamaPelunasan, "bulan")
	fmt.Printf("| %-23s: %-51s |\n", "Status tagihan bulanan", statusPembayaran(A[i].sudahBayar))
}

func formatUang(x int) string {
	//mengubah format uang x dalam integer, dikembalikan sebagai string format rupiah

	var strAngka string
	var i int = 1
	for x > 0 {
		strAngka = string(x%10+48) + strAngka
		if i%3 == 0 && x/10 > 0 {
			strAngka = "," + strAngka
		}
		i++
		x = x / 10
	}
	strAngka = "Rp. " + strAngka + ",-"
	return strAngka
}

func statusPembayaran(x bool) string {
	//mengembalikan string "Sudah dibayar" atau "Belum dibayar" sesuai nilai boolean x
	if x {
		return "Sudah dibayar"
	} else {
		return "Belum dibayar"
	}
}

func editHapusData(A *dataPinjaman, n *int, listIndex *[100]int, nIndex *int, nama string) {
	//user dapat memilih untuk mengedit atau menghapus data sesuai nama user
	//I.S Ditampilkan data pinjaman dari user yang terisi sembarang
	//F.S Jika user menyetujui, data yang dipilih akan berubah sesuai dengan input user

	var pilihan int
	var i int
	var nomor int
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Printf("| %-76s |\n", "Pilih operasi yang diinginkan: ")
	fmt.Println("+------------------------------------------------------------------------------+")
	fmt.Println("1. Edit data")
	fmt.Println("2. Hapus data")
	fmt.Println("3. Kembali")
	inputPilihan(&pilihan, 1, 3)
	clearScreen()
	switch pilihan {
	case 1:
		for i = 0; i < *nIndex; i++ {
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %s %-61d |\n", "Pinjaman nomor", i+1)
			tampilkanData(*A, i)
			fmt.Println("+------------------------------------------------------------------------------+")
		}
		fmt.Println("Data mana yang ingin anda ubah?")
		inputPilihan(&nomor, 1, i)
		clearScreen()
		fmt.Print("Silahkan lakukan input ulang data yang baru.\n")
		inputPinjaman(A, &listIndex[nomor-1], nama)
		fmt.Println("Data anda telah berhasil diubah!")
	case 2:
		for i = 0; i < *nIndex; i++ {
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %s %-61d |\n", "Pinjaman nomor", i+1)
			tampilkanData(*A, i)
			fmt.Println("+------------------------------------------------------------------------------+")
		}
		fmt.Println("Data mana yang ingin anda hapus?")
		inputPilihan(&nomor, 1, i)
		hapusData(A, n, listIndex[nomor-1])
		fmt.Println("Data telah berhasil dihapus!")
	}
}

func hapusData(A *dataPinjaman, n *int, idxHapus int) {
	//I.S data user terisi sembarang
	//F.S data pada indeks idxHapus ditimpa dengan data pada idxHapus+1
	//    data pada indeks idxHapus+1 ditimpa dengan data pada idxHapus+2, dst

	var i int
	for i = idxHapus; i < *n; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func sortDataInsert(A *dataPinjaman, n int, versi string) {
	//sorting data metode insertion sort
	//I.S data pinjaman terisi sembarang
	//F.S data pinjaman terurut secara ascending atau descending berdasarkan versi

	var pass, i int
	var temp pinjaman
	pass = 1
	if versi == "asc" {
		for pass < n {
			temp = A[pass]
			i = pass - 1
			for i >= 0 && A[i].namaPeminjam < temp.namaPeminjam {
				A[i+1] = A[i]
				i--
			}
			A[i+1] = temp
			pass++
		}
	} else {
		for pass < n {
			temp = A[pass]
			i = pass - 1
			for i >= 0 && A[i].namaPeminjam > temp.namaPeminjam {
				A[i+1] = A[i]
				i--
			}
			A[i+1] = temp
			pass++
		}
	}
}

func sortDataSelect(A *dataPinjaman, n int, versi string) {
	//sorting data metode selection sort
	//I.S data pinjaman terisi sembarang
	//F.S data pinjaman terurut secara ascending atau descending berdasarkan versi

	var pass, idx int
	if versi == "desc" {
		for pass = 1; pass < n; pass++ {
			idx = maxPos(*A, n, pass)
			tukar(A, idx, pass-1)
		}
	} else {
		for pass = 1; pass < n; pass++ {
			idx = minPos(*A, n, pass)
			tukar(A, idx, pass-1)
		}
	}
}

func maxPos(A dataPinjaman, n, pass int) int {
	//mengembalikan nilai jumlahPinjaman yang tertinggi

	var idx, i int
	idx = pass - 1
	for i = pass; i < n; i++ {
		if A[i].jumlahPinjaman > A[idx].jumlahPinjaman {
			idx = i
		}
	}
	return idx
}

func minPos(A dataPinjaman, n, pass int) int {
	//mengembalikan nilai jumlahPinjaman yang terendah

	var idx, i int
	idx = pass - 1
	for i = pass; i < n; i++ {
		if A[i].jumlahPinjaman < A[idx].jumlahPinjaman {
			idx = i
		}
	}
	return idx
}

func tukar(A *dataPinjaman, i, j int) {
	//I.S data pinjaman terisi sembarang
	//F.S isi data pinjaman pada indeks i dan j telah ditukar

	var temp pinjaman
	temp = A[i]
	A[i] = A[j]
	A[j] = temp
}

func bayarTagihan(A *dataPinjaman, n int, nama string) {
	//I.S menampilkan data pinjaman dengan nama user yang tagihannya belum terbayar
	//F.S jika user menyetujui, user dapat membayar tagihan, mengubah field sudahDibayar menjadi true

	var listIndex [100]int
	var i, j int
	var pilihan int
	var yesno string
	var tagihan float64
	for i = 0; i < n; i++ {
		if !A[i].sudahBayar && A[i].namaPeminjam == nama {
			listIndex[j] = i
			j++
		}
	}
	if j == 0 {
		fmt.Println("Anda tidak memiliki tagihan yang belum terbayar.")
		kembaliKeMenu()
	} else {
		fmt.Println("+------------------------------------------------------------------------------+")
		fmt.Printf("| %-76s |\n", "Berikut adalah data tagihan yang belum anda bayar.")
		fmt.Println("+------------------------------------------------------------------------------+")
		for i = 0; i < j; i++ {
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %s %-62d |\n", "Tagihan nomor", i+1)
			tampilkanData(*A, listIndex[i])
			fmt.Println("+------------------------------------------------------------------------------+")
		}
		fmt.Println("+------------------------------------------------------------------------------+")
		fmt.Printf("| %-76s |\n", "Apakah anda ingin membayar tagihan tersebut?")
		fmt.Println("+------------------------------------------------------------------------------+")
		inputYesNo(&yesno)
		if yesno == "Y" {
			clearScreen()
			for i = 0; i < j; i++ {
				fmt.Println("+------------------------------------------------------------------------------+")
				fmt.Printf("| %s %-61d |\n", "Tagihan nomor", i+1)
				tampilkanData(*A, listIndex[i])
				fmt.Println("+------------------------------------------------------------------------------+")
			}
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %-76s |\n", "Tagihan mana yang ingin anda bayar?")
			fmt.Println("+------------------------------------------------------------------------------+")
			inputPilihan(&pilihan, 1, j)
			if A[pilihan-1].keterlambatan == 0 {
				tagihan = float64(A[pilihan-1].jumlahPinjaman) * 1.05 / float64(A[pilihan-1].lamaPelunasan)
				tagihan = math.Round(tagihan)
			} else {
				tagihan = float64(A[pilihan-1].jumlahPinjaman) * 1.05 / float64(A[pilihan-1].lamaPelunasan)
				tagihan = float64(A[pilihan-1].keterlambatan+1) * tagihan
				tagihan = tagihan + 0.1*float64(A[pilihan-1].jumlahPinjaman)
				tagihan = math.Round(tagihan)
			}
			clearScreen()
			fmt.Println("+------------------------------------------------------------------------------+")
			fmt.Printf("| %-76s |\n", "Berikut adalah nominal dari tagihan anda:")
			fmt.Printf("| %-76s |\n", formatUang(int(tagihan)))
			fmt.Printf("| %-76s |\n", "Apakah anda ingin membayarnya?")
			fmt.Println("+------------------------------------------------------------------------------+")
			inputYesNo(&yesno)
			if yesno == "Y" {
				A[pilihan-1].sudahBayar = true
				fmt.Println("Anda telah berhasil membayar tagihan!")
				kembaliKeMenu()
			}
		}
	}
}

func clearScreen() {
	//membersihkan layar pada terminal, didapat dari:
	//https://stackoverflow.com/a/22892171

	fmt.Print("\033[H\033[2J")
}

func kembaliKeMenu() {
	//program akan menunggu input dari user sebelum berjalan kembali
	//sehingga user dapat membaca informasi baru sebelum terminal dibersihkan kembali

	var esc int
	fmt.Println("Ketik apapun untuk kembali ke menu")
	fmt.Scan(&esc)
}

func inputPilihan(pilihan *int, a, b int) {
	//I.S variabel pilihan berisi sembarang
	//F.S variabel pilihan terisi nilai sesuai input user,
	//    dengan batas atas a dan batas bawah b untuk mencegah error dari input yang tidak valid

	fmt.Print("Ketik angkanya saja: ")
	fmt.Scan(pilihan)
	for *pilihan < a || *pilihan > b {
		fmt.Println("Input tidak sesuai")
		fmt.Print("Masukkan input: ")
		fmt.Scan(pilihan)
	}
}

func inputYesNo(yesno *string) {
	//I.S variabel yesno berisi sembarang
	//F.S variabel yesno berisi "Y" atau "N" sesuai input user,
	//    jika input tidak sesuai, akan diminta lagi sampai sesuai

	fmt.Print("(Y/N): ")
	fmt.Scan(yesno)
	for *yesno != "N" && *yesno != "Y" {
		fmt.Println("Input tidak sesuai")
		fmt.Print("(Y/N): ")
		fmt.Scan(yesno)
	}
}

func simulasi30hari(A *dataPinjaman, n int) {
	//mensimulasikan jika aplikasi telah berjalan selama 1 bulan
	//I.S data pinjaman terisi sembarang
	//F.S field sisaPelunasan data pinjaman dikurangi 1, field sudahBayar diisi dengan false
	//    data pinjaman yang tagihan sebelumnya belum dibayar, akan ditingkatkan nilai keterlambatannya

	var i int
	for i = 0; i < n; i++ {
		A[i].sisaPelunasan--
		if !A[i].sudahBayar {
			A[i].keterlambatan++
		}
		A[i].sudahBayar = false
	}
}

func isiDummy(A *dataPinjaman, n *int) {
	//I.S data pinjaman terisi sembarang
	//F.S data pinjaman menerima data dummy mulai dari indeks n

	A[*n].namaPeminjam = "Tornade"
	A[*n].jumlahPinjaman = 3000000
	A[*n].lamaPelunasan = 12
	A[*n].sisaPelunasan = 12
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Anima"
	A[*n].jumlahPinjaman = 10000000
	A[*n].lamaPelunasan = 24
	A[*n].sisaPelunasan = 24
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Shinako"
	A[*n].jumlahPinjaman = 1000000
	A[*n].lamaPelunasan = 36
	A[*n].sisaPelunasan = 36
	A[*n].sudahBayar = false
	A[*n].keterlambatan = 1
	*n++

	A[*n].namaPeminjam = "Slimy"
	A[*n].jumlahPinjaman = 15000000
	A[*n].lamaPelunasan = 30
	A[*n].sisaPelunasan = 30
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Lusmif"
	A[*n].jumlahPinjaman = 1000000
	A[*n].lamaPelunasan = 12
	A[*n].sisaPelunasan = 12
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Shrub"
	A[*n].jumlahPinjaman = 5000000
	A[*n].lamaPelunasan = 18
	A[*n].sisaPelunasan = 18
	A[*n].sudahBayar = false
	A[*n].keterlambatan = 1
	*n++

	A[*n].namaPeminjam = "Mewo"
	A[*n].jumlahPinjaman = 3000000
	A[*n].lamaPelunasan = 24
	A[*n].sisaPelunasan = 24
	A[*n].sudahBayar = false
	A[*n].keterlambatan = 1
	*n++

	A[*n].namaPeminjam = "Neon"
	A[*n].jumlahPinjaman = 10000000
	A[*n].lamaPelunasan = 18
	A[*n].sisaPelunasan = 18
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Nick"
	A[*n].jumlahPinjaman = 1000000
	A[*n].lamaPelunasan = 12
	A[*n].sisaPelunasan = 12
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Domo"
	A[*n].jumlahPinjaman = 15000000
	A[*n].lamaPelunasan = 30
	A[*n].sisaPelunasan = 30
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Orca"
	A[*n].jumlahPinjaman = 10000000
	A[*n].lamaPelunasan = 36
	A[*n].sisaPelunasan = 36
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Zalter"
	A[*n].jumlahPinjaman = 3000000
	A[*n].lamaPelunasan = 24
	A[*n].sisaPelunasan = 24
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++

	A[*n].namaPeminjam = "Vernep"
	A[*n].jumlahPinjaman = 10000000
	A[*n].lamaPelunasan = 18
	A[*n].sisaPelunasan = 18
	A[*n].sudahBayar = false
	A[*n].keterlambatan = 1
	*n++

	A[*n].namaPeminjam = "Ren"
	A[*n].jumlahPinjaman = 5000000
	A[*n].lamaPelunasan = 12
	A[*n].sisaPelunasan = 12
	A[*n].sudahBayar = true
	A[*n].keterlambatan = 0
	*n++
}
