package str

import (
	"math/rand"
	"strings"
)

func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	for _, r := range []rune(str) {
		if r != ' ' {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsAllBlank(str ...string) bool {
	for _, s := range str {
		if IsNotBlank(s) {
			return false
		}
	}
	return true
}

func ContainsAny(str string, subs ...string) bool {
	for _, s := range subs {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

func IsAllNotBlank(str ...string) bool {
	for _, s := range str {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

func IsNotAllBlank(str ...string) bool {
	return IsAllBlank(str...)
}

func RandNum(size int) string {
	if size == 0 {
		return ""
	}
	bf := NewBuffer()
	for i := 0; i < size; i++ {
		bf.Append(rand.Intn(9))
	}
	return bf.String()
}

func IsAllEmpty(str ...string) bool {
	for _, s := range str {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

func IsAllNotEmpty(str ...string) bool {
	for _, s := range str {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}

//Rep 将字符串之间字符替换
func Rep(str, start, end, rep string) string {
	strLen := len(str)
	if strLen == 0 {
		return str
	}
	startIdx := strings.Index(str, start) + 1
	endIdx := strings.Index(str, end)
	ret := str[0:startIdx] + rep + str[endIdx:strLen]
	return ret
}
