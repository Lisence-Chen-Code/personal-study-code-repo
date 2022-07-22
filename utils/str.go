package govalidator

import (
	"strings"
	"unsafe"
)

// ToSnake string, XxYy to xx_yy, X_Y to x_y
func ToSnake(s string) string {
	num := len(s)
	need := false // need determine if it's necessary to add a '_'

	snake := make([]byte, 0, len(s)*2)
	for i := 0; i < num; i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
			if need {
				snake = append(snake, '_')
			}
		} else {
			// if previous is '_' or ' ',
			// there is no need to add extra '_' before
			need = c != '_' && c != ' '
		}

		snake = append(snake, c)
	}

	return string(snake)
}

// ToCamel string, xx_yy to XxYy, xx__yy to Xx_Yy
// xx _yy to Xx Yy, the rule is that a lower case letter
// after '_' will combine to a upper case letter
func ToCamel(s string) string {
	num := len(s)
	need := true

	var prev byte = ' '
	camel := make([]byte, 0, len(s))
	for i := 0; i < num; i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			if need {
				c = c - 'a' + 'A'
				need = false
			}
		} else {
			if prev == '_' {
				camel = append(camel, '_')
			}
			need = c == '_' || c == ' '
			if c == '_' {
				prev = '_'
				continue
			}
		}

		prev = c
		camel = append(camel, c)
	}

	return string(camel)
}

// TrimQuote trim quote for string, return error if quote don't match
func TrimQuote(str string) (string, bool) {
	str = strings.TrimSpace(str)
	l := len(str)
	if l == 0 {
		return "", true
	}

	if s, e := str[0], str[l-1]; s == '\'' || s == '"' || s == '`' || e == '\'' || e == '"' || e == '`' {
		if l != 1 && s == e {
			str = str[1 : l-1]
		} else {
			return "", false
		}
	}

	return str, true
}

// String2bytes zero-copy
func String2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
