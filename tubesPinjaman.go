package main
import "fmt"

const NMAX int = 1000
type pinjaman struct {
	namaPeminjam string
	jumlahPinjaman int
	lamaPelunasan int
	bunga float64
}
type dataPinjaman [NMAX] pinjaman

func main() {
	var data dataPinjaman
	var nData int
	var nama string
	var i int
	var pilihan int
	for pilihan != 4 {
		menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1: 
			inputPinjaman(&data, &nData)
			nData++
		case 2: 
			fmt.Print("Masukkan nama anda: ")
			fmt.Scan(&nama)
			cariData(&data, &nData, &nama)
		case 3:
			fmt.Println("Berikut adalah seluruh pinjaman yang ada saat ini, sesuai waktu pengajuan")
			for i=0; i<nData; i++ {
				fmt.Printf("Data peminjam nomor %d\n", i+1)
				tampilkanData(&data, i)
			}
			fmt.Println("Apakah anda ingin mengurutkan data ini?")
			fmt.Print("(Y/N)? ")
			fmt.Scan(&nama)
			for nama != "N" {
				if nama == "Y" {
				fmt.Println("1. Urutkan berdasarkan jumlah pinjaman terbesar")
				fmt.Println("2. Urutkan berdasarkan jumlah pinjaman terkecil")
				fmt.Println("3. Urutkan berdasarkan lama pelunasan terbesar")
				fmt.Println("4. Urutkan berdasarkan lama pelunasan terkecil")
				fmt.Scan(&pilihan)
				switch pilihan {
				case 1: //selection sort max
					sortDataJumlah(&data, nData, "max")
				case 2: //selection sort min
					sortDataJumlah(&data, nData, "min")
				case 3: //insertion sort max
				case 4: //insertion sort min
				}
				pilihan = 0
				nama = "N"
			} else {
				fmt.Println("Input tidak sesuai")
				fmt.Scan(&nama)
			}
			fmt.Println("Data telah berhasil diurutkan!")
			fmt.Println("Berikut adalah data pinjaman dengan pengurutan yang dipilih")
			for i=0; i<nData; i++ {
				fmt.Printf("Data peminjam nomor %d\n", i+1)
				tampilkanData(&data, i)
			}
	}
		case 4:
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
		}
	}
}

func inputPinjaman(A *dataPinjaman, n *int) {
	fmt.Print("Input nama: ")
	fmt.Scan(&A[*n].namaPeminjam)
	fmt.Println("Tentukan jumlah pinjaman yang diinginkan: ")
	listPinjaman(A, n)
	fmt.Println("Tentukan lama pelunasan yang diinginkan: ")
	listPelunasan(A, n)
	A[*n].bunga = 5
}

func cariData(A *dataPinjaman, n *int, nama *string) {
	var listIndex [100]int
	var i, j int
	var yesno string
	j = 1
	for i=0; i<*n; i++ {
		if A[i].namaPeminjam == *nama {
			listIndex[j-1] = i
			fmt.Printf("Pinjaman nomor %d\n", j)
			tampilkanData(A, i)
			j++
		}
	}
	fmt.Println("Apakah ada data yang ingin diubah?")
	fmt.Print("(Y/N)? ")
	fmt.Scan(&yesno)
	for yesno != "N" {
		if yesno == "Y" {
			j--
			editHapusData(A, n, &listIndex, &j)
			yesno = "N"
		} else {
			fmt.Println("Input tidak sesuai")
			fmt.Scan(&yesno)
		}
	}
}

func sortDataJumlah(A *dataPinjaman, n int, versi string) {
	var pass, idx int
	if versi == "max" {
		for pass=1; pass<n; pass++ {
			idx = maxPos(*A,n,pass)
			tukar(A, idx, pass-1)
		}
	} else {
		for pass=1; pass<n; pass++ {
			idx = minPos(*A,n,pass)
			tukar(A, idx, pass-1)
		}
	}
}

func maxPos(A dataPinjaman, n, pass int) int {
	var idx, i int
	idx = pass-1
	for i=pass; i<n; i++ {
		if A[i].jumlahPinjaman > A[idx].jumlahPinjaman {
			idx = i
		}
	}
	return idx
}

func minPos(A dataPinjaman, n, pass int) int {
	var idx, i int
	idx = pass-1
	for i=pass; i<n; i++ {
		if A[i].jumlahPinjaman < A[idx].jumlahPinjaman {
			idx = i
		}
	}
	return idx
}

func tukar(A *dataPinjaman, i,j int) {
	var temp pinjaman
	temp = A[i]
	A[i] = A[j]
	A[j] = temp
}

func editHapusData(A *dataPinjaman, n *int, listIndex *[100]int, nIndex *int) {
	var pilihan int
	var i int
	var nomor int
	fmt.Println("Pilih operasi yang diinginkan")
	fmt.Println("1. Edit data")
	fmt.Println("2. Hapus data")
	fmt.Println("3. Kembali")
	fmt.Scan(&pilihan)
	for pilihan != 3 {
		switch pilihan {
		case 1:
			for i=0; i<*nIndex; i++ {
				fmt.Printf("Pinjaman nomor %d\n", i+1)
				tampilkanData(A, listIndex[i])
			}
			fmt.Print("Masukkan nomor data yang ingin di edit: ")
			fmt.Scan(&nomor)
			fmt.Print("Silahkan lakukan input ulang data yang baru\n")
			inputPinjaman(A, &listIndex[nomor-1])
			fmt.Println("Data anda sudah berhasil diubah")
		case 2:
			for i=0; i<*nIndex; i++ {
				fmt.Printf("Pinjaman nomor %d\n", i+1)
				tampilkanData(A, listIndex[i])
			}
			fmt.Print("Masukkan nomor data yang ingin di hapus: ")
			fmt.Scan(&nomor)
			hapusData(A, n, listIndex[nomor-1])
			fmt.Println("Data telah dihapus")
		}
		pilihan = 3
	}
}

func hapusData(A *dataPinjaman, n *int, idxHapus int) {
	var i int
	for i=idxHapus; i<*n; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func tampilkanData(A *dataPinjaman, i int) {
	fmt.Printf("Nama: %s\n", A[i].namaPeminjam)
	fmt.Printf("Jumlah pinjaman: %d\n", A[i].jumlahPinjaman)
	fmt.Printf("Lama pelunasan: %d bulan\n\n", A[i].lamaPelunasan)
}

func listPinjaman(A *dataPinjaman, n *int) {
	var pilihan int
	fmt.Println("1. Rp. 1,000,000,-")
	fmt.Println("2. Rp. 3,000,000,-")
	fmt.Println("3. Rp. 5,000,000,-")
	fmt.Println("4. Rp. 10,000,000,-")
	fmt.Println("5. Rp. 15,000,000,-")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		A[*n].jumlahPinjaman = 1000000
	case 2:
		A[*n].jumlahPinjaman = 3000000
	case 3:
		A[*n].jumlahPinjaman = 5000000
	case 4: 
		A[*n].jumlahPinjaman = 10000000
	case 5:
		A[*n].jumlahPinjaman = 15000000
	}
}

func listPelunasan(A *dataPinjaman, n *int) {
	var pilihan int
	fmt.Println("1. 12 Bulan")
	fmt.Println("2. 18 Bulan")
	fmt.Println("3. 24 Bulan")
	fmt.Println("4. 30 Bulan")
	fmt.Println("5. 36 Bulan")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		A[*n].lamaPelunasan = 12
	case 2:
		A[*n].lamaPelunasan = 18
	case 3:
		A[*n].lamaPelunasan = 24
	case 4: 
		A[*n].lamaPelunasan = 30
	case 5:
		A[*n].lamaPelunasan = 36
	}
}

func menu() {
	fmt.Println("Selamat Datang!")
	fmt.Println("1. Ajukan pinjaman")
	fmt.Println("2. Pinjaman anda saat ini")
	fmt.Println("3. Daftar pinjaman saat ini")
	fmt.Println("4. Exit")
}

// a. Pengguna dapat menambahkan, mengubah, dan menghapus data peminjam serta jumlah pinjaman yang diajukan.

// b. Sistem menghitung bunga dan cicilan bulanan berdasarkan jumlah pinjaman dan tenor yang dipilih.

// c. Pengguna dapat mencari data peminjam berdasarkan nama menggunakan Sequential dan Binary Search.

// d. Pengguna dapat mengurutkan daftar peminjam berdasarkan jumlah pinjaman atau tenor menggunakan Selection dan Insertion Sort.

// e. Sistem menampilkan laporan jumlah pinjaman yang diberikan dan status pembayaran.
