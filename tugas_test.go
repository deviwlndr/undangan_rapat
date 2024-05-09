package _714220054

import (
	"fmt"
	"testing"
	//"time"
	"github.com/deviwlndr/undangan_rapat/model"
	"github.com/deviwlndr/undangan_rapat/module"

	
)

func TestInsertRapatMakrab(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	ukm := "Logic Coffee"
	alamat := "Cafe Logic"
	kepada := model.Undangan_Rapat{
		Kepada : "Irfanto",
	}
	waktu := model.Tanggal{
		Waktu : "senin, 09-10-2024",
	}
	hasil:= module.InsertRapatMakrab(long ,lat , ukm , alamat , kepada , waktu)
	fmt.Println(hasil)
}

func TestInsertTanggal(t *testing.T) {
	waktu := "senin, 09-10-2024"
	hasil := module.InsertTanggal(waktu) // Memanggil fungsi InsertTanggal dari tugas.go
	fmt.Println(hasil)
}

func TestInsertUndanganRapat(t *testing.T){
    kepada := "Irfanto"
	divisi := "Logistik 4"
	hasil:=module.InsertUndanganRapat(kepada, divisi)
	fmt.Println(hasil)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllRapatMakrab()
	fmt.Println(data)
}

