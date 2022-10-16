package file

import "testing"

func TestSha256part(t *testing.T) {
	data := []byte("Hello")
	want := "185f8db32271"
	if have := Sha256part(data); want != have {
		t.Errorf("ERROR: got %s hash for string \"%s\", but want %s", have, string(data), want)
	}
}

func TestSha512part(t *testing.T) {
	data := []byte("Hello")
	want := "3615f80c9d29"
	if have := Sha512part(data); want != have {
		t.Errorf("ERROR: got %s hash for string \"%s\", but want %s", have, string(data), want)
	}
}

func TestXxhash64part(t *testing.T) {
	data := []byte("Hello")
	want := "0a75a91375b2"
	if have := Xxhash64part(data); want != have {
		t.Errorf("ERROR: got %s hash for string \"%s\", but want %s", have, string(data), want)
	}
}
