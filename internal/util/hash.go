package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPIN(pin string) string {
	hash := sha256.Sum256([]byte(pin))
	return hex.EncodeToString(hash[:])
}
