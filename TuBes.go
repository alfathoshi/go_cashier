package main

import (
	"fmt"
)

const NMAX int = 100

type barang struct {
	nama     string
	harga    float64
	qty      int
	kategori string
}

type penjualan struct {
	modal      float64
	pendapatan float64
	data       barang
	tanggal    string
	hargaJual  float64
}

type arrBarang [NMAX]barang

type arrPenjualan [NMAX]penjualan

func main() {
	var B arrBarang
	var T arrPenjualan
	var nT, nB int
	var x string
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Print("Press anything to continue.....")
	fmt.Scanln(&x)
	fmt.Println()
	mainMenu(&B, &nB, &T, &nT)
}

func menuHeader() {
	fmt.Println("###	---------------------------------------------       ###")
	fmt.Println("###		    Aplikasi Jual Beli			    ###")
	fmt.Println("###		    Created by Group 19			    ###")
	fmt.Println("###	Muhammad Bintang Al-Fath & Prabusti Alwan Fauzan    ###")
	fmt.Println("###		Algoritma Pemrograman 2023		    ###")
	fmt.Println("###	---------------------------------------------	    ###")
}

func mainMenu(B *arrBarang, nB *int, T *arrPenjualan, nT *int) {
	var pilihan string
	var qty int
	var harga float64
	var nama, kategori string
	menuHeader()
	fmt.Println()

	fmt.Println("###	Menu Utama     ###")
	fmt.Println()

	fmt.Println("1. Menu Transaksi")
	fmt.Println("2. Menu Barang")
	fmt.Println("3. Keluar")
	fmt.Println()
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	for pilihan != "3" && pilihan != "2" && pilihan != "1" {
		fmt.Print("Pilihan :")
		fmt.Scanln(&pilihan)
	}

	if pilihan == "1" {
		menuTransaction(T, nT, B, nB)
	} else if pilihan == "2" {
		itemMenu(B, nB, T, nT, nama, kategori, harga, qty)
	}
}

func itemMenu(B *arrBarang, nB *int, T *arrPenjualan, nT *int, nama, kategori string, harga float64, qty int) {
	var pilihan string
	var x string
	fmt.Println()
	menuHeader()
	fmt.Println()

	fmt.Println("###	Menu Admin	 ###")
	fmt.Println()

	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Lihat Barang")
	fmt.Println("3. Cari Barang")
	fmt.Println("4. Kembali")
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	for pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
		fmt.Println("Inputan tidak valid!")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
	}

	if pilihan == "1" {
		addItem(B, nB, T, nT, nama, kategori, harga, qty)
	} else if pilihan == "2" {
		viewItem(B, nB, T, nT)
	} else if pilihan == "3" {
		searchItem(*B, *nB, *T, *nT, x)
	} else if pilihan == "4" {
		mainMenu(B, nB, T, nT)
	}
}

func addItem(B *arrBarang, n *int, T *arrPenjualan, nT *int, nama, kategori string, harga float64, kuantitas int) {
	var pilihan string
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Println("###	Menu Tambah Barang	###")
	fmt.Println()
	fmt.Println("Masukan data barang : (nama, harga, stok, kategori)")
	fmt.Println()
	fmt.Scanln(&nama, &harga, &kuantitas, &kategori)
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	fmt.Println("1. Simpan	2. Batal")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	for pilihan != "1" && pilihan != "2" {
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
	}
	if *n < NMAX {
		if pilihan == "1" {
			B[*n].nama = nama
			B[*n].harga = harga
			B[*n].qty = kuantitas
			B[*n].kategori = kategori
			for i := 0; i < *n; i++ {
				if nama == B[i].nama && harga == B[i].harga && kategori == B[i].kategori {
					B[i].nama = nama
					B[i].harga = harga
					B[i].qty += kuantitas
					B[i].kategori = kategori
					*n--
				}
			}
			*n++
			fmt.Println()
			fmt.Println("Data barang tersimpan")
		}
	} else {
		fmt.Println("Maaf, array sudah penuh")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
	itemMenu(B, n, T, nT, nama, kategori, harga, kuantitas)
}

func viewItem(B *arrBarang, n *int, T *arrPenjualan, nT *int) {
	var pilihan, nama, kategori string
	var qty int
	var harga float64
	fmt.Println()
	menuHeader()
	fmt.Println()

	fmt.Println("### Menu Tampilkan Barang ###")
	fmt.Println()

	fmt.Println("List data barang")
	fmt.Println()
	if *n != 0 {
		for i := 0; i < *n; i++ {
			fmt.Println(i+1, ". ", "| Nama :", B[i].nama, "| Harga :", B[i].harga, "| Stok :", B[i].qty, "| Kategori :", B[i].kategori, "|")
		}
		fmt.Println()
		fmt.Println("###	--------------------------------------------- 	    ###")
		fmt.Println()
		fmt.Println("Opsi : ")
		fmt.Println("1. Edit	2. Urutkan		3.Hapus		4. Kembali")
		fmt.Println()
		fmt.Println("###	--------------------------------------------- 	    ###")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
		for pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
			fmt.Print("Pilihan : ")
			fmt.Scanln(&pilihan)
		}
		if pilihan == "1" {
			editItem(B, n, T, nT)
		} else if pilihan == "2" {
			fmt.Println()
			fmt.Println("1. Nama	2. Harga		3.Stok		4. Kategori")
			fmt.Print("Urutkan berdasarkan : ")
			fmt.Scanln(&pilihan)
			sortItem(*B, *n, *T, *nT, pilihan)
			fmt.Println()
			fmt.Print("0. Kembali.............")
			fmt.Scanln(&pilihan)
		} else if pilihan == "3" {
			deleteItem(B, n, T, nT)

		}
	} else {
		fmt.Println()
		fmt.Println("Data barang kosong")
		fmt.Println("Silahkan tambahkan data barang ")
		fmt.Println()
		fmt.Print("0. Kembali.............")
		fmt.Scanln(&pilihan)
	}
	itemMenu(B, n, T, nT, nama, kategori, harga, qty)
}

func deleteItem(B *arrBarang, n *int, T *arrPenjualan, nT *int) {
	var i, idx int
	var pilihan string
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Println("### Menu Hapus Barang ###")
	fmt.Println()
	fmt.Println("List data barang")
	fmt.Println()

	if *n != 0 {

		for i := 0; i < *n; i++ {
			fmt.Println(i+1, ". ", "| Nama :", B[i].nama, "| Harga :", B[i].harga, "| Stok :", B[i].qty, "| Kategori :", B[i].kategori, "|")
		}
		fmt.Println()

		fmt.Print("Pilih barang yang akan dihapus : ")
		fmt.Scanln(&idx)
		if idx > 0 && idx <= *n {
			for i = idx - 1; i < *n; i++ {
				B[i] = B[i+1]
			}
			*n--
			fmt.Println("Data barang telah dihapus")
		} else {
			fmt.Println("Tidak ada data barang yang dihapus")
		}
	} else {
		fmt.Println()
		fmt.Println("Data barang kosong")
		fmt.Println("Silahkan tambahkan data barang ")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
}

func editItem(B *arrBarang, n *int, T *arrPenjualan, nT *int) {
	var pilihan string
	var idx int
	fmt.Println()
	menuHeader()
	fmt.Println()

	fmt.Println("### Menu Edit Barang ###")
	fmt.Println()

	if *n != 0 {
		for i := 0; i < *n; i++ {
			fmt.Println(i+1, ". ", "| Nama :", B[i].nama, "| Harga :", B[i].harga, "| Stok :", B[i].qty, "| Kategori :", B[i].kategori, "|")
		}
		fmt.Println()
		fmt.Println("Pilih data barang yang akan diedit")
		fmt.Println()
		fmt.Print("Data ke : ")
		fmt.Scanln(&idx)
		if idx > 0 && idx <= *n {
			fmt.Println()
			fmt.Println("Pilih bagian yang akan diedit.")
			fmt.Println()
			fmt.Println("1. Nama		2. Harga		3.Stok		4. Kategori")
			fmt.Println()
			fmt.Print("Pilih bagian yang akan diedit : ")
			fmt.Scanln(&pilihan)
			for pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
				fmt.Println("Input tidak valid!")
				fmt.Print("Pilih bagian yang akan diedit : ")
				fmt.Scanln(&pilihan)
			}
			if pilihan == "1" {
				var editNama string
				fmt.Println()
				fmt.Println("Data sekarang :", B[idx-1].nama)
				fmt.Println()
				fmt.Print("Masukan data baru :")
				fmt.Scanln(&editNama)
				fmt.Println()
				fmt.Println("###	--------------------------------------------- 	    ###")
				fmt.Println()
				fmt.Println("1. Simpan	2. Batal")
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
				for pilihan != "1" && pilihan != "2" {
					fmt.Print("Pilihan :")
					fmt.Scanln(&pilihan)
				}
				if pilihan == "1" {
					var found int = -1
					for i := 0; i < *n; i++ {
						if editNama == B[i].nama {
							found = i
						}
					}
					var i int
					if found != -1 {
						B[found].qty += B[idx-1].qty
						for i = idx - 1; i < *n; i++ {
							B[i] = B[i+1]
						}
						*n--
					} else {
						B[idx-1].nama = editNama
					}
				}
			} else if pilihan == "2" {
				var editHarga float64
				fmt.Println()
				fmt.Println("Data sekarang :", B[idx-1].harga)
				fmt.Println()
				fmt.Print("Masukan data baru : ")
				fmt.Scanln(&editHarga)
				fmt.Println()
				fmt.Println("###	--------------------------------------------- 	    ###")
				fmt.Println()
				fmt.Println("1. Simpan	2. Batal")
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
				for pilihan != "1" && pilihan != "2" {
					fmt.Print("Pilihan :")
					fmt.Scanln(&pilihan)
				}
				if pilihan == "1" {
					B[idx-1].harga = editHarga
				}
			} else if pilihan == "3" {
				var editStok int
				fmt.Println()
				fmt.Println("Data sekarang :", B[idx-1].qty)
				fmt.Println()
				fmt.Print("Masukan data baru :")
				fmt.Scanln(&editStok)
				fmt.Println()
				fmt.Println("###	--------------------------------------------- 	    ###")
				fmt.Println()
				fmt.Println("1. Simpan	2. Batal")
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
				for pilihan != "1" && pilihan != "2" {
					fmt.Print("Pilihan :")
					fmt.Scanln(&pilihan)
				}
				if pilihan == "1" {
					B[idx-1].qty = editStok
				}
			} else if pilihan == "4" {
				var editKategori string
				fmt.Println()
				fmt.Println("Data sekarang :", B[idx-1].kategori)
				fmt.Println()
				fmt.Print("Masukan data baru :")
				fmt.Scanln(&editKategori)
				fmt.Println()
				fmt.Println("###	--------------------------------------------- 	    ###")
				fmt.Println()
				fmt.Println("1. Simpan	2. Batal")
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
				for pilihan != "1" && pilihan != "2" {
					fmt.Print("Pilihan :")
					fmt.Scanln(&pilihan)
				}
				if pilihan == "1" {
					B[idx-1].kategori = editKategori
				}
			}
			fmt.Println()
			fmt.Println("Data barang telah diperbarui!")
		} else {
			fmt.Println("Tidak ada data barang yang diedit")
		}
	} else {
		fmt.Println()
		fmt.Println("Data barang kosong")
		fmt.Println("Silahkan tambahkan data barang ")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
}

func sortItem(B arrBarang, n int, T arrPenjualan, nT int, x string) {
	var i, j, idx int
	var temp barang
	var pilihan string
	fmt.Println()
	fmt.Println("1. A-Z		2. Z-A")
	fmt.Print("Urutkan berdasarkan : ")
	fmt.Scanln(&pilihan)
	if pilihan == "1" {
		i = 1
		for i <= n-1 {
			idx = i - 1
			j = i
			if x == "1" {
				for j < n {
					if B[idx].nama > B[j].nama {
						idx = j
					}
					j++
				}
			} else if x == "2" {
				for j < n {
					if B[idx].harga > B[j].harga {
						idx = j
					}
					j++
				}
			} else if x == "3" {
				for j < n {
					if B[idx].qty > B[j].qty {
						idx = j
					}
					j++
				}
			} else if x == "4" {
				for j < n {
					if B[idx].kategori > B[j].kategori {
						idx = j
					}
					j++
				}
			}
			temp = B[idx]
			B[idx] = B[i-1]
			B[i-1] = temp
			i++
		}
	} else if pilihan == "2" {
		i = 1
		for i <= n-1 {
			idx = i - 1
			j = i
			if x == "1" {
				for j < n {
					if B[idx].nama < B[j].nama {
						idx = j
					}
					j++
				}
			} else if x == "2" {
				for j < n {
					if B[idx].harga < B[j].harga {
						idx = j
					}
					j++
				}
			} else if x == "3" {
				for j < n {
					if B[idx].qty < B[j].qty {
						idx = j
					}
					j++
				}
			} else if x == "4" {
				for j < n {
					if B[idx].kategori < B[j].kategori {
						idx = j
					}
					j++
				}
			}
			temp = B[idx]
			B[idx] = B[i-1]
			B[i-1] = temp
			i++
		}
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ". ", "| Nama :", B[i].nama, "| Harga :", B[i].harga, "| Stok :", B[i].qty, "| Kategori :", B[i].kategori, "|")
	}
}

func searchItem(B arrBarang, n int, T arrPenjualan, nT int, x string) {
	var idx int = -1
	var qty int
	var harga float64
	var pilihan, nama, kategori string
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Println("### Menu Cari Barang ###")
	fmt.Println()
	fmt.Println("Cari berdasarkan :")
	fmt.Println()
	fmt.Println("1. Nama	2. Kategori")
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	fmt.Print("Pilihan :")
	fmt.Scanln(&pilihan)
	for pilihan != "1" && pilihan != "2" {
		fmt.Print("Pilihan :")
		fmt.Scanln(&pilihan)
	}
	if pilihan == "1" {
		var temp barang
		var i, j int
		i = 1
		for i <= n-1 {
			j = i
			temp = B[j]
			for j > 0 && temp.nama < B[j-1].nama {
				B[j] = B[j-1]
				j--
			}
			B[j] = temp
			i++
		}
		fmt.Println()
		fmt.Print("Input nama barang : ")
		fmt.Scanln(&x)
		var med int
		var left int = 0
		var right int = n - 1
		for left <= right && idx == -1 {
			med = (left + right) / 2
			if x < B[med].nama {
				right = med - 1
			} else if x > B[med].nama {
				left = med + 1
			} else {
				idx = med
			}
		}
		if idx == -1 {
			fmt.Println()
			fmt.Println("Data barang tidak ditemukan")
		} else {
			fmt.Println()
			fmt.Println(idx+1, ". ", "| Nama :", B[idx].nama, "| Harga :", B[idx].harga, "| Stok :", B[idx].qty, "| Kategori :", B[idx].kategori, "|")
		}
	} else if pilihan == "2" {
		var temp barang
		var i, j int
		i = 1
		for i <= n-1 {
			j = i
			temp = B[j]
			for j > 0 && temp.kategori < B[j-1].kategori {
				B[j] = B[j-1]
				j--
			}
			B[j] = temp
			i++
		}
		fmt.Println()
		fmt.Println("List kategori barang :")
		fmt.Println()
		for i := 0; i < n; i++ {
			fmt.Println(i+1, ".", B[i].kategori)
		}
		fmt.Println()
		fmt.Print("Input kategori barang : ")
		fmt.Scanln(&x)
		fmt.Println()
		for i = 0; i < n; i++ {
			if B[i].kategori == x {
				idx = i
				fmt.Println(i+1, ". ", "| Nama :", B[idx].nama, "| Harga :", B[idx].harga, "| Stok :", B[idx].qty, "| Kategori :", B[idx].kategori, "|")
			}
		}
		if idx == -1 {
			fmt.Println()
			fmt.Println("Data barang tidak ditemukan")
		}
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
	itemMenu(&B, &n, &T, &nT, nama, kategori, harga, qty)

}

func viewTransaction(T arrPenjualan, nT int, B arrBarang, nB int) {
	var pilihan string
	var labaKotor, modal, temp float64
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Println("### Data Transaksi Penjualan ###")
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	for i := 0; i < nT; i++ {
		labaKotor += T[i].pendapatan
		temp += T[i].data.harga * float64(T[i].data.qty)
	}
	modal = transactionCapital(B, nB) + (labaKotor - (labaKotor - temp))
	fmt.Println("Modal :", modal)
	if nT != 0 {
		fmt.Println()
		for i := 0; i < nT; i++ {
			fmt.Println(i+1, ".", "| Tanggal Pembelian :", T[i].tanggal, "| Nama Barang:", T[i].data.nama, "| QTY :", T[i].data.qty, "| Harga Barang :", T[i].data.harga, "| Harga Jual :", T[i].hargaJual, "| Pemasukan :", T[i].pendapatan, "|")
		}
		fmt.Println()
		fmt.Println("Laba Kotor :", labaKotor)
		fmt.Println("Laba bersih :", labaKotor-modal)
		fmt.Println()

		fmt.Println("###	--------------------------------------------- 	    ###")
		fmt.Println()
		fmt.Println("Opsi : ")
		fmt.Println("1. Edit		2. Hapus		3.Top 5		4. Kembali")
		fmt.Println()
		fmt.Println("###	--------------------------------------------- 	    ###")
		fmt.Println()
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
		for pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
			fmt.Print("Pilihan : ")
			fmt.Scanln(&pilihan)
		}
		if pilihan == "1" {
			editTransaction(&T, &nT, &B, &nB)
		} else if pilihan == "3" {
			fmt.Println()
			top5(T, nT)
			fmt.Println()
		} else if pilihan == "2" {
			deleteTransaction(&T, &nT, &B, &nB)
		}
	} else {
		fmt.Println()
		fmt.Println("Data trasaksi kosong")
		fmt.Println("Silahkan tambahkan data transaksi")
	}
	menuTransaction(&T, &nT, &B, &nB)
}

func addTransaction(T *arrPenjualan, nT *int, B *arrBarang, nB *int) {
	var nama, pilihan, tanggal string
	var qty int
	var bayar float64
	var idx int = -1
	fmt.Println()
	menuHeader()
	fmt.Println("### Transkasi ###")
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	fmt.Println("Tanggal Pembelian : (DD/MM/YYYY)")
	fmt.Scanln(&tanggal)
	fmt.Println()
	fmt.Print("Barang yang dibeli : ")
	fmt.Scanln(&nama)
	for i := 0; i < *nB; i++ {
		if B[i].nama == nama {
			idx = i
		}
	}
	if idx != -1 {
		if B[idx].qty != 0 {
			fmt.Println("Harga :", (B[idx].harga*1/10)+B[idx].harga)
			fmt.Println("Stok :", B[idx].qty)
			fmt.Println()
			fmt.Print("Jumlah yang dibeli : ")
			fmt.Scanln(&qty)
			fmt.Println()
			if qty <= B[idx].qty {
				fmt.Println("Jumlah yang harus dibayar :", ((B[idx].harga*1/10)+B[idx].harga)*float64(qty))
				fmt.Println()
				fmt.Print("Jumlah bayar : ")
				fmt.Scanln(&bayar)
				fmt.Println()
				if bayar >= ((B[idx].harga*1/10)*float64(qty))+B[idx].harga*float64(qty) {
					fmt.Println("###	--------------------------------------------- 	    ###")
					fmt.Println("1. Simpan	2. Batal")
					fmt.Print("Pilihan : ")
					fmt.Scanln(&pilihan)
					for pilihan != "1" && pilihan != "2" {
						fmt.Print("Pilihan : ")
						fmt.Scanln(&pilihan)
					}
					if *nT < NMAX {
						if pilihan == "1" {
							B[idx].qty -= qty
							T[*nT].pendapatan = ((B[idx].harga * 1 / 10) * float64(qty)) + B[idx].harga*float64(qty)
							T[*nT].tanggal = tanggal
							T[*nT].data.nama = nama
							T[*nT].data.kategori = B[idx].kategori
							T[*nT].data.harga = B[idx].harga
							T[*nT].data.qty = qty
							T[*nT].hargaJual = (B[idx].harga * 1 / 10) + B[idx].harga
							fmt.Println()
							fmt.Println("Kembalian :", bayar-T[*nT].pendapatan)
							fmt.Println()
							fmt.Println("Sisa barang :", B[idx].qty)
							fmt.Println("Total pendapatan :", T[*nT].pendapatan)
							*nT++
						}
					} else {
						fmt.Println("Maaf, array sudah penuh")
					}
				} else {
					fmt.Println("Uang tidak cukup")
				}
			} else {
				fmt.Println("Stok kurang dari jumlah yang diminta")
			}
		} else {
			fmt.Println("Stok barang kosong")
		}
	} else {
		fmt.Println("Barang tidak ada")

	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
	menuTransaction(T, nT, B, nB)
}

func deleteTransaction(T *arrPenjualan, nT *int, B *arrBarang, nB *int) {
	var i, idx int
	var pilihan string
	fmt.Println()
	menuHeader()
	fmt.Println("### Menu Hapus Transaksi ###")
	fmt.Println()
	fmt.Println("List data transaksi")
	fmt.Println()
	for i := 0; i < *nT; i++ {
		fmt.Println(i+1, ".", "| Tanggal Pembelian :", T[i].tanggal, "| Nama Barang:", T[i].data.nama, "| QTY :", T[i].data.qty, "| Harga Barang :", T[i].data.harga, "| Harga Jual :", T[i].hargaJual, "| Pemasukan :", T[i].pendapatan, "|")
	}
	fmt.Println()
	fmt.Print("Pilih barang yang akan dihapus : ")
	fmt.Scanln(&idx)
	if idx > 0 && idx <= *nT {
		for i = idx - 1; i < *nT; i++ {
			T[i] = T[i+1]
		}
		*nT--
		fmt.Println()
		fmt.Println("Data transaksi telah dihapus")
	} else {
		fmt.Println("Tidak ada data transaksi yang dihapus")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
}

func editTransaction(T *arrPenjualan, nT *int, B *arrBarang, nB *int) {
	var pilihan string
	var idx int
	fmt.Println()
	menuHeader()
	fmt.Println()

	fmt.Println("### Menu Edit Transaksi ###")
	fmt.Println()
	for i := 0; i < *nT; i++ {
		fmt.Println(i+1, ".", "| Tanggal Pembelian :", T[i].tanggal, "| Nama Barang:", T[i].data.nama, "| QTY :", T[i].data.qty, "| Harga Barang :", T[i].data.harga, "| Harga Jual :", T[i].hargaJual, "| Pemasukan :", T[i].pendapatan, "|")
	}
	fmt.Println()
	fmt.Println("Pilih data transaksi yang akan diedit")
	fmt.Println()
	fmt.Print("Data ke : ")
	fmt.Scanln(&idx)
	if idx > 0 && idx <= *nT {
		fmt.Println()
		fmt.Println("Pilih bagian yang akan diedit.")
		fmt.Println()
		fmt.Println("1. Tanggal		2. Nama barang		3.Jumlah")
		fmt.Println()
		fmt.Println("###	--------------------------------------------- 	    ###")
		fmt.Print("Pilih bagian yang akan diedit : ")
		fmt.Scanln(&pilihan)
		for pilihan != "1" && pilihan != "2" && pilihan != "3" {
			fmt.Println("Input tidak valid!")
			fmt.Print("Pilih bagian yang akan diedit : ")
			fmt.Scanln(&pilihan)
		}
		if pilihan == "1" {
			var editTanggal string
			fmt.Println("Data sekarang :", T[idx-1].tanggal)
			fmt.Print("Masukan data baru :")
			fmt.Scanln(&editTanggal)
			fmt.Println()
			fmt.Println("###	--------------------------------------------- 	    ###")
			fmt.Println()
			fmt.Println("1. Simpan	2. Batal")
			fmt.Print("Pilihan :")
			fmt.Scanln(&pilihan)
			for pilihan != "1" && pilihan != "2" {
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
			}
			if pilihan == "1" {

				T[idx-1].tanggal = editTanggal
			}

		} else if pilihan == "2" {
			var editNama string
			fmt.Println("Data sekarang :", T[idx-1].data.nama)
			fmt.Print("Masukan data baru :")
			fmt.Scanln(&editNama)
			var cari int = -1
			for i := 0; i < *nB; i++ {
				if B[i].nama == T[idx-1].data.nama {
					cari = i
				}
			}
			var found int = -1
			for i := 0; i < *nB; i++ {
				if B[i].nama == editNama {
					found = i
				}
			}
			if found != -1 {
				var totalQTY int
				for i := 0; i < *nT; i++ {
					if B[cari].nama == T[i].data.nama {
						totalQTY += T[i].data.qty
					}
				}
				fmt.Println("###	--------------------------------------------- 	    ###")
				fmt.Println("1. Simpan	2. Batal")
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
				for pilihan != "1" && pilihan != "2" {
					fmt.Print("Pilihan :")
					fmt.Scanln(&pilihan)
				}
				if pilihan == "1" {
					if totalQTY > T[found].data.qty {
						B[found].qty -= T[idx-1].data.qty
						B[cari].qty += T[idx].data.qty
						T[idx-1].data.nama = editNama
						T[idx-1].data.harga = B[found].harga
						T[idx-1].hargaJual = B[found].harga + (B[found].harga * 0.1)
						T[idx-1].pendapatan = float64(T[idx-1].data.qty) * T[idx-1].hargaJual
						fmt.Println("Data barang telah diperbarui!")
					} else {
						fmt.Println()
						fmt.Println("Stok tidak cukup")
					}
				}
			} else {
				fmt.Println()
				fmt.Println("Nama barang tidak ada pada data barang")
				menuTransaction(T, nT, B, nB)
			}
		} else if pilihan == "3" {
			var editStok int
			fmt.Println("Data sekarang :", T[idx-1].data.qty)
			temp := T[idx-1].data.qty
			fmt.Print("Masukan data baru :")
			fmt.Scanln(&editStok)
			fmt.Println("###	--------------------------------------------- 	    ###")
			fmt.Println("1. Simpan	2. Batal")
			fmt.Print("Pilihan :")
			fmt.Scanln(&pilihan)
			for pilihan != "1" && pilihan != "2" {
				fmt.Print("Pilihan :")
				fmt.Scanln(&pilihan)
			}
			var found int = -1
			var totalQTY int
			for i := 0; i < *nB; i++ {
				if B[i].nama == T[idx-1].data.nama {
					found = i
				}
			}
			if pilihan == "1" {
				for i := 0; i < *nT; i++ {
					if B[found].nama == T[i].data.nama {
						totalQTY += T[i].data.qty + B[found].qty
					}
				}
				if totalQTY > editStok && B[found].qty > editStok {
					T[idx-1].data.qty = editStok
					B[found].qty += temp - editStok
					T[idx-1].pendapatan = float64(T[idx].data.qty) * T[idx-1].hargaJual
					fmt.Println()
					fmt.Println("Data barang telah diperbarui!")
				} else {
					fmt.Println()
					fmt.Println("Stok tidak cukup")
				}
			}
		}
	} else {
		fmt.Println("Tidak ada data barang yang diedit")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
	menuTransaction(T, nT, B, nB)
}

func top5(T arrPenjualan, nT int) {
	var i, j, k, idx int
	var pilihan string
	var temp penjualan
	if nT != 0 {
		for i := 0; i < nT; i++ {
			j = i + 1
			for j < nT {
				if T[i].data.nama == T[j].data.nama {
					T[i].data.qty += T[j].data.qty
					k++
				}
				j++
			}
		}
		i = 1
		for i <= nT-1 {
			idx = i - 1
			j = i
			for j < nT {
				if T[idx].data.qty < T[j].data.qty {
					idx = j
				}
				j++
			}
			temp = T[idx]
			T[idx] = T[i-1]
			T[i-1] = temp
			i++
		}
		for k > 0 {
			if idx > 0 && idx <= nT {
				for i = idx - 1; i < nT; i++ {
					T[i] = T[i+1]
				}
				nT--
			}
			k--
		}
		fmt.Println("Top 5 Data Barang Paling Banyak Terjual")
		fmt.Println()
		for i := 0; i < 5; i++ {
			fmt.Println(i+1, ".", "| Nama Barang:", T[i].data.nama, "| Harga Barang :", T[i].data.harga, "| QTY :", T[i].data.qty, "| Kategori :", T[i].data.kategori, "|")
		}
	} else {
		fmt.Println("Data transaksi kosong")
	}
	fmt.Println()
	fmt.Print("0. Kembali.............")
	fmt.Scanln(&pilihan)
}

func transactionCapital(B arrBarang, n int) float64 {
	var modal float64
	for i := 0; i < n; i++ {
		modal += B[i].harga * float64(B[i].qty)
	}
	return modal
}

func menuTransaction(T *arrPenjualan, nT *int, B *arrBarang, nB *int) {
	var pilihan string
	fmt.Println()
	menuHeader()
	fmt.Println()
	fmt.Println("### Menu Transaksi ###")
	fmt.Println()
	fmt.Println("1. Pembelian	2. Lihat Data Transaksi		3. Kembali")
	fmt.Println()
	fmt.Println("###	--------------------------------------------- 	    ###")
	fmt.Println()
	fmt.Print("Pilihan :")
	fmt.Scanln(&pilihan)
	for pilihan != "1" && pilihan != "2" && pilihan != "3" {
		fmt.Print("Pilihan :")
		fmt.Scanln(&pilihan)
	}
	if pilihan == "1" {
		addTransaction(T, nT, B, nB)
	} else if pilihan == "2" {
		viewTransaction(*T, *nT, *B, *nB)
	} else if pilihan == "3" {
		mainMenu(B, nB, T, nT)
	}
}
