package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buf []byte) (int, error) {
	bufSize := len(buf)
	for i := 0; i < bufSize; i++ {
		buf[i] = 'A'
	}
	return bufSize, nil
}

func main() {
	reader.Validate(MyReader{})
}
