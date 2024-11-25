package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func VerifySignature(message, pubKey string, signature []byte) error {
	// Parse the PKCS#8 public key
	block, _ := pem.Decode([]byte(pubKey))
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// Assert as an RSA public key
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("fqf")
	}

	// Calculate the SHA256 hash value
	hashed := sha256.Sum256([]byte(message))

	// Verify the signature
	return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signature)
}

func SHA256WithRSA2048(message string, privateKeyPEM string) (string, error) {
	messageByte := []byte(message)
	// Decode private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("ras private key err")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Assert as an RSA public key
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not the ras private key")
	}

	hash := sha256.New()
	hash.Write(messageByte)
	hashed := hash.Sum(nil)

	// Signature using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}
