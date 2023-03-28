package namapackage

import (
	"fmt"
	"testing"

	model "github.com/daniferdinandall/be_dhs2/model"
	module "github.com/daniferdinandall/be_dhs2/module"
)

func TestInsertPresensi(t *testing.T) {
	mahasiswa := model.Mahasiswa{
		Npm:  1214036,
		Nama: "Erdito Nausha Adam",
		Fakultas: model.Fakultas{
			Nama: "Sekolah Vokasi",
		},
		ProgramStudi: model.ProgramStudi{
			Nama: "D4 Teknik Informatika",
		},
		DosenWali: model.Dosen{
			Nama: " Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
		},
	}
	mata_kuliah := []model.MataKuliah{
		{
			KodeMatkul: "TI41061",
			Nama:       "ALGORITMA DAN STRUKTUR DATA I",
			Sks:        3,
			Nilai:      "B",
		}, {
			KodeMatkul: "TI41092",
			Nama:       "ALJABAR LINIER",
			Sks:        2,
			Nilai:      "AB",
		}, {
			KodeMatkul: "PPI01040",
			Nama:       "BAHASA INDONESIA",
			Sks:        2,
			Nilai:      "A",
		}, {
			KodeMatkul: "TI42011",
			Nama:       "LITERASI TEKNOLOGI",
			Sks:        2,
			Nilai:      "A",
		}, {
			KodeMatkul: "TI41071",
			Nama:       "ALGORITMAPEMOGRAMAN I",
			Sks:        3,
			Nilai:      "AB",
		},
	}

	hasil := module.InsertDHS(mahasiswa, mata_kuliah)
	fmt.Println(hasil)
}

func TestGetDhsFromNPM(t *testing.T) {
	npm := 1214049
	biodata := module.GetDhsFromNPM(npm)
	fmt.Println(biodata)
}

func TestGetDhsAll(t *testing.T) {
	biodata := module.GetDhsAll()
	fmt.Println(biodata)
}

// mhs
func TestInsertMhs(t *testing.T) {
	npm := 1214049
	nama := "Auliana Fahrian Bani Ridwan"
	fakultas := model.Fakultas{
		Nama: "Sekolah Vokasi",
	}
	programStudi := model.ProgramStudi{
		Nama: "D4 Teknik Informatika",
	}
	dosen := model.Dosen{
		Nama: " Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
	}

	hasil := module.InsertMhs(npm, nama, fakultas, dosen, programStudi)
	fmt.Println(hasil)
}

func TestGetMhsFromNPM(t *testing.T) {
	npm := 1214049
	biodata := module.GetMhsFromNPM(npm)
	fmt.Println(biodata)
}

func TestGetMhsAll(t *testing.T) {
	biodata := module.GetMhsAll()
	fmt.Println(biodata)
}

// dosen
func TestInsertDosen(t *testing.T) {
	kode := "003"
	nama := "NISA HANUM HARANI"
	hp := "09876215321"

	hasil := module.InsertDosen(kode, nama, hp)
	fmt.Println(hasil)
}

func TestGetDosenFromKodeDosen(t *testing.T) {
	kode := "001"
	biodata := module.GetDosenFromKodeDosen(kode)
	fmt.Println(biodata)
}

func TestGetDosenAll(t *testing.T) {
	biodata := module.GetDosenAll()
	fmt.Println(biodata)
}

// dosen
func TestInsertMatkul(t *testing.T) {
	kode := "TI41092"
	nama := "LITERASI TEKNOLOGI"
	sks := 3
	dosen := model.Dosen{
		Nama: "Roni Habibi, S.Kom., M.T., SFPC",
	}

	hasil := module.InsertMatkul(kode, nama, sks, dosen)
	fmt.Println(hasil)
}

func TestMatkulFromKodeMatkul(t *testing.T) {
	kode := "TI41061"
	biodata := module.GetMatkulFromKodeMatkul(kode)
	fmt.Println(biodata)
}

func TestMatkulAll(t *testing.T) {
	biodata := module.GetMatkulAll()
	fmt.Println(biodata)
}
