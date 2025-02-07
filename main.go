package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const encoded = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

func encodeRot13(v byte) byte {
	switch {
	case v >= 'A' && v <= 'Z':
		return encoded[v-'A']
	case v >= 'a' && v <= 'z':
		return encoded[v-'a']
	default:
		return v
	}
}

func (rot *rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = rot.r.Read(buf)
	for i := 0; i < n; i++ {
		buf[i] = encodeRot13(buf[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
}
