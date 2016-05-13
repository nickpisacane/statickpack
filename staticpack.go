package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const elmsPerLine = 13

// StaticPack static pack
func StaticPack(r io.Reader, name string) (b []byte, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	name = capitalizeFirst(name)
	b = []byte(fmt.Sprintf("var %s []byte = []byte{\n%s}", name,
    createByteBody(data)))
	return
}

func createByteBody(data []byte) string {
	buf := &bytes.Buffer{}
	perLineCount := 1

	for i, b := range data {
		if perLineCount == 1 {
			buf.WriteString("\t")
		}
		buf.WriteString(fmt.Sprintf("%s, ", createHexStr(b)))
		if perLineCount == elmsPerLine || i == len(data)-1 {
			buf.WriteString("\n")
			perLineCount = 1
		} else {
			perLineCount++
		}
	}

	return string(buf.Bytes())
}

func createHexStr(b byte) string {
	hx := fmt.Sprintf("%x", b)
	if len(hx) == 1 {
		hx = "0" + hx
	}
	return "0x" + hx
}

func capitalizeFirst(s string) string {
	if len(s) < 1 {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
