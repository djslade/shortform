package main

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

func generateURLID(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idSlice := make([]string, length)
	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		idSlice[i] = string(chars[charIndex.Int64()])
	}
	idString := strings.Join(idSlice, "")
	return idString, nil
}
