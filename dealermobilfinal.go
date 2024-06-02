package main

import "fmt"

const NMAX int = 1000

type dealer struct {
	idproduk, pabrik, mobil                             string
	harga, tahunkeluar, stokproduk, totalterjual, total int
	lainnya                                             spareparts
}

type spareparts struct {
	koderangka, kodemodel string
}
type tabstring [NMAX]dealer

func main() {
	var ndata, pilihmenu, cariindeks int
	var countinput, switchproses, switchsort int
	var datasearchstr string
	var datasearchint int
	var data tabstring
	for pilihmenu != 8 {
		menu()
		fmt.Print("Pilih Opsi Menu : ")
		fmt.Scan(&pilihmenu)
		if countinput == 0 {
			if pilihmenu > 1 && pilihmenu <= 7 {
				fmt.Println("Harap Masukkan Data Terlebih Dahulu")
			} else if pilihmenu == 1 {
				bacadata(&data, &ndata, 0)
				countinput++
			} else if pilihmenu == 8 {
				fmt.Println("Terima Kasih Karena Sudah Menjalankan Program Ini")
			} else {
				fmt.Println("Opsi Tidak Valid Mohon Masukkan Opsi Yang Benar")
			}

		} else if countinput > 0 {
			if pilihmenu == 1 {
				bacadata(&data, &ndata, 1)
			} else if pilihmenu == 2 {
				fmt.Print("Masukkan ID Mobil Yang Akan Diedit Datanya : ")
				fmt.Scan(&datasearchstr)
				editdata(&data, &ndata, datasearchstr)
			} else if pilihmenu == 3 {
				fmt.Print("Masukkan ID Mobil Yang Akan Dihapus Datanya : ")
				fmt.Scan(&datasearchstr)
				hapusdata(&data, &ndata, datasearchstr)
			} else if pilihmenu == 4 {
				fmt.Println("Pencarian Data Berdasarkan ? :")
				fmt.Printf("%s\n", "1. ID Mobil")
				fmt.Printf("%s\n", "2. Pabrik Mobil")
				fmt.Printf("%s\n", "3. Merek Mobil")
				fmt.Printf("%s\n", "4. Kode Model Mobil")
				fmt.Printf("%s\n", "5. Kode Rangka Mobil")
				fmt.Printf("%s\n", "6. Tahun Keluar Mobil")
				fmt.Printf("%s\n", "7. Stok Mobil")
				fmt.Printf("%s\n", "8. Total Mobil Yang Terjual")
				fmt.Printf("%s\n", "9. Harga Mobil")
				fmt.Print("Opsi Pencarian : ")
				fmt.Scan(&switchproses)
				for switchproses < 1 && switchproses > 9 {
					fmt.Println("Opsi Tidak Valid, Mohon Masukkan Opsi Yang Benar")
					fmt.Print("Opsi Pencarian : ")
					fmt.Scan(&switchproses)
				}
				if switchproses == 1 {
					fmt.Printf("%s", "Masukkan ID Mobil Yang Akan Dicari : ")
				} else if switchproses == 2 {
					fmt.Printf("%s", "Masukkan Pabrik Mobil Yang Akan Dicari : ")
				} else if switchproses == 3 {
					fmt.Printf("%s", "Masukkan Merek Mobil Yang Akan Dicari : ")
				} else if switchproses == 4 {
					fmt.Printf("%s", "Masukkan Kode Model Mobil Yang Akan Dicari : ")
				} else if switchproses == 5 {
					fmt.Printf("%s", "Masukkan Kode Rangka Mobil Yang Akan Dicari : ")
				} else if switchproses == 6 {
					fmt.Printf("%s", "Masukkan Tahun Keluar Yang Akan Dicari : ")
				} else if switchproses == 7 {
					fmt.Printf("%s", "Masukkan Stok Mobil Yang Akan Dicari : ")
				} else if switchproses == 8 {
					fmt.Printf("%s", "Masukkan Total Jumlah Mobil Yang Terjual Yang Akan Dicari : ")
				} else if switchproses == 9 {
					fmt.Printf("%s", "Masukkan Harga Mobil Yang Akan Dicari : ")
				}
				if switchproses >= 1 && switchproses <= 5 {
					fmt.Scan(&datasearchstr)
				} else if switchproses >= 6 && switchproses <= 9 {
					fmt.Scan(&datasearchint)
				}
				if switchproses >= 2 && switchproses <= 9 {
					caridata(&data, &ndata, switchproses, datasearchstr, datasearchint, &cariindeks)
				} else if switchproses == 1 {
					caridata2(&data, ndata, datasearchstr)
				}
			} else if pilihmenu == 5 {
				fmt.Println("Tampilan Data ? :")
				fmt.Printf("%s\n", "1. Semua Data")
				fmt.Printf("%s\n", "2. Sebuah Data")
				fmt.Print("Opsi Tampilan : ")
				fmt.Scan(&switchproses)
				for switchproses < 1 && switchproses > 2 {
					fmt.Println("Opsi Tidak Valid, Mohon Masukkan Opsi Yang Benar")
					fmt.Print("Opsi Pencarian : ")
					fmt.Scan(&switchproses)
				}
				if switchproses == 1 {
					cetakdata(&data, ndata, switchproses, 0)
				} else if switchproses == 2 {
					fmt.Print("Masukkan Id Data Yang Ingin Ditampilkan : ")
					fmt.Scan(&datasearchstr)
					caridata(&data, &ndata, 1, datasearchstr, datasearchint, &cariindeks)
					fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					cetakdata(&data, ndata, switchproses, cariindeks)
				}
			} else if pilihmenu == 6 {
				fmt.Println("Pengurutan Data Secara ? :")
				fmt.Printf("%s\n", "1. Menaik/Ascending")
				fmt.Printf("%s\n", "2. Menurun/Descending")
				fmt.Print("Opsi Pencarian : ")
				fmt.Scan(&switchproses)
				for switchproses < 1 && switchproses > 2 {
					fmt.Println("Opsi Tidak Valid, Mohon Masukkan Opsi Yang Benar")
					fmt.Print("Opsi Pencarian : ")
					fmt.Scan(&switchproses)
				}
				opsisort()
				fmt.Scan(&switchsort)
				sortdata(&data, ndata, switchproses, switchsort)
			} else if pilihmenu == 7 {
				sortdata2(&data, ndata, 1)
				fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
				cariindeks = 0
				for cariindeks < ndata && cariindeks < 3 {
					cetakdata(&data, ndata, 2, cariindeks)
					cariindeks++
				}
			} else if pilihmenu == 8 {
				fmt.Println("Terima Kasih Karena Sudah Menjalankan Program Ini")
			} else if pilihmenu > 8 {
				fmt.Println("Opsi Tidak Valid Mohon Masukkan Opsi Yang Benar")
			}

		}
	}
}

func menu() {
	fmt.Printf("%s\n", "                                    ")
	fmt.Printf("%s\n", "|||||||||||||=> MENU <=|||||||||||||")
	fmt.Printf("%s\n", "|| 1. Input/Tambah Data           ||")
	fmt.Printf("%s\n", "|| 2. Edit Data                   ||")
	fmt.Printf("%s\n", "|| 3. Hapus Data                  ||")
	fmt.Printf("%s\n", "|| 4. Cari Data                   ||")
	fmt.Printf("%s\n", "|| 5. Tampilkan Data              ||")
	fmt.Printf("%s\n", "|| 6. Urutkan Data                ||")
	fmt.Printf("%s\n", "|| 7. Top 3 Penjualan             ||")
	fmt.Printf("%s\n", "|| 8. Exit                        ||")
	fmt.Printf("%s\n", "||||||||||||||||||||||||||||||||||||")
	fmt.Printf("%s\n", "                                    ")
}

func bacadata(M *tabstring, n *int, switchinput int) {
	var idx int
	var inputidproduk string
	var stop int = 0
	if switchinput == 0 {
		i := 0
		*n = 0
		for i < NMAX && stop == 0 {
			idx = -1
			fmt.Print("Mohon Masukkan ID Data Terlebih Dahulu (ID Data Tidak Boleh Sama Dengan Yang Lain): ")
			fmt.Scan(&inputidproduk)
			caridata(M, n, 1, inputidproduk, -1, &idx)
			if idx != -1 {
				for idx != -1 {
					fmt.Println("ID Data Tidak Boleh Sama ID sudah dipakai")
					fmt.Print("Mohon Masukkan ID Data Selain Itu: ")
					fmt.Scan(&inputidproduk)
					caridata(M, n, 1, inputidproduk, -1, &idx)
				}

			}

			if idx == -1 {
				M[i].idproduk = inputidproduk
				fmt.Print("Masukkan Data Nama Pabrik: ")
				fmt.Scan(&M[i].pabrik)
				fmt.Print("Masukkan Data Merek Mobil: ")
				fmt.Scan(&M[i].mobil)
				fmt.Print("Masukkan Kode Model Mobil: ")
				fmt.Scan(&M[i].lainnya.kodemodel)
				fmt.Print("Masukkan Kode Rangka Mobil: ")
				fmt.Scan(&M[i].lainnya.koderangka)
				fmt.Print("Masukkan Tahun Produksi Mobil: ")
				fmt.Scan(&M[i].tahunkeluar)
				fmt.Print("Masukkan Stok Mobil: ")
				fmt.Scan(&M[i].stokproduk)
				fmt.Print("Masukkan Total Mobil Yang Sudah Terjual: ")
				fmt.Scan(&M[i].totalterjual)
				fmt.Print("Masukkan Harga Mobil: ")
				fmt.Scan(&M[i].harga)
				M[i].total = M[i].totalterjual * M[i].harga
				fmt.Println("Data Berhasil Dimasukkan")

			}
			i++
			*n++
			fmt.Print("Ketik Angka Apapun Kecuali 0 Untuk Selesai Memasukkan Data/Ketik 0 Untuk Melanjutkan: ")
			fmt.Scan(&stop)
		}
		fmt.Println("Semua Data Berhasil Dimasukkan")

	} else if switchinput == 1 {
		i := *n
		for i < NMAX && stop == 0 {
			idx = -1
			fmt.Print("Mohon Masukkan ID Data Terlebih Dahulu (ID Data Tidak Boleh Sama Dengan Yang Lain): ")
			fmt.Scan(&inputidproduk)
			caridata(M, n, 1, inputidproduk, -1, &idx)
			if idx != -1 {
				for idx != -1 {
					fmt.Println("ID Data Tidak Boleh Sama ID sudah dipakai")
					fmt.Print("Mohon Masukkan ID Data Selain Itu: ")
					fmt.Scan(&inputidproduk)
					caridata(M, n, 1, inputidproduk, -1, &idx)
				}

			}

			if idx == -1 {
				M[i].idproduk = inputidproduk
				fmt.Print("Masukkan Data Nama Pabrik: ")
				fmt.Scan(&M[i].pabrik)
				fmt.Print("Masukkan Data Merek Mobil: ")
				fmt.Scan(&M[i].mobil)
				fmt.Print("Masukkan Kode Model Mobil: ")
				fmt.Scan(&M[i].lainnya.kodemodel)
				fmt.Print("Masukkan Kode Rangka Mobil: ")
				fmt.Scan(&M[i].lainnya.koderangka)
				fmt.Print("Masukkan Tahun Produksi Mobil: ")
				fmt.Scan(&M[i].tahunkeluar)
				fmt.Print("Masukkan Stok Mobil: ")
				fmt.Scan(&M[i].stokproduk)
				fmt.Print("Masukkan Total Mobil Yang Sudah Terjual: ")
				fmt.Scan(&M[i].totalterjual)
				fmt.Print("Masukkan Harga Mobil: ")
				fmt.Scan(&M[i].harga)
				M[i].total = M[i].totalterjual * M[i].harga
				fmt.Println("Data Berhasil Dimasukkan")
			}
			i++
			*n++
			fmt.Print("Ketik Angka Apapun Kecuali 0 Untuk Selesai Memasukkan Data/Ketik 0 Untuk Melanjutkan: ")
			fmt.Scan(&stop)
		}

		fmt.Println("Semua Data Berhasil Ditambahkan")
	}
}

func caridata(M *tabstring, n *int, switchproses int, datasearchstr string, datasearchint int, cariindeks *int) {
	var i, ketemu, countdata int
	ketemu = -1
	countdata = 0
	if switchproses == 1 {
		for i < *n && ketemu == -1 {
			if M[i].idproduk == datasearchstr {
				ketemu = i
			}
			i++

		}
		*cariindeks = ketemu
	} else if switchproses >= 2 && switchproses <= 9 {
		for i < *n {
			ketemu = -1
			for i < *n && ketemu == -1 {
				if switchproses == 2 && M[i].pabrik == datasearchstr {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 3 && M[i].mobil == datasearchstr {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 4 && M[i].lainnya.kodemodel == datasearchstr {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 5 && M[i].lainnya.koderangka == datasearchstr {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 6 && M[i].tahunkeluar == datasearchint {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 7 && M[i].stokproduk == datasearchint {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 8 && M[i].totalterjual == datasearchint {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				} else if switchproses == 9 && M[i].harga == datasearchint {
					if countdata == 0 {
						fmt.Println("Data Ditemukan : ")
						fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
					}
					ketemu = i
					cetakdata(M, *n, 2, ketemu)
					countdata++
				}
				i++
			}

		}
		if countdata == 0 {
			fmt.Println("Maaf, Data Tidak Ditemukan")
		}
	}
}

func editdata(M *tabstring, n *int, datasearchstr string) {
	var pabrikupdate, merekupdate, kodemodelupdate, koderangkaupdate string
	var hargaupdate, tahunkeluarupdate, jumlahstokupdate, jumlahpenjualanupdate, idx, selesai, switchedit int
	caridata(M, n, 1, datasearchstr, -1, &idx)
	if idx == -1 {
		fmt.Println("Tidak Ada ID Data Yang Bisa Diedit")
	} else if idx != -1 {
		for selesai != 1 {
			fmt.Println("Data Apa Yang Akan Diedit ? :")
			fmt.Printf("%s\n", "1. Pabrik Mobil")
			fmt.Printf("%s\n", "2. Merek Mobil")
			fmt.Printf("%s\n", "3. Kode Model Mobil")
			fmt.Printf("%s\n", "4. Kode Rangka Mobil")
			fmt.Printf("%s\n", "5. Tahun Keluar Mobil")
			fmt.Printf("%s\n", "6. Stok Mobil")
			fmt.Printf("%s\n", "7. Total Mobil Yang Terjual")
			fmt.Printf("%s\n", "8. Harga Mobil")
			fmt.Print("Opsi Pencarian : ")
			fmt.Scan(&switchedit)
			if switchedit == 1 {
				fmt.Print("Masukkan Data Nama Pabrik Terbaru: ")
				fmt.Scan(&pabrikupdate)
				M[idx].pabrik = pabrikupdate
			} else if switchedit == 2 {
				fmt.Print("Masukkan Data Merek Mobil Terbaru: ")
				fmt.Scan(&merekupdate)
				M[idx].mobil = merekupdate
			} else if switchedit == 3 {
				fmt.Print("Masukkan Kode Model Mobil Terbaru: ")
				fmt.Scan(&kodemodelupdate)
				M[idx].lainnya.kodemodel = kodemodelupdate
			} else if switchedit == 4 {
				fmt.Print("Masukkan Kode Rangka Mobil Terbaru: ")
				fmt.Scan(&koderangkaupdate)
				M[idx].lainnya.koderangka = koderangkaupdate
			} else if switchedit == 5 {
				fmt.Print("Masukkan Tahun Produksi Mobil Terbaru: ")
				fmt.Scan(&tahunkeluarupdate)
				M[idx].tahunkeluar = tahunkeluarupdate
			} else if switchedit == 6 {
				fmt.Print("Masukkan Stok Mobil Terbaru: ")
				fmt.Scan(&jumlahstokupdate)
				M[idx].stokproduk = jumlahstokupdate
			} else if switchedit == 7 {
				fmt.Print("Masukkan Total Mobil Yang Sudah Terjual Terbaru: ")
				fmt.Scan(&jumlahpenjualanupdate)
				M[idx].totalterjual = jumlahpenjualanupdate
			} else if switchedit == 8 {
				fmt.Print("Masukkan Harga Mobil Terbaru: ")
				fmt.Scan(&hargaupdate)
				M[idx].harga = hargaupdate
			}
			fmt.Print("Selesai (Ketik 1 Untuk Selesai/Ketik 0 Untuk Melanjutkan Edit Data): ")
			fmt.Scan(&selesai)
		}
		M[idx].total = M[idx].totalterjual * M[idx].harga
		fmt.Println("Data Berhasil Diedit")
	}

}

func hapusdata(M *tabstring, n *int, datasearchstr string) {
	var idx int
	caridata(M, n, 1, datasearchstr, -1, &idx)
	if idx == -1 {
		fmt.Println("Tidak Ada ID Data Yang Bisa Dihapus")
	} else {
		for i := idx; i < *n-1; i++ {
			M[i] = M[i+1]
		}
		*n--
		fmt.Println("Data Berhasil Dihapus")
	}
}

func sortdata(M *tabstring, n int, switchsort int, switchdata int) {
	var i, j, idx int
	var temp dealer
	if switchsort == 1 {
		for i = 1; i < n; i++ {
			idx = i - 1
			for j = i; j < n; j++ {
				if switchdata == 1 && M[idx].idproduk > M[j].idproduk {
					idx = j
				} else if switchdata == 2 && M[idx].pabrik > M[j].pabrik {
					idx = j
				} else if switchdata == 3 && M[idx].mobil > M[j].mobil {
					idx = j
				} else if switchdata == 4 && M[idx].lainnya.kodemodel > M[j].lainnya.kodemodel {
					idx = j
				} else if switchdata == 5 && M[idx].lainnya.koderangka > M[j].lainnya.koderangka {
					idx = j
				} else if switchdata == 6 && M[idx].tahunkeluar > M[j].tahunkeluar {
					idx = j
				} else if switchdata == 7 && M[idx].stokproduk > M[j].stokproduk {
					idx = j
				} else if switchdata == 8 && M[idx].totalterjual > M[j].totalterjual {
					idx = j
				} else if switchdata == 9 && M[idx].harga > M[j].harga {
					idx = j
				} else if switchdata == 10 && M[idx].total > M[j].total {
					idx = j
				}
			}
			temp = M[idx]
			M[idx] = M[i-1]
			M[i-1] = temp
		}
	} else if switchsort == 2 {
		for i = 1; i < n; i++ {
			idx = i - 1
			for j = i; j < n; j++ {
				if switchdata == 1 && M[idx].idproduk < M[j].idproduk {
					idx = j
				} else if switchdata == 2 && M[idx].pabrik < M[j].pabrik {
					idx = j
				} else if switchdata == 3 && M[idx].mobil < M[j].mobil {
					idx = j
				} else if switchdata == 4 && M[idx].lainnya.kodemodel < M[j].lainnya.kodemodel {
					idx = j
				} else if switchdata == 5 && M[idx].lainnya.koderangka < M[j].lainnya.koderangka {
					idx = j
				} else if switchdata == 6 && M[idx].tahunkeluar < M[j].tahunkeluar {
					idx = j
				} else if switchdata == 7 && M[idx].stokproduk < M[j].stokproduk {
					idx = j
				} else if switchdata == 8 && M[idx].totalterjual < M[j].totalterjual {
					idx = j
				} else if switchdata == 9 && M[idx].harga < M[j].harga {
					idx = j
				} else if switchdata == 10 && M[idx].total < M[j].total {
					idx = j
				}
			}
			temp = M[idx]
			M[idx] = M[i-1]
			M[i-1] = temp
		}
	}
	cetakdata(M, n, 1, 0)
}

func cetakdata(M *tabstring, n int, switchcetak int, x int) {
	if switchcetak == 1 {
		fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
		for i := 0; i < n; i++ {
			fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14d %-14d %-14d %2s %-14d %2s %-14d\n", M[i].idproduk, M[i].pabrik, M[i].mobil, M[i].lainnya.kodemodel, M[i].lainnya.koderangka, M[i].tahunkeluar, M[i].stokproduk, M[i].totalterjual, "Rp", M[i].harga, "Rp", M[i].total)
		}
	} else if switchcetak == 2 {
		fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14d %-14d %-14d %2s %-14d %2s %-14d\n", M[x].idproduk, M[x].pabrik, M[x].mobil, M[x].lainnya.kodemodel, M[x].lainnya.koderangka, M[x].tahunkeluar, M[x].stokproduk, M[x].totalterjual, "Rp", M[x].harga, "Rp", M[x].total)
	}
}

func sortdata2(M *tabstring, n int, switchsort int) {
	var i, j int
	var temp dealer
	if switchsort == 1 {
		for i = 1; i < n; i++ {
			j = i
			temp = M[i]
			for j > 0 && M[j-1].total < temp.total {
				M[j] = M[j-1]
				j--
			}
			M[j] = temp
		}
	} else if switchsort == 2 {
		for i = 1; i < n; i++ {
			j = i
			temp = M[i]
			for j > 0 && M[j-1].idproduk < temp.idproduk {
				M[j] = M[j-1]
				j--
			}
			M[j] = temp
		}
	}
}

func opsisort() {
	fmt.Println("Pengurutan Data Berdasarkan ? :")
	fmt.Printf("%s\n", "1. ID Mobil")
	fmt.Printf("%s\n", "2. Pabrik Mobil")
	fmt.Printf("%s\n", "3. Merek Mobil")
	fmt.Printf("%s\n", "4. Kode Model Mobil")
	fmt.Printf("%s\n", "5. Kode Rangka Mobil")
	fmt.Printf("%s\n", "6. Tahun Keluar Mobil")
	fmt.Printf("%s\n", "7. Stok Mobil")
	fmt.Printf("%s\n", "8. Total Mobil Yang Terjual")
	fmt.Printf("%s\n", "9. Harga Mobil")
	fmt.Printf("%s\n", "10. Total Hasil Penjualan Mobil")
	fmt.Print("Opsi Pencarian : ")
}

func caridata2(M *tabstring, n int, datasearchstr string) {
	var bawah, atas, tengah, ketemu int
	ketemu = -1
	sortdata2(M, n, 2)
	bawah = 0
	atas = n - 1
	for bawah <= atas && ketemu == -1 {
		tengah = (bawah + atas) / 2
		if datasearchstr < M[tengah].idproduk {
			bawah = tengah + 1
		} else if datasearchstr > M[tengah].idproduk {
			atas = tengah - 1
		} else {
			ketemu = tengah
		}

	}
	if ketemu != -1 {
		fmt.Println("Data Ditemukan : ")
		fmt.Printf("%-4s %-14s %-14s %-14s %-14s %-14s %-14s %-14s %-17s %-16s\n", "Id", "Nama Pabrik", "Merek Mobil", "Kode Model", "Kode Rangka", "Tahun Keluar", "Stok Mobil", "Total Terjual", "Harga Mobil", "Total")
		cetakdata(M, n, 2, ketemu)
	} else if ketemu == -1 {
		fmt.Println("Maaf, Data Tidak Ditemukan")
	}
}
