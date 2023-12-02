package konseling

import (
	"fmt"
	"testing"

	model "github.com/bimbingankonseling/backendbk/model"
	module "github.com/bimbingankonseling/backendbk/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "konseling")

// PEMASUKAN

func TestInsertReservasi(t *testing.T) {
	var doc model.Reservasi
	doc.Nama = "bella"
	doc.No_telp = "081388765778"
	// Isi ID_sumber dan ID_user sesuai kebutuhan, contoh:
	doc.TTL = "23 mei 2003"
	doc.Status = "mahasiswa"
	doc.Keluhan = "pusing"

	hasil, err := module.InsertReservasi(db, "reservasi", doc.Nama, doc.No_telp, doc.TTL, doc.Status, doc.Keluhan)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Printf("Data berhasil disimpan dengan id %s\n", hasil.Hex())
	}
	fmt.Println(hasil)
}

// func TestGetAllPemasukan(t *testing.T) {
// 	var docs []model.Pemasukan
// 	docs, err := module.GetAllPemasukan(db)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan id :", docs)
// 	}
// 	fmt.Println(docs)
// }

// func TestGetPemasukanFromID(t *testing.T) {
// 	id := "6565676bb3e79ceef0540910"
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Errorf("Error getting document: %v", err)
// 	} else {
// 		user, err := module.GetPemasukanFromID(objectId, db)
// 		if err != nil {
// 			t.Errorf("Error getting document: %v", err)
// 		} else {
// 			fmt.Println(user)
// 		}
// 	}
// }

// func TestUpdatePemasukan(t *testing.T) {
// 	var doc model.Pemasukan
// 	doc.Tanggal_masuk = "22/02/2023"
// 	doc.Jumlah_masuk = 230000
// 	doc.ID_sumber.ID, _ = primitive.ObjectIDFromHex("65643f6242ef94271017c94a")
// 	doc.Deskripsi = "dari joki ngoding"
// 	doc.ID_user.ID, _ = primitive.ObjectIDFromHex("65631b4de009209dea4dc55e")
// 	id, err := primitive.ObjectIDFromHex("6565676bb3e79ceef0540910")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil diubah dengan id")
// 	} else {
// 		err = module.UpdatePemasukan(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil diubah dengan id :", doc.ID)
// 		}
// 	}
// }

// func TestUpdatePemasukan(t *testing.T) {
// 	var doc model.Pemasukan
// 	doc.Tanggal_masuk = "22/02/2023"
// 	doc.Jumlah_masuk = 230000
// 	doc.ID_sumber.Nama_sumber = "Freelance"
// 	doc.Deskripsi = "dari joki ngoding"
// 	doc.ID_user.Name = "Fedhira"
// 	id, err := primitive.ObjectIDFromHex("6565676bb3e79ceef0540910")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil diubah dengan id")
// 	} else {

// 		err = module.UpdatePemasukan(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil diubah dengan id :", doc.ID)
// 		}
// 	}

// }

// func TestDeletePemasukan(t *testing.T) {
// 	var doc model.Pemasukan
// 	id, err := primitive.ObjectIDFromHex("6564639a6e6e2f66eee84ddd")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil dihapus dengan id")
// 	} else {

// 		err = module.DeletePemasukan(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil dihapus dengan id :", doc.ID)
// 		}
// 	}
// }

// PENGELUARAN

// func TestInsertPengeluaran(t *testing.T) {
// 	var doc model.Pengeluaran
// 	doc.Tanggal_keluar = "26/02/2023"
// 	doc.Jumlah_keluar = 50000
// 	// Isi ID_sumber dan ID_user sesuai kebutuhan, contoh:
// 	doc.ID_sumber = model.Sumber{Nama_sumber: "Konsumsi"}
// 	doc.ID_user = model.User{Name: "Fedhira Syaila"}

// 	hasil, err := module.InsertPengeluaran(db, "pengeluaran", doc.Tanggal_keluar, doc.Jumlah_keluar, doc.ID_sumber, doc.Deskripsi, doc.ID_user)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Printf("Data berhasil disimpan dengan id %s\n", hasil.Hex())
// 	}
// 	fmt.Println(hasil)
// }

// func TestGetAllPengeluaran(t *testing.T) {
// 	var docs []model.Pengeluaran
// 	docs, err := module.GetAllPengeluaran(db)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan id :", docs)
// 	}
// 	fmt.Println(docs)
// }

// func TestGetPengeluaranFromID(t *testing.T) {
// 	id := "65646471f789492812e11a7a"
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Errorf("Error getting document: %v", err)
// 	} else {
// 		user, err := module.GetPengeluaranFromID(objectId, db)
// 		if err != nil {
// 			t.Errorf("Error getting document: %v", err)
// 		} else {
// 			fmt.Println(user)
// 		}
// 	}
// }

// func TestUpdatePemasukan(t *testing.T) {
// 	var doc model.Pemasukan
// 	doc.Tanggal_masuk = "22/02/2023"
// 	doc.Jumlah_masuk = 230000
// 	doc.ID_sumber.ID, _ = primitive.ObjectIDFromHex("65643f6242ef94271017c94a")
// 	doc.Deskripsi = "dari joki ngoding"
// 	doc.ID_user.ID, _ = primitive.ObjectIDFromHex("65631b4de009209dea4dc55e")
// 	id, err := primitive.ObjectIDFromHex("6565676bb3e79ceef0540910")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil diubah dengan id")
// 	} else {
// 		err = module.UpdatePemasukan(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil diubah dengan id :", doc.ID)
// 		}
// 	}
// }

// func TestUpdatePengeluaran(t *testing.T) {
// 	var doc model.Pengeluaran
// 	doc.Tanggal_keluar = "22/02/2023"
// 	doc.Jumlah_keluar = 230000
// 	doc.ID_sumber.Nama_sumber = "Kesehatan"
// 	doc.Deskripsi = "ke rs"
// 	doc.ID_user.Name = "Fedhira"
// 	id, err := primitive.ObjectIDFromHex("656464a7b23a402327223b46")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil diubah dengan id")
// 	} else {

// 		err = module.UpdatePengeluaran(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil diubah dengan id :", doc.ID)
// 		}
// 	}

// }

// func TestDeletePengeluaran(t *testing.T) {
// 	var doc model.Pengeluaran
// 	id, err := primitive.ObjectIDFromHex("6565774763f64428805965ef")
// 	doc.ID = id
// 	if err != nil {
// 		fmt.Printf("Data tidak berhasil dihapus dengan id")
// 	} else {

// 		err = module.DeletePengeluaran(db, doc)
// 		if err != nil {
// 			t.Errorf("Error updateting document: %v", err)
// 		} else {
// 			fmt.Println("Data berhasil dihapus dengan id :", doc.ID)
// 		}
// 	}
// }

// TEST GET USER
// func TestGetUserFromID(t *testing.T) {
// 	id := "65631b4de009209dea4dc55e"
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		t.Errorf("Error getting document: %v", err)
// 	} else {
// 		user, err := module.GetUserFromID(objectId, db)
// 		if err != nil {
// 			t.Errorf("Error getting document: %v", err)
// 		} else {
// 			fmt.Println("Welcome back:", user)
// 		}
// 	}

// }



// func TestGetUserFromEmail(t *testing.T) {
// 	email := "yellow12@gmail.com"
// 	user, err := module.GetUserFromEmail(email, db)
// 	if err != nil {
// 		t.Errorf("Error getting user: %v", err)
// 	} else {
// 		fmt.Println("Welcome back:", user)
// 	}
// }

// func TestGetUserFromName(t *testing.T) {
// 	name := "Fedhira Syaila"
// 	user, err := module.GetUserFromName(name, db)
// 	if err != nil {
// 		t.Errorf("Error getting user: %v", err)
// 	} else {
// 		fmt.Println("Welcome back:", user)
// 	}
// }



// func TestGeneratePasswordHash(t *testing.T) {
// 	password := "yellow"
// 	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
// 	fmt.Println("Password:", password)
// 	fmt.Println("Hash:    ", hash)

// 	match := CheckPasswordHash(password, hash)
// 	fmt.Println("Match:   ", match)
// }
// func TestGeneratePrivateKeyPaseto(t *testing.T) {
// 	privateKey, publicKey := watoken.GenerateKey()
// 	fmt.Println("privateKey")
// 	fmt.Println(privateKey)
// 	fmt.Println("publicKey")
// 	fmt.Println(publicKey)
// 	fmt.Println("HASIL")
// 	hasil, err := watoken.Encode("Billblis", privateKey)
// 	fmt.Println(hasil, err)
// }

// func TestHashFunction(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "billblis")
// 	var userdata User
// 	userdata.Username = "renjun"
// 	userdata.Password = "yellow"

// 	filter := bson.M{"username": userdata.Username}
// 	res := atdb.GetOneDoc[User](mconn, "user", filter)
// 	fmt.Println("Mongo User Result: ", res)
// 	hash, _ := HashPassword(userdata.Password)
// 	fmt.Println("Hash Password : ", hash)
// 	match := CheckPasswordHash(userdata.Password, res.Password)
// 	fmt.Println("Match:   ", match)

// }

// func TestIsPasswordValid(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "billblis")
// 	var userdata User
// 	userdata.Username = "renjun"
// 	userdata.Password = "yellow"

// 	anu := IsPasswordValid(mconn, "user", userdata)
// 	fmt.Println(anu)
// }

// func TestInsertUser(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "billblis")
// 	var userdata User
// 	userdata.Username = "renjun"
// 	userdata.Password = "yellow"

// 	nama := InsertUser(mconn, "user", userdata)
// 	fmt.Println(nama)
// }