package utils

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 会在第一次使用该包时被调用，并且只会被调用一次
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 随机生成一个整数，范围是[min, max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// 随机生成包含n个字符串的字符串
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// 随机所有者
func RandomOwner() string {
	return RandomString(6)
}

// 随机账户余额
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// 随机货币类型
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP", "JPY", "CNY"}
	return currencies[rand.Intn(len(currencies))]
}
