package password

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func Generate(minLength int, maxLength int) string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength)))
	if err != nil {
		panic(err) // handle this gracefully
	}
	length.Add(length, big.NewInt(int64(minLength)))

	intLength := int(length.Int64())

	newPassword := make([]byte, intLength)
	randomData := make([]byte, intLength+intLength/4)
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPassword[i] = chars[c%clen]
			i++
			if i == intLength {
				return string(newPassword)
			}
		}
	}
}
func ConvertToHash(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func CheckHash(currentPasswordHash string, password string) bool {
	isTheSame := bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(password))

	if isTheSame != nil {
		return false
	}
	return true
}
