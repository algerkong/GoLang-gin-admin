package utils

import (
	"fmt"
	"strings"
	"time"
)

// 获取当前时间并格式化
func GetCurrentTime() string {
    return time.Now().Format("20060102150405")
}

// 当前时间加上10位哈希值
func GetCurrentTimeHash() string {
    return GetCurrentTime() + GetHash(10)
}

// 获取哈希值
func GetHash(length int) string {
    hash := time.Now().UnixNano()
    return strings.Replace(fmt.Sprintf("%d", hash), " ", "", -1)[:length]
}