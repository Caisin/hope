package str

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

//LeftUpper 首字母转大写
func LeftUpper(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

//LeftLower 首字母转小写
func LeftLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
}

// Camel2Case 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	return Camel2Split(name, "_")
}

// Camel2Split 驼峰式写法转为指定分隔符写法
func Camel2Split(name, split string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append(split)
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	return Split2Camel(name, "_")
}

func Split2Camel(name, split string) string {
	name = strings.Replace(name, split, " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
