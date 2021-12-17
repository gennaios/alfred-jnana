package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// FileHash create sha256 file hash for later comparison
func FileHash(file string) (string, error) {
	if _, err := os.Stat(file); err != nil {
		return "", err
	}
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString, err
}
