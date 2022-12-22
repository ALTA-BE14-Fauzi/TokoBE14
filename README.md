# TokoBE14

PROGRAM INI MENGGUNAKAN BAHASA PEMOGRAMAN GOLANG DAN SQL

Di dalam program ini terdapat
LOGIN :
(Disini Terdapat Proses Menginputkan username(Nama) dan Password , Jika Cocok di database maka akan ditentukan untuk masuk ke Menu Admin atau Menu Pegawai Berdasarkan Role nya , Role : 1 untuk Admin dan 2 untuk Pegawai)

MENU ADMIN :
-> Tambahkan Pegawai Baru
-> Hapus Data

- Hapus Pegawai
  (Disini Kita Akan Menghapus Pegawai Berdasarkan ID yang Diinput, sebelum di hapus akan di cek apakah ID tersebut cocok dengan data di database , jika ID tidak ada maka akan error)
- Hapus Item/Barang
  (Disini Kita Akan Menghapus Pegawai Berdasarkan ID yang Diinput, sebelum di hapus akan di cek apakah ID tersebut cocok dengan data di database , jika ID tidak ada maka akan error)
- Hapus Transaksi
  (Disini Kita Akan Menghapus Transaksi Berdasarkan ID yang Diinput, sebelum di hapus akan di cek apakah ID tersebut cocok dengan data di database , jika ID tidak ada maka akan error)
- Hapus Customer
  (Disini Kita Akan Menghapus Customer Berdasarkan ID yang Diinput, sebelum di hapus akan di cek apakah ID tersebut cocok dengan data di database , jika ID tidak ada maka akan error)

MENU PEGAWAI :
-> Tambahkan Barang Baru

- Input Nama Barang & Input Stok Barang
  (Disini Kita Akan Menambahkan Barang Baru, sebelum di tambahkan akan di cek apakah nama barang tersebut sudah ada di database atau belum, jika ada maka akan error karena barang yang diinputkan sudah ada di database, jika tidak maka barang tersebut akan masuk ke database)

-> Edit Nama Barang & Stok

- Edit Nama Item/Barang
  (Disini kita akan mengedit nama barang dengan pertama-tama memilih barang tersebut dengan menginputkan namanya, setelah nama diinputkan akan di cek apakah barang tersebut memang ada di database atau belum, jika belum maka error, jika ada maka akan ke proses ubah nama baru dengan menginputkan nama barunya)
- Tambah Stok Barang
  (Disini kita akan menambah stok barang dengan pertama-tama memilih barang tersebut dengan menginputkan namanya, setelah nama diinputkan akan di cek apakah barang tersebut memang ada di database atau belum, jika belum maka error, jika ada maka akan ke proses tambah stok dengan menginputkan berapa stok barang yang akan ditambahkan)

-> Tambahkan Customer Baru

- Input Nama Customer
  (Disini pegawai Akan Menambahkan Customer Baru, sebelum di tambahkan akan di cek apakah nama Customer tersebut sudah ada di database atau belum, jika ada maka akan error karena Customer yang diinputkan sudah ada di database, jika tidak maka Customer tersebut akan masuk ke database)

-> Buat Transaksi

- Input Nama Customer
  (Di proses ini pegawai menambah transaksi baru dengan pertama-tama pegawai menginputkan nama customer yang akan membeli barang, jika customer tidak ada di database maka proses transaksi akan gagal, customer perlu didaftarkan terlebih dahulu. Apabila Customer cocok dan telah didaftarkan di database, maka Transaksi akan dibuat selanjutnya pegawai akan menambahkan barang yang akan dibeli customer di toko)
- Input Nama Barang
  (Di Proses ini pegawai akan menginput nama barang yang akan dibeli oleh customer, pertama tama akan dicek apakah nama barang tersebut ada didatabase atau tidak, jika tidak ada maka akan error sehingga pegawai perlu menginputkan kembali barang dengan nama yang benar atau menginputkan nama barang lain yang akan dibeli customer. Jika nama barang cocok maka akan menghasilkan output OK dana pegawai akan menginputkan nama barang selanjutnya, jika tidak ada barang lagi pegawai akan menginputkan angka 0 untuk exit, selanjutnya data transaksi dan barang yang dibeli akan masuk ke database transaksi item )

-> Lihat Transaksi

- Lihat Transaksi Items Berdasarkan ID Transaksi
  (Pada proses ini kita akan melihat semua transaksi yang pernah dibuat, jika kita ingin melihat transaksi tersebut secara lebih detail, kita akan menginputkan ID transaksi, kemudian akan tampil data yang lebih jelas Pada Tabel Transaksi Item, jika kita menginputkan id yang tidak cocok maka akan error karena ID yang diinputkan tidak ada di database)
