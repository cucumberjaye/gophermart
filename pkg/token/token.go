package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

func GenerateToken(id string, signingKey []byte) (string, error) {
	key := md5.Sum(signingKey)

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	enc := aesgcm.Seal(nil, key[:12], []byte(id), nil)

	return hex.EncodeToString(enc), nil
}

func CheckToken(token string, signingKey []byte) (string, error) {
	data, err := hex.DecodeString(token)
	if err != nil {
		return "", err
	}

	key := md5.Sum(signingKey)

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	dec, err := aesgcm.Open(nil, key[:12], data, nil)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}
