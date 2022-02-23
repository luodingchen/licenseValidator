package service

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"log"
)

func DecodePublicKeyString(publicKeyString string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyString)
	if err != nil {
		return nil, err
	}
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func Verify(publicKey *rsa.PublicKey, origData []byte, signatureString string) error {
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureString)
	if err != nil {
		return err
	}
	msgHash := sha256.New()
	msgHash.Write(origData)
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signatureBytes, nil)
	if err != nil {
		return err
	}
	return nil
}

// 公钥加密
func Encrypt(publicKey *rsa.PublicKey, origData []byte) (string, error) {
	partLen := publicKey.N.BitLen()/16 - 66
	chunks := Split(origData, partLen)
	var encrypts bytes.Buffer
	for _, chunk := range chunks {
		encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, chunk, []byte{})
		if err != nil {
			log.Println(len(chunk), publicKey.Size())
			return "", err
		}
		encrypts.Write(encrypted)
	}
	return base64.StdEncoding.EncodeToString(encrypts.Bytes()), nil
}

// 数据分块
func Split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
