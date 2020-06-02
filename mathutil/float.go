package mathutil

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// 四舍五入基准
	baseTimes float64 = 5

	// 两位小数（货币）
	baseDecimal int = 2
)

// 将float64按精度四舍五入保留成浮点数
// f浮点数据
// 保留小数位数
// 转换结果
func CF64ToDecimal(f float64, r int) float64 {

	// 计算小数位数
	var m float64 = baseTimes / math.Pow10(r+1)

	return float64(math.Floor((f+m)*10) / 10)
}

// 保留2位小数
// f浮点数据
// 转换结果
func CF64ToTwoDecimal(f float64) float64 {
	return CF64ToDecimal(f, baseDecimal)
}

// 根据精度float64转换成int类型
// f 输入数据
// 转换结果
func TryParseCF64ToInt(f float64) (int, error) {

	s := fmt.Sprintf("%.0f", f)

	i, err := strconv.Atoi(s)

	return i, err
}
