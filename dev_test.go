package konseling

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// func TestUpdateGetData(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "billblis")
// 	datagedung := GetAllBangunanLineString(mconn, "billblis")
// 	fmt.Println(datagedung)
// }

func TestGeneratePasswordHash(t *testing.T) {
	password := "youzanna"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("gabriella", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "konseling")
	var userdata User
	userdata.Username = "gabriella"
	userdata.Password = "youzanna"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "konseling")
	var userdata User
	userdata.Username = "gabriella"
	userdata.Password = "youzanna"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "konseling")
	var userdata User
	userdata.Username = "gabriella"
	userdata.Password = "youzanna"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}