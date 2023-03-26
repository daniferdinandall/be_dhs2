package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	mahasiswa := Mahasiswa{
		Npm:  1214050,
		Nama: "Dani Ferdinan",
		Fakultas: Fakultas{
			Nama: "Sekolah Vokasi",
		},
		ProgramStudi: ProgramStudi{
			Nama: "D4 Teknik Informatika",
		},
		DosenWali: Dosen{
			Nama: " Rd. NURAINI SITI FATHONAH, S.S., M.Hum.,SFPC",
		},
	}
	mata_kuliah := []MataKuliah{
		{
			KodeMatkul: "TI41061",
			Nama:       "ALGORITMA DAN STRUKTUR DATA I ",
			Sks:        3,
			Nilai:      "A",
		}, {
			KodeMatkul: "TI41092",
			Nama:       "ALJABAR LINIER",
			Sks:        2,
			Nilai:      "A",
		}, {
			KodeMatkul: "PPI01040",
			Nama:       "BAHASA INDONESIAI ",
			Sks:        2,
			Nilai:      "B",
		}, {
			KodeMatkul: "TI42011",
			Nama:       "LITERASI TEKNOLOGI",
			Sks:        2,
			Nilai:      "B",
		}, {
			KodeMatkul: "TI41071",
			Nama:       "ALGORITMAPEMOGRAMAN I",
			Sks:        3,
			Nilai:      "A",
		},
	}

	hasil := InsertDHS(mahasiswa, mata_kuliah)
	fmt.Println(hasil)
}

func TestGetDhsFromNPM(t *testing.T) {
	npm := 1214049
	biodata := GetDhsFromNPM(npm)
	fmt.Println(biodata)
}

// func TestGetDhsAll(t *testing.T) {
// 	biodata := GetDhsAll()
// 	fmt.Println(biodata.Mahasiswa.Nama)
// }
