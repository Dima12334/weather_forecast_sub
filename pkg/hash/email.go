package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type EmailHasher interface {
	GenerateEmailHash(email string) string
}

type SHA256Hasher struct{}

func NewSHA256Hasher() *SHA256Hasher {
	return &SHA256Hasher{}
}

func (h *SHA256Hasher) GenerateEmailHash(email string) string {
	hash := sha256.Sum256([]byte(email))
	return hex.EncodeToString(hash[:])
}
