package namapackage

import (
	"context"
	"fmt"
	"os"
	"time"

	model "github.com/daniferdinandall/be_dhs2/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// dhs
func InsertDHS(mahasiswa model.Mahasiswa, mata_kuliah []model.MataKuliah) (InsertedID interface{}) {
	var dhs model.Dhs
	dhs.Mahasiswa = mahasiswa
	dhs.MataKuliah = mata_kuliah
	dhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("db_dhs", "dhs", dhs)
}

func GetDhsFromNPM(npm int) (dhs model.Dhs) {
	data_dhs := MongoConnect("db_dhs").Collection("dhs")
	filter := bson.M{"mahasiswa.npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dhs)
	if err != nil {
		fmt.Printf("GetDhsFromNPM: %v\n", err)
	}
	return dhs
}

func GetDhsAll() (dhs []model.Dhs) {
	data_dhs := MongoConnect("db_dhs").Collection("dhs")
	filter := bson.D{}
	// var results []Dhs
	cur, err := data_dhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetDhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &dhs)
	if err != nil {
		fmt.Println(err)
	}
	return dhs
}

// mahasiswa
func InsertMhs(npm int, nama string, fakultas model.Fakultas, dosen model.Dosen, programStudi model.ProgramStudi) (InsertedID interface{}) {
	var mhs model.Mahasiswa
	mhs.Npm = npm
	mhs.Nama = nama
	mhs.Fakultas = fakultas
	mhs.DosenWali = dosen
	mhs.ProgramStudi = programStudi
	// mhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("db_dhs", "mahasiswa", mhs)
}

func GetMhsFromNPM(npm int) (mhs model.Mahasiswa) {
	data_dhs := MongoConnect("db_dhs").Collection("mahasiswa")
	filter := bson.M{"npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetMhsFromNPM: %v\n", err)
	}
	return mhs
}

func GetMhsAll() (mhs []model.Mahasiswa) {
	data_mhs := MongoConnect("db_dhs").Collection("mahasiswa")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetmhsFromNPM: %v\n", err)
	}
	err = cur.All(context.TODO(), &mhs)
	if err != nil {
		fmt.Println(err)
	}
	return mhs
}

// dosen
func InsertDosen(kode string, nama string, hp string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.KodeDosen = kode
	dosen.Nama = nama
	dosen.PhoneNumber = hp
	return InsertOneDoc("db_dhs", "dosen", dosen)
}

func GetDosenFromKodeDosen(kode string) (dosen model.Dosen) {
	data_dhs := MongoConnect("db_dhs").Collection("dosen")
	filter := bson.M{"kode_dosen": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dosen)
	if err != nil {
		fmt.Printf("GetDosenFromKodeDosen: %v\n", err)
	}
	return dosen
}

func GetDosenAll() (dosen []model.Dosen) {
	data_mhs := MongoConnect("db_dhs").Collection("dosen")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllDosen: %v\n", err)
	}
	err = cur.All(context.TODO(), &dosen)
	if err != nil {
		fmt.Println(err)
	}
	return dosen
}

// matkul
func InsertMatkul(kode string, nama string, sks int, dosen model.Dosen) (InsertedID interface{}) {
	var matkul model.MataKuliah
	matkul.KodeMatkul = kode
	matkul.Nama = nama
	matkul.Sks = sks
	matkul.Dosen = dosen
	return InsertOneDoc("db_dhs", "matkul", matkul)
}

func GetMatkulFromKodeMatkul(kode string) (matkul model.MataKuliah) {
	data_dhs := MongoConnect("db_dhs").Collection("matkul")
	filter := bson.M{"kode_matkul": kode}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&matkul)
	if err != nil {
		fmt.Printf("GetMatkulFromKodeMatkul: %v\n", err)
	}
	return matkul
}

func GetMatkulAll() (matkul []model.MataKuliah) {
	data_mhs := MongoConnect("db_dhs").Collection("matkul")
	filter := bson.D{}
	// var results []mhs
	cur, err := data_mhs.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetMatkulAll: %v\n", err)
	}
	err = cur.All(context.TODO(), &matkul)
	if err != nil {
		fmt.Println(err)
	}
	return matkul
}
