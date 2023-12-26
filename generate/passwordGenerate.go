package generate

import (
	_ "PasswordGenerate/leng"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+[{]}|;:'\",<.>/?"

func GeneratePassword(length int) string {
	var sb strings.Builder
	max := big.NewInt(int64(len(characters)))

	for i := 0; i < length; i++ {
		// generate random index
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			fmt.Println("Error generating password", err)
			os.Exit(1)
		}

		character := characters[idx.Int64()]

		sb.WriteByte(character)
	}

	return sb.String()
}
