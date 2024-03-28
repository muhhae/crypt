package crypt

import "encoding/base64"

func SimpleShiftEncrypt(content string, key string) (string, error) {
	byteC := []byte(content)
	res := make([]byte, len(byteC))
	k := 0
	for i := range byteC {
		res[i] = byteC[i] << (int(key[k]) % 2)
		k++
		if k >= len(key) {
			k = 0
		}
	}
	resStr := base64.StdEncoding.EncodeToString(res)
	return resStr, nil
}

func SimpleShiftDecrypt(code string, key string) (string, error) {
	byteC, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	res := make([]byte, len(byteC))
	k := 0
	for i := range byteC {
		res[i] = byteC[i] >> (int(key[k]) % 2)
		k++
		if k >= len(key) {
			k = 0
		}
	}
	return string(res), nil
}
