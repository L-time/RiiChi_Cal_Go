package utils

import "strconv"

var Pai = map[string]bool{"m": true, "p": true, "s": true, "z": true}

// Paras 将字符串初步处理为牌型，同时提取赤宝牌
func Paras(text string) ([]string, int) {
	var temp []string
	var RedDora int
	for _, s := range text {
		char := string(s)
		_, err := strconv.Atoi(char)
		if err == nil {
			if char == "0" {
				char = "5"
				RedDora++
			}
			temp = append(temp, char)
		} else {
			if _, ok := Pai[char]; ok {
				for k, _ := range temp {
					if _, err = strconv.Atoi(temp[k]); err == nil {
						temp[k] = temp[k] + char
					}
				}
			}
		}
	}
	return temp, RedDora
}
