package utils

import "time"

var cstSh, _ = time.LoadLocation("Asia/Shanghai")

func GetNow() string {
	return time.Now().In(cstSh).Format("2006/01/02 15:04:05")
}
