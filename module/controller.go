package module

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/bimbingankonseling/backendbk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// CRUD
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertManyDocsRegistrasi(db *mongo.Database, col string, registrasi []model.Registrasi) (insertedIDs []primitive.ObjectID, err error) {
	var interfaces []interface{}
	for _, registrasi := range registrasi {
		interfaces = append(interfaces, registrasi)
	}
	result, err := db.Collection(col).InsertMany(context.Background(), interfaces)
	if err != nil {
		return insertedIDs, fmt.Errorf("kesalahan server: insert")
	}
	for _, id := range result.InsertedIDs {
		insertedIDs = append(insertedIDs, id.(primitive.ObjectID))
	}
	return insertedIDs, nil
}

func InsertManyDocsReservasi(db *mongo.Database, col string, reservasi []model.Reservasi) (insertedIDs []primitive.ObjectID, err error) {
	var interfaces []interface{}
	for _, reservasi := range reservasi {
		interfaces = append(interfaces, reservasi)
	}
	result, err := db.Collection(col).InsertMany(context.Background(), interfaces)
	if err != nil {
		return insertedIDs, fmt.Errorf("kesalahan server: insert")
	}
	for _, id := range result.InsertedIDs {
		insertedIDs = append(insertedIDs, id.(primitive.ObjectID))
	}
	return insertedIDs, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// SIGN IN
func SignIn(db *mongo.Database, col string, insertedDoc model.User) (user model.User, Status bool, err error) {
	if insertedDoc.Username == "" || insertedDoc.Password == "" {
		return user, false, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Username); err != nil {
		return user, false, fmt.Errorf("Username tidak valid")
	}
	if !CheckPasswordHash(insertedDoc.Password, existsDoc.Password) {
		return user, false, fmt.Errorf("password salah")
	}

	return existsDoc, true, nil
}

// REGISTRASI

func InsertRegistrasi(db *mongo.Database, col string, nama_lengkap string, no_telp string, ttl string, nim string, alamat string) (insertedID primitive.ObjectID, err error) {
	registrasi := bson.M{
		"nama_lengkap"	: nama_lengkap,
		"no_telp"		: no_telp,
		"ttl"			: ttl,
		"nim"			: nim,
		"alamat"		: alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), registrasi)
	if err != nil {
		fmt.Printf("InsertRegistrasi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllRegistrasi(db *mongo.Database) (registrasi []model.Registrasi, err error) {
	collection := db.Collection("registrasi")
	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return registrasi, fmt.Errorf("error GetAllRegistrasi mongo: %s", err)
	}

	// Iterate through the cursor and decode each document
	for cursor.Next(context.Background()) {
		var p model.Registrasi
		if err := cursor.Decode(&p); err != nil {
			return registrasi, fmt.Errorf("error decoding document: %s", err)
		}
		registrasi = append(registrasi, p)
	}

	if err := cursor.Err(); err != nil {
		return registrasi, fmt.Errorf("error during cursor iteration: %s", err)
	}

	return registrasi, nil
}

// RESERVASI

func InsertReservasi(db *mongo.Database, col string, nama string, no_telp string, ttl string, status string, keluhan string) (insertedID primitive.ObjectID, err error) {
	reservasi := bson.M{
		"nama"		: nama,
		"no_telp"	: no_telp,
		"ttl"		: ttl,
		"status"	: status,
		"keluhan"	: keluhan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), reservasi)
	if err != nil {
		fmt.Printf("InsertReservasi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllReservasi(db *mongo.Database) (reservasi []model.Reservasi, err error) {
	collection := db.Collection("reservasi")
	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return reservasi, fmt.Errorf("error GetAllReservasi mongo: %s", err)
	}

	// Iterate through the cursor and decode each document
	for cursor.Next(context.Background()) {
		var p model.Reservasi
		if err := cursor.Decode(&p); err != nil {
			return reservasi, fmt.Errorf("error decoding document: %s", err)
		}
		reservasi = append(reservasi, p)
	}

	if err := cursor.Err(); err != nil {
		return reservasi, fmt.Errorf("error during cursor iteration: %s", err)
	}

	return reservasi, nil
}

//func GetPengeluaranFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Pengeluaran, err error) {
//	collection := db.Collection("pengeluaran")
//	filter := bson.M{"_id": _id}
//	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return doc, fmt.Errorf("_id tidak ditemukan")
//		}
//		return doc, fmt.Errorf("kesalahan server")
//	}
//	return doc, nil
//}

func UpdateReservasi(db *mongo.Database, doc model.Reservasi) (err error) {
	filter := bson.M{"_id": doc.ID}
	result, err := db.Collection("reservasi").UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		fmt.Printf("UpdateReservasi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified id")
		return
	}
	return nil
}

func DeleteReservasi(db *mongo.Database, doc model.Reservasi) error {
	collection := db.Collection("reservasi")
	filter := bson.M{"_id": doc.ID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", doc.ID, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", doc.ID)
	}

	return nil
}
