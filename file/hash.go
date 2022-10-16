package file

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strconv"

	"github.com/cespare/xxhash"
)

// Sha256part returns sha256 sum of the bytes b
// cropped to 12 symbols
func Sha256part(b []byte) string {
	return sha256Sum(b)[:12]
}

// Sha512part returns sha512 sum of the bytes b
// cropped to 12 symbols
func Sha512part(b []byte) string {
	return sha512Sum(b)[:12]
}

// Xxhash64part returns xxhash64 sum of the bytes b
// cropped to 12 symbols
func Xxhash64part(b []byte) string {
	return xxhash64Sum(b)[:12]
}

func xxhash64Sum(b []byte) string {
	n := xxhash.Sum64(b)
	return fmt.Sprintf("%016s", strconv.FormatUint(n, 16))
}

func sha256Sum(b []byte) string {
	hasher := sha256.New()
	hasher.Write(b)

	n := hasher.Sum(nil)
	return fmt.Sprintf("%x", n)
}

func sha512Sum(b []byte) string {
	hasher := sha512.New()
	hasher.Write(b)

	n := hasher.Sum(nil)
	return fmt.Sprintf("%x", n)
}
