package random

import (
	"crypto/rand"
)

func HexChar() int64 {
	b := make([]byte, 1)
	rand.Read(b)
	return int64(b[0] % 4)
}
