package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func powSHA256(data string, difficulty int) ([]byte, int) {
	const (
		maxNonce = 100000000000
	)

	checker := initChecker(difficulty)

	for nonce := 0; nonce < maxNonce; nonce++ {
		payload := fmt.Sprintf("%d:%s", nonce, data)
		hash := sha256.Sum256([]byte(payload))
		if checker(hash[:]) {
			return hash[:], nonce
		}
	}

	return nil, 0
}

func initChecker(difficulty int) func([]byte) bool {
	prefix := make([]byte, difficulty)
	bufferIdx := 0
	for n := 0; n < difficulty; n++ {
		bufferIdx += copy(prefix[bufferIdx:], []byte{0x00})
	}

	return func(hash []byte) bool {
		return bytes.HasPrefix(hash, prefix)
	}
}

func checkProof(data, proof []byte, nonce int) bool {
	payload := fmt.Sprintf("%d:%s", nonce, data)
	hash := sha256.Sum256([]byte(payload))

	return bytes.Equal(hash[:], proof)
}
