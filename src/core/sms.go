package core

import (
	"fmt"
)

func SendSMS(phoneNumber string, code int64) {
	fmt.Printf("response is %#v\n", code)
}
