package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// presensi
type Karyawan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja   []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

// ================================================================================================
// DHS
type Dhs struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mahasiswa  Mahasiswa          `bson:"mahasiswa,omitempty" json:"mahasiswa,omitempty"`
	MataKuliah []MataKuliah       `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm          int                `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Fakultas     Fakultas           `bson:"fakultas,omitempty" json:"fakultas,omitempty"`
	ProgramStudi ProgramStudi       `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	DosenWali    Dosen              `bson:"dosen_wali,omitempty" json:"dosen_wali,omitempty"`
}

type MataKuliah struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeMatkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama       string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Dosen      Dosen              `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Sks        int                `bson:"sks,omitempty" json:"sks,omitempty"`
	Nilai      string             `bson:"nilai,omitempty" json:"nilai,omitempty"`
}

type Dosen struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeDosen   string             `bson:"kode_dosen,omitempty" json:"kode_dosen,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type Fakultas struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeFakultas string             `bson:"kode_fakultas,omitempty" json:"kode_fakultas,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type ProgramStudi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KodeProgramStudi string             `bson:"kode_program_studi,omitempty" json:"kode_program_studi,omitempty"`
	Nama             string             `bson:"nama,omitempty" json:"nama,omitempty"`
}
