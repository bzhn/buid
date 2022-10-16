package file

import (
	"strconv"
)

// sizeMod returns (size mod 16),
// where size is a size of file in bytes
// The output will be in the hex number system
func SizeMod(filesize int64) string {
	return strconv.FormatInt(filesize%16, 16)
}

// sizeAvg returns hex digit that stands
// for amount of symbols after first
// symbol of the file size in bytes
func SizeApprox(filesize int64) string {
	l := digits(filesize)
	if l > 16 {
		return "f"
	}
	return strconv.FormatInt(l-1, 16)
}

func digits(number int64) (count int64) {
	if number == 0 {
		return 1
	}
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}
