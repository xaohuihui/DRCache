package utils

import (
	"fmt"
	"time"
)

// GetNowFormatTodayTime 获取当天年月日函数
func GetNowFormatTodayTime() string {
	now := time.Now()
	dateStr  := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
	return dateStr
}

