package module

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	model "github.com/bimbingankonseling/backendbk/model"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func GCFHandlerSignup(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	var Response model.Credential
// 	Response.Status = false
// 	var dataUser model.User
// 	err := json.NewDecoder(r.Body).Decode(&dataUser)
// 	if err != nil {
// 		Response.Message = "error parsing application/json: " + err.Error()
// 		return GCFReturnStruct(Response)
// 	}
// 	err = SignUp(conn, collectionname, dataUser)
// 	if err != nil {
// 		Response.Message = err.Error()
// 		return GCFReturnStruct(Response)
// 	}
// 	Response.Status = true
// 	Response.Message = "Halo " + dataUser.Username
// 	return GCFReturnStruct(Response)
// }

// func GCFHandlerSignin(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	var Response model.Credential
// 	Response.Status = false
// 	var dataUser model.User
// 	err := json.NewDecoder(r.Body).Decode(&dataUser)
// 	if err != nil {
// 		Response.Message = "error parsing application/json: " + err.Error()
// 		return GCFReturnStruct(Response)
// 	}
// 	user, status1, err := SignIn(conn, collectionname, dataUser)
// 	if err != nil {
// 		Response.Message = err.Error()
// 		return GCFReturnStruct(Response)
// 	}
// 	Response.Status = true
// 	// Response.Message = "Halo " + dataUser.Name
// 	tokenstring, err := watoken.Encode(dataUser.Email, os.Getenv(PASETOPRIVATEKEYENV))
// 	if err != nil {
// 		Response.Message = "Gagal Encode Token : " + err.Error()
// 	} else {
// 		Response.Message = "Selamat Datang " + user.Email + " di KeeKonseling" + strconv.FormatBool(status1)
// 		Response.Token = tokenstring
// 	}
// 	return GCFReturnStruct(Response)
// }

// PEMASUKAN

func GCFHandlerInsertReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response model.ReservasiResponse
	Response.Status = false
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var dataReservasi model.Reservasi

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	err = json.NewDecoder(r.Body).Decode(&dataReservasi)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}

	_, err = InsertReservasi(mconn, collectionname, dataReservasi.Nama, dataReservasi.No_telp, dataReservasi.TTL, dataReservasi.Status, dataReservasi.Keluhan)
	if err != nil {
		Response.Message = "error inserting Reservasi: " + err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Insert Reservasi success"
	return GCFReturnStruct(Response)
}

func GCFHandlerGetPemasukanFromID(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PemasukanResponse
	Response.Status = false
	var dataUser model.User

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:" + token
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	pemasukan, err := GetPemasukanFromID(dataUser.ID, conn)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Selamat Datang " + dataUser.Email
	Response.Data = []model.Pemasukan{pemasukan}
	return GCFReturnStruct(Response)
}

func GCFHandlerGetAllPemasukan(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PemasukanResponse
	Response.Status = false
	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	pemasukan, err := GetAllPemasukan(conn)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Berhasil mendapatkan semua pemasukan"
	Response.Data = pemasukan
	return GCFReturnStruct(Response)
}

func GCFHandlerUpdatePemasukan(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PemasukanResponse
	Response.Status = false
	var dataPemasukan model.Pemasukan

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	err := json.NewDecoder(r.Body).Decode(&dataPemasukan)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}
	err = UpdatePemasukan(conn, dataPemasukan)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Pemasukan berhasil diupdate"
	Response.Data = []model.Pemasukan{dataPemasukan}
	return GCFReturnStruct(Response)
}

func GCFHandlerDeletePemasukan(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PemasukanResponse
	Response.Status = false
	var dataPemasukan model.Pemasukan

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	err := DeletePemasukan(conn, dataPemasukan)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Pemasukan berhasil dihapus"
	return GCFReturnStruct(Response)
}

// PENGELUARAN

func GCFHandlerInsertPengeluaran(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response model.PengeluaranResponse
	Response.Status = false
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var dataPengeluaran model.Pengeluaran

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		Response.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	err = json.NewDecoder(r.Body).Decode(&dataPengeluaran)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}

	_, err = InsertPengeluaran(mconn, collectionname, dataPengeluaran.Tanggal_keluar, dataPengeluaran.Jumlah_keluar, dataPengeluaran.ID_sumber, dataPengeluaran.Deskripsi, dataPengeluaran.ID_user)
	if err != nil {
		Response.Message = "error inserting Pengeluaran: " + err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = "Insert Pengeluaran success"
	return GCFReturnStruct(Response)
}

func GCFHandlerGetPengeluaranFromID(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PengeluaranResponse
	Response.Status = false
	var dataUser model.User

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:" + token
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	pengeluaran, err := GetPengeluaranFromID(dataUser.ID, conn)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Selamat Datang " + dataUser.Email
	Response.Data = []model.Pengeluaran{pengeluaran}
	return GCFReturnStruct(Response)
}

func GCFHandlerGetAllPengeluaran(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PengeluaranResponse
	Response.Status = false
	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	pengeluaran, err := GetAllPengeluaran(conn)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Berhasil mendapatkan semua pengeluaran"
	Response.Data = pengeluaran
	return GCFReturnStruct(Response)
}

func GCFHandlerUpdatePengeluaran(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PengeluaranResponse
	Response.Status = false
	var dataPengeluaran model.Pengeluaran

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}

	err := json.NewDecoder(r.Body).Decode(&dataPengeluaran)
	if err != nil {
		Response.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(Response)
	}
	err = UpdatePengeluaran(conn, dataPengeluaran)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Pengeluaran berhasil diupdate"
	Response.Data = []model.Pengeluaran{dataPengeluaran}
	return GCFReturnStruct(Response)
}

func GCFHandlerDeletePengeluaran(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.PengeluaranResponse
	Response.Status = false
	var dataPengeluaran model.Pengeluaran

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		Response.Message = "error parsing application/json1:"
		return GCFReturnStruct(Response)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		Response.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(Response)
	}
	err := DeletePengeluaran(conn, dataPengeluaran)
	if err != nil {
		Response.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Pengeluaran berhasil dihapus"
	return GCFReturnStruct(Response)
}

// return
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// get id
func GetID(r *http.Request) string {
	return r.URL.Query().Get("id")
}
