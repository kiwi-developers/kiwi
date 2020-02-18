package utils

import "reflect"

// Capitalize: change first character to upper
// 改变字符串首字母为大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// return true if item in slice  check:=ItemInSlice([]slice, "value")
// 检查slice中是否存在某个元素
func ItemInSlice(s []interface{}, val interface{}) bool {
	targetValue := reflect.ValueOf(s)
	switch reflect.TypeOf(s).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == val {
				return true
			}
		}
	}
	return false
}
