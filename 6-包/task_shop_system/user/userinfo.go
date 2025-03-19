// @FileName : userinfo.go
// @Time : 2025/3/18 下午6:50
// @Author : luobozi

package user

import (
	"encoding/json"
	"os"
	"unicode/utf8"
)

// 此处开始编写代码
func User_regriter(userinfo map[string]string) bool {

	if utf8.RuneCountInString(userinfo["name"]) <= 8 && utf8.RuneCountInString(userinfo["passwd"]) >= 10 {
		json_data, e := json.Marshal(userinfo)
		if e != nil {
			return false
		}
		err := os.WriteFile("userinfo.json", json_data, 0644)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
