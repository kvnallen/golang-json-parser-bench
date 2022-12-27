package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/francoispqt/gojay"
	jsoniter "github.com/json-iterator/go"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
var fastestIter = jsoniter.ConfigFastest

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func (u *User) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.ID)
	case "name":
		return dec.String(&u.Name)
	case "email":
		return dec.String(&u.Email)
	case "age":
		return dec.Int(&u.Age)
	}
	return nil
}

func (u *User) NKeys() int {
	return 4
}

func main() {
}

func parseJson(data []byte) User {
	var user User
	_ = json.Unmarshal(data, &user)
	return user
}

func parseJsonIterDefault(data []byte) User {
	var user User
	_ = jsoniter.Unmarshal(data, &user)
	return user
}

func parseJsonIter(data []byte) User {
	var user User

	_ = jsonIter.Unmarshal(data, &user)
	return user
}

func parseJsonIterFastest(data []byte) User {
	var user User

	_ = fastestIter.Unmarshal(data, &user)
	return user
}

func parseGoJay(data []byte) User {
	var result User
	_ = gojay.UnmarshalJSONObject(data, &result)
	return result
}

func parseSonic(data []byte) User {
	var result User
	_ = sonic.Unmarshal(data, &result)
	return result
}
