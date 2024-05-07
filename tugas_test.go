package _714220054

import (
	"fmt"
	"testing"
	"time"
	"github.com/deviwlndr/undangan_rapat/model"
	"github.com/deviwlndr/undangan_rapat/module"
)

func TestInsertRapatMakrab(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	ukm := "Logic Coffee"
	alamat := "Cafe Logic"
	kepada := model.Undangan_Rapat{
		Kepada : "Devi Wulandari",
	}
	waktu := model.Tanggal{
		Waktu : time.Date(2024, time.March, 30, 10, 0, 0, 0, time.UTC),
	}
	hasil:= module.InsertRapatMakrab(long ,lat , ukm , alamat , kepada , waktu)
	fmt.Println(hasil)
}

func TestInsertTanggal (t *testing.T){
    waktu := time.Date(2024, time.March, 30, 10, 0, 0, 0, time.UTC)
	hasil:=module.InsertTanggal(waktu)
	fmt.Println(hasil)
}

func TestInsertUndanganRapat(t *testing.T){
    kepada := "Devi Wulandari"
	divisi := "Humas"
	hasil:=module.InsertUndanganRapat(kepada, divisi)
	fmt.Println(hasil)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllRapatMakrab()
	fmt.Println(data)
}

