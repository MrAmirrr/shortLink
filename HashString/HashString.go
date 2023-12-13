package HashString

import (
	"crypto/sha256"
	"fmt"
)

func HashString(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
}
