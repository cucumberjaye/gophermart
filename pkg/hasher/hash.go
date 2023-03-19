package hasher

import (
	"crypto/sha256"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
