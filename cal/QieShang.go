package cal

// QieShang 返回向上取整的一个百位整数
// 例如：1440 返回 1500
func QieShang(Nums int) int {
	if Nums%100 == 0 {
		return Nums
	} else {
		Nums = Nums + 100 - (Nums % 100)
		return Nums
	}
}
