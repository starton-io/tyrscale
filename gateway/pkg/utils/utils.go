package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GetSHA256Checksum(filePath string, checksum string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return false, err
	}
	computedChecksum := hex.EncodeToString(h.Sum(nil))
	return computedChecksum == checksum, nil
}
