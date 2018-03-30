// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"math/big"
)

// GetRandomString generate random string by specify chars.
func GetRandomString(n int) (string, error) {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	buffer := make([]byte, n)
	max := big.NewInt(int64(len(alphanum)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return "", err
		}

		buffer[i] = alphanum[index]
	}

	return string(buffer), nil
}

// NewLfsJwtSecret generate a new value intended to be used by LFS_JWT_SECRET.
func NewLfsJwtSecret() (string, error) {
	JWTSecretBytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, JWTSecretBytes)
	if err != nil {
		return "", err
	}

	JWTSecretBase64 := base64.RawURLEncoding.EncodeToString(JWTSecretBytes)
	return JWTSecretBase64, nil
}

// NewSecretKey generate a new value intended to be used by SECRET_KEY.
func NewSecretKey() (string, error) {
	secretKey, err := GetRandomString(64)
	if err != nil {
		return "", err
	}

	return secretKey, nil
}

func randomInt(max *big.Int) (int, error) {
	rand, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return int(rand.Int64()), nil
}
