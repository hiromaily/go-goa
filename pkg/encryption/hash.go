package encryption

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

//-----------------------------------------------------------------------------
// HASH
//-----------------------------------------------------------------------------

// HashMD5 hashes by MD5 message-digest algorithm
func HashMD5(target string) (string, error) {
	if target == "" {
		return "", nil
	}

	h := md5.New()
	if _, err := io.WriteString(h, target); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// HashSHA1 hashes by SHA1
func HashSHA1(target string) (string, error) {
	if target == "" {
		return "", nil
	}

	h := sha1.New()
	if _, err := io.WriteString(h, target); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// HashSHA256 hashes by SHA256
func HashSHA256(target string) (string, error) {
	if target == "" {
		return "", nil
	}

	h := sha256.New()
	if _, err := io.WriteString(h, target); err != nil {
		return "", nil
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//-----------------------------------------------------------------------------
// Interface
//-----------------------------------------------------------------------------

// Hasher interface
type Hasher interface {
	Hash(target string) string
}

//-----------------------------------------------------------------------------
// MD5 Plus
//-----------------------------------------------------------------------------

// MD5 interface
type MD5 interface {
	Hash(target string) (string, error)
	HashWith(target, additional string) (string, error)
}

type md5Hash struct {
	salt1 string
	salt2 string
}

// NewMD5 returns MD5 interface
func NewMD5(salt1, salt2 string) MD5 {
	return &md5Hash{
		salt1: salt1,
		salt2: salt2,
	}
}

// Hash hashes target
func (m *md5Hash) Hash(target string) (string, error) {
	return m.hash(target, "")
}

// HashWith hashes target with additional salt
func (m *md5Hash) HashWith(target, additional string) (string, error) {
	return m.hash(target, additional)
}

func (m *md5Hash) hash(target, additional string) (string, error) {
	if target == "" {
		return "", nil
	}

	h := md5.New()
	if _, err := io.WriteString(h, target); err != nil {
		return "", err
	}
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	if _, err := io.WriteString(h, m.salt1); err != nil {
		return "", err
	}
	if _, err := io.WriteString(h, m.salt2); err != nil {
		return "", err
	}
	if additional != "" {
		if _, err := io.WriteString(h, additional); err != nil {
			return "", err
		}
	}
	if _, err := io.WriteString(h, hashed); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//-----------------------------------------------------------------------------
// Scrypt
//-----------------------------------------------------------------------------

// Scrypt interface
type Scrypt interface {
	Hash(target string) string
}

type scryptHash struct {
	salt string
}

// NewScrypt returns Scrypt interface
func NewScrypt(salt string) Scrypt {
	return &scryptHash{
		salt: salt,
	}
}

func (s *scryptHash) Hash(target string) string {
	if target == "" {
		return ""
	}

	key, _ := scrypt.Key([]byte(target), []byte(s.salt), 16384, 8, 1, 32)
	result := base64.StdEncoding.EncodeToString(key)
	return result
}
