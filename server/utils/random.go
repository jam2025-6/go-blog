package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 生成一个指定长度的随机验证码
func GenerateVerficationCode(length int) string {
	// 创建一个随机数种子。
	// time.Now().UnixNano()当前时间纳秒
	// rand.NewSource 根据一个初始值（种子）生成一个随机数序列
	new := rand.NewSource(time.Now().UnixNano())
	// 使用上面创建的种子，生成一个随机数生成器 r
	r := rand.New(new)
	// 生成随机整数：r.Intn(10000) 得到一个 0-9999 之间的随机数，比如 123。
	return fmt.Sprintf("%0*d", length, r.Intn(int(math.Pow10(length))))
}
