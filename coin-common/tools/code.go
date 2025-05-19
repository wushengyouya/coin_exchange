package tools

import (
	"fmt"
	"math/rand"
)

// 随机生成四位验证码
func Rand4Num() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn += 1000
	}
	return fmt.Sprintf("%d", intn)
}
