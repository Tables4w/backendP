package types

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type Form struct {
	Fio      string `json:"Fio"`
	Tel      string `json:"Tel"`
	Email    string `json:"Email"`
	Date     string `json:"Date"`
	Gender   string `json:"Gender"`
	Favlangs []int  `json:"Favlangs"`
	Bio      string `json:"Bio"`
}

type FormErrors struct {
	Fio      string `json:"Fio"`
	Tel      string `json:"Tel"`
	Email    string `json:"Email"`
	Date     string `json:"Date"`
	Gender   string `json:"Gender"`
	Favlangs string `json:"Favlangs"`
	Bio      string `json:"Bio"`
	Familiar string `json:"Familiar"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SALT")+password), 14)
	return string(bytes), err
}

// if ok return nil; else error
func CheckPassword(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(os.Getenv("SALT")+password))
	if err != nil {
		return errors.New("username or password invalid")
	}
	return nil
}
