package cryptopkg

import (
	"crypto/sha256"
	"crypto/x509"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashFromPassword generates password hash
func GenerateHashFromPassword(pass string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(h)
}

// CompareHashAndPassword compares pass with hash
func CompareHashAndPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func AppendCertsFromPEM(pemCerts []byte) (*x509.CertPool, bool) {
	cp := x509.NewCertPool()
	b := cp.AppendCertsFromPEM(pemCerts)
	return cp, b
}

func SHA256Sum(data []byte) [32]byte {
	return sha256.Sum256(data)
}
