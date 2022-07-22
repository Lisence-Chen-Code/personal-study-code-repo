package test

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"strconv"
	"testing"
)

func compressLZW(testStr string) []int {
	code := 256
	dictionary := make(map[string]int)
	for i := 0; i < 256; i++ {
		dictionary[strconv.Itoa(i)] = i
	}

	currChar := ""
	result := make([]int, 0)
	for _, c := range []byte(testStr) {
		phrase := currChar + string(c)
		if _, isTrue := dictionary[phrase]; isTrue {
			currChar = phrase
		} else {
			result = append(result, dictionary[currChar])
			dictionary[phrase] = code
			code++
			currChar = string(c)
		}
	}
	if currChar != "" {
		result = append(result, dictionary[currChar])
	}
	return result
}

func decompressLZW(compressed []int) string {
	code := 256
	dictionary := make(map[int]string)
	for i := 0; i < 256; i++ {
		dictionary[i] = strconv.Itoa(i)
	}

	currChar := strconv.Itoa(compressed[0])
	result := currChar
	for _, element := range compressed[1:] {
		var word string
		if x, ok := dictionary[element]; ok {
			word = x
		} else if element == code {
			word = currChar + currChar[:1]
		} else {
			panic(fmt.Sprintf("Bad compressed element: %d", element))
		}

		result += word

		dictionary[code] = currChar + word[:1]
		code++

		currChar = word
	}
	return result
}

func TestStt(t *testing.T) {

	var in bytes.Buffer
	b := []byte(`xiorui.cc`)
	w := zlib.NewWriter(&in)
	w.Write(b)
	w.Close()

	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.String())
}
