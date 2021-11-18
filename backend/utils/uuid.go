package utils

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
)

func GenerateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func MakeUUID() (string, error) {
	return uuid.MakeUUID()
}
