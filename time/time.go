package time

import (
	"fmt"

	"github.com/google/uuid"
)

func Timestamp() string {
	h, m, l := threeParts()
	return fmt.Sprintf("%x-%x-%x", h, m, l)
}

func threeParts() (time_low uint32, time_mid uint16, time_hi uint16) {
	now, _, _ := uuid.GetTime()
	time_low = uint32(now & 0xffffffff)
	time_mid = uint16((now >> 32) & 0xffff)
	time_hi = uint16((now >> 48) & 0x0fff)
	time_hi |= 0xb000 // Version b

	return
}
