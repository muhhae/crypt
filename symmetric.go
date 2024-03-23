package crypt

import (
	"encoding/base64"
)

func SimpleSymmetricEncrypt(content, key string) (string, error) {
	byteC := []byte(content)
	byteK := []byte(key)
	res := make([]byte, len(byteC))
	k := 0
	for i := range byteC {
		res[i] = byteC[i] ^ byteK[k]
		k++
		if k >= len(byteK) {
			k = 0
		}
	}
	resStr := base64.StdEncoding.EncodeToString(res)
	return resStr, nil
}

func SimpleSymmetricDecrypt(code, key string) (string, error) {
	byteC, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	byteK := []byte(key)
	res := make([]byte, len(byteC))
	k := 0
	for i := range byteC {
		res[i] = byteC[i] ^ byteK[k]
		k++
		if k >= len(byteK) {
			k = 0
		}
	}
	return string(res), nil
}

func ModifiedSimpleSymmetricEncrypt(content, key string) (string, error) {
	byteC := []byte(content)
	byteK := []byte(key)
	res := make([]byte, len(byteC))
	k := 0
	kI := len(byteK) - 1
	for i := range byteC {
		res[i] = byteC[i] ^ byteK[k] ^ byteK[kI]
		k++
		if k >= len(byteK) {
			k = 0
		}
		kI--
		if kI < 0 {
			kI = len(byteK) - 1
		}
	}
	resStr := base64.StdEncoding.EncodeToString(res)
	return resStr, nil
}

func ModifiedSimpleSymmetricDecrypt(code, key string) (string, error) {
	byteC, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	byteK := []byte(key)
	res := make([]byte, len(byteC))
	k := 0
	kI := len(byteK) - 1
	for i := range byteC {
		res[i] = byteC[i] ^ byteK[k] ^ byteK[kI]
		k++
		if k >= len(byteK) {
			k = 0
		}
		kI--
		if kI < 0 {
			kI = len(byteK) - 1
		}
	}
	return string(res), nil
}
