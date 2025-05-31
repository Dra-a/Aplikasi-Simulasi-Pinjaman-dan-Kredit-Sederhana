# Aplikasi-Simulasi-Pinjaman-dan-Kredit-Sederhana
Aplikasi ini mensimulasikan sistem pinjaman dan kredit sederhana dengan skema bunga tetap atau variabel. Data utama yang digunakan adalah daftar peminjam, jumlah pinjaman, tenor, dan status pembayaran. Pengguna aplikasi adalah individu yang ingin mensimulasikan sistem pinjaman secara virtual.

Aplikasi ini digunakan untuk mencatat dan mengelola proyek freelance yang sedang dikerjakan oleh pengguna. Data utama yang digunakan adalah daftar proyek, klien, deadline, dan status proyek. Pengguna aplikasi adalah pekerja lepas yang ingin melacak perkembangan proyek mereka. 

# Spesifikasi
1. Pengguna dapat menambahkan, mengubah, dan menghapus proyek freelance yang sedang atau telah dikerjakan.
2. Pengguna dapat memperbarui status proyek (misalnya: sedang dikerjakan, selesai, pending).
3. Pengguna dapat mencari proyek berdasarkan nama klien atau nama proyek menggunakan Sequential dan Binary Search.
4. Pengguna dapat mengurutkan daftar proyek berdasarkan deadline atau bayaran tertinggi menggunakan Selection dan Insertion Sort.
5. Sistem menampilkan laporan proyek yang sudah selesai dan yang masih berjalan dalam bentuk tabel atau ringkasan.

# Fitur
## 1. Login/Logout
Akan ditampilkan halaman login saat aplikasi pertama kali dibuka.
Pengguna dapat masuk dalam aplikasi dengan menginputkan nama, nama tersebut akan disimpan ke dalam program untuk digunakan pada prosedur dan fungsi lain. Pengguna kemudian akan masuk ke menu utama.
Dengan menginputkan 7 pada menu utama,
Pengguna akan logout, sehingga nama dalam program akan direset.
## 2. Ajuan Pinjaman
Dengan menginputkan 1 pada menu utama,
Pengguna dapat mengajukan pinjaman dari opsi jumlah pinjaman dan lama pinjaman yang diberikan. Ketentuan dari bunga dan denda dari pinjaman tersebut akan disampaikan kepada pengguna, jika pengguna bersedia untuk mengikuti ketentuan tersebut, maka data ajuan akan diproses untuk masuk ke dalam array utama.
## 3. Lihat Pinjaman Pengguna
Dengan menginputkan 2 pada menu utama,
Program akan menampilkan data pinjaman dengan nama peminjam yang sama, dengan nama yang telah diinput pengguna pada halaman login.
Pengguna dapat memilih untuk mengubah atau menghapus data tersebut.
## 4. Edit dan Hapus Data
Jika pengguna memilih untuk mengedit data, data pengguna akan ditampilkan lagi, kemudian pengguna dapat menginputkan nomor dari data yang ingin diubah. Pengguna kemudian akan diminta untuk menginputkan ulang data baru yang diinginkan.
Jika pengguna memilih untuk menghapus data, data pengguna akan ditampilkan lagi, kemudian pengguna dapat menginputkan nomor dari data yang ingin dihapus. Data tersebut kemudian akan dihapus oleh program.
## 5. Tampilkan dan Urutkan Seluruh Data
Dengan menginputkan 3 pada menu utama,
Pengguna akan dapat melihat seluruh data yang ada di dalam program. Pengguna kemudian dapat mengurutkan data tersebut berdasarkan jumlah pinjamannya, atau berdasarkan nama peminjam.
## 6. Cari Data Seseorang
Dengan menginputkan 4 pada menu utama,
Pengguna akan diminta untuk menginputkan nama dari peminjam yang ingin dicari. Jika data tersebut ditemukan, program akan menampilkan data tersebut.
## 7. Bayar Tagihan Pengguna
Dengan menginputkan 5 pada menu utama,
Pengguna akan diarahkan untuk dapat membayar tagihannya. Program akan menampilkan data dengan nama user, yang status pembayaran(sudahBayar)nya masih bernilai false. Jika pengguna ingin membayar tagihan tersebut, program akan menampilkan data tersebut agar user dapat memilih tagihan yang ingin dibayarkan sesuai nomor. Setelah suatu data pinjaman terpilih, program akan menunjukkan besar tagihan yang perlu dibayarkan oleh pengguna, berdasarkan jumlah pinjaman awal, bunga, dan ada/tidaknya denda.
## 8. Simulasi 30 Hari ke Depan
Dengan menginputkan 6 pada menu utama,
Program akan menjalankan prosedur untuk mengurangi 1 bulan dari sisa pinjaman seluruh data. Dari data tersebut, jika ada data yang belum dibayarkan, maka akan dikenakan denda. Status pembayaran tiap data kemudian dikembalikan menjadi false untuk bulan yang baru.
Program juga akan menjalankan prosedur untuk mengisi beberapa data dummy.
