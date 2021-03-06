// Package crypto handles cryptographic related operations require to
// generate OAuth1.0a header.
package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// Sha256 generates the SHA256 hash of the provided data
func Sha256(data []byte) []byte {

	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

// Sign signs the given signing data by using the RSA PrivateKey.
func Sign(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	digest := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest[:])
}
