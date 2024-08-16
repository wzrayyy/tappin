package config

import "encoding/base64"

func (c *DailyCipher) Decode() (string, error) {
	enc, err := base64.StdEncoding.DecodeString(c.Cipher[:3] + c.Cipher[4:])

	if err != nil {
		return "", err
	}
	return string(enc), nil
}
