package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bzhn/buid/random"

	file "github.com/bzhn/buid/file"
	tm "github.com/bzhn/buid/time"
)

const (
	variationF = "f" // stands for files

	// hash to use
	sha256 = 0
	sha512 = 1
	xxh64  = 2

	SHA256HASHnum = 0 // 00xx
	SHA512HASHnum = 4 // 01xx
	XXH64HASHnum  = 8 // 10xx
)

var hashnum uint
var pathtofile string

func init() {
	flag.UintVar(&hashnum, "hash", 0, "Hash to use\n0 for SHA256\n1 for SHA512\n2 for XXH64")
	flag.StringVar(&pathtofile, "path", `C:\Program Files\7-Zip\7z.exe`, "Path to file")
	flag.Parse()
}

func main() {

	uuid, err := UUID(uint8(hashnum), pathtofile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", uuid)
}

func UUID(hashToUse uint8, pathToFile string) (string, error) {
	var variation, randomhex, hashPart string

	variation = variationF

	// Check that hashToUse is between 0 and 2
	if hashToUse < 0 || hashToUse > 2 {
		return "", errors.New("Number of the hash to use is not in the range between 0 and 2")
	}
	// TODO: Check that file exists and accessible

	f, err := os.Open(pathToFile)
	defer f.Close()
	fstat, err := f.Stat()
	if err != nil {
		return "", err
	}
	fileSize := fstat.Size()
	fileContent, err := os.ReadFile(pathToFile)
	if err != nil {
		return "", err
	}
	switch hashToUse {
	case sha256:
		randomhex = fmt.Sprintf("%x", SHA256HASHnum+random.HexChar())
		hashPart = file.Sha256part(fileContent)
	case sha512:
		randomhex = fmt.Sprintf("%x", SHA512HASHnum+random.HexChar())
		hashPart = file.Sha512part(fileContent)
	case xxh64:
		randomhex = fmt.Sprintf("%x", XXH64HASHnum+random.HexChar())
		hashPart = file.Xxhash64part(fileContent)
	default:
		return "", errors.New("wrong number of hash to use")
	}

	return fmt.Sprintf("%s-%s%s%s%s-%s", tm.Timestamp(), variation, file.SizeApprox(fileSize), file.SizeMod(fileSize), randomhex, hashPart), nil
}
