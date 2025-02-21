package main

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net"
	"net/http"
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

func getClientIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return strings.Split(forwarded, ",")[0] // First IP in the list
	}
	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // As last resort
	}
	return ip
}
