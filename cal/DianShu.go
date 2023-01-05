package cal

import (
	"Richii_cal/enum/Xian"
	"Richii_cal/enum/Zhuang"
	"math"
)

// DianShuCal 点数计算的主函数，仅处理到累计役满
// Fan:番数；Fu:符数；isZhuang:是否为庄家。
// 返回点数值，不是应付点数！
// 符数应提前取整！
func DianShuCal(Fan int, Fu int, isZhuang bool) int {
	if isZhuang {
		if Fan > 12 {
			return Zhuang.LeiJiYiMa
		} else if Fan >= 11 {
			return Zhuang.SanBeiMan
		} else if Fan >= 8 {
			return Zhuang.BeiMan
		} else if Fan >= 6 {
			return Zhuang.TiaoMan
		} else {
			dedian := 6.0 * float64(Fu) * math.Pow(2, float64(Fan+2))
			if Fan < 6 && dedian > 12000 {
				return Zhuang.ManGuan
			}
			return QieShang(int(dedian))
		}
	} else {
		if Fan > 12 {
			return Xian.LeiJiYiMan
		} else if Fan >= 11 {
			return Xian.SanBeiMan
		} else if Fan >= 8 {
			return Xian.BeiMan
		} else if Fan >= 6 {
			return Xian.TiaoMan
		} else {
			dedian := 4.0 * float64(Fu) * math.Pow(2, float64(Fan+2))
			if Fan < 6 && dedian > 8000 {
				return Xian.ManGuan
			}
			return QieShang(int(dedian))
		}
	}
}
