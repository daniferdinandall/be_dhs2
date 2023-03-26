package namapackage

import (
	"context"
	"fmt"
	"os"
	"time"

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

// func InsertDHS(mahasiswa Mahasiswa) (InsertedID interface{}) {
func InsertDHS(mahasiswa Mahasiswa, mata_kuliah []MataKuliah) (InsertedID interface{}) {
	var dhs Dhs
	dhs.Mahasiswa = mahasiswa
	dhs.MataKuliah = mata_kuliah
	dhs.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("db_dhs", "dhs", dhs)
}

func GetDhsFromNPM(npm int) (dhs Dhs) {
	data_dhs := MongoConnect("db_dhs").Collection("dhs")
	filter := bson.M{"mahasiswa.npm": npm}
	err := data_dhs.FindOne(context.TODO(), filter).Decode(&dhs)
	if err != nil {
		fmt.Printf("GetDhsFromNPM: %v\n", err)
	}
	return dhs
}

// func GetDhsAll() (dhs Dhs) {
// 	data_dhs := MongoConnect("db_dhs").Collection("dhs")
// 	filter := bson.D{}
// 	var results []Dhs
// 	cur, err := data_dhs.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetDhsFromNPM: %v\n", err)
// 	}
// 	for cur.Next(context.TODO()) {
// 		//Create a value into which the single document can be decoded
// 		var dhs Dhs
// 		err := cur.Decode(&dhs)
// 		if err != nil {
// 			// log.Fatal(err)
// 		}

// 		results = append(results, dhs)

// 	}
// 	return dhs
// }
