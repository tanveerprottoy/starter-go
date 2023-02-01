package crypto

import (
	"crypto/x509"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashFromPassword generates password hash
func GenerateHashFromPassword(p string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(h)
}

// CompareHashAndPassword compares pass with hash
func CompareHashAndPassword(h string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}

func AppendCertsFromPEM(pemCerts []byte) (*x509.CertPool, bool) {
	cp := x509.NewCertPool()
	b := cp.AppendCertsFromPEM(pemCerts)
	return cp, b
}
