package utils

import (
	"fmt"
)

func Logger(uid interface{}, level interface{}, message ...interface{}) {
	vs := ""
	for _, _ = range message {
		vs += "%+v"
	}
	fmt.Printf("[%s][%s][%s]: %s\n", GetNow(), uid, level, fmt.Sprintf(vs, message...))
}
