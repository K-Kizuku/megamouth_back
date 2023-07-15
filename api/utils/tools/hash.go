package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(s string) (h string) {
	b := getSHA256Binary(s)
	h = hex.EncodeToString(b)
	return
}

func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
