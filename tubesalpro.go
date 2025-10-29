//komentar nich
package main

import "fmt"

const NMAX int = 100

type comment struct{
	pesan string
	like int
}

type tabCmnt [NMAX]comment

func main () {
	var data tabCmnt
	var nData int
	var x1 string
	fmt.Print("masukan jumlah komentar: ")
	fmt.Scan(&nData)
	fmt.Print("masukkan kata yang akan dicari: ")
	fmt.Scan(&x1)
	bacaKomentar(&data,nData)
	fmt.Println()
	cetakKomentar(data, nData)
	if sequentialSearch(data,nData,x1){
		freq := frekuensiKata(data,nData,x1)
		fmt.Printf("ditemukan sebanyak %d kata %s ",freq, x1)
	}else{
		fmt.Printf("kata tidak ditemukan")
	}
	fmt.Printf("hasil analisis dari komentar tersebut: %s ", analisisSentimen(data,nData,x1))
}

func bacaKomentar(A *tabCmnt, n int){
	var i int
	fmt.Println("Masukkan bintang: ")
	fmt.Scan(&A[i].like)
	fmt.Println("Masukkan komentar: ")
	for i := 0; i < n; i++{
		fmt.Scan(&A[i].pesan)
	}
}

func cetakKomentar(A tabCmnt, n int){
	var i int
	fmt.Printf("like: %d",A[i].like)
	fmt.Println()
	fmt.Println("komentar: ")
	for i := 0; i < n ; i++{
		fmt.Printf("%s ",A[i].pesan)
		fmt.Println()
	}
}


func frekuensiKata(A tabCmnt,n int, x string) int {
	freq := 0
	for i := 0; i < n; i++ {
		if A[i].pesan == x {
			freq++
		}
	}
	return freq
}

func sequentialSearch(A tabCmnt, n int, x string) bool{
	var ketemu bool
	var i int
	ketemu = false
	i = 0
	
	for !ketemu && i < n {
		ketemu = A[i].pesan == x
		i++
	}
	return ketemu
}

func analisisSentimen(A tabCmnt,n int, x string) string{
	var kata string
	var i int
	i = 0
	for i := 0; i < n; i++ {
		if A[i].pesan = x == "bagus", "keren", "mantap" {
			return "positif"
		}else if A[i].pesan = x == "jelek", "kurang", "aneh" {
			return "negatif"
		}
	}
	return "netral"
}