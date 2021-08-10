package utils

import (
	"fmt"
)

func Logger(uid string, level string, message ...interface{}) {
	fmt.Printf("[%s][%s][%s]: %+v\n", GetNow(), uid, level, message)
}
