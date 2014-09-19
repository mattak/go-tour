package main

import (
	"io"
	"os"
	"strings"
)

const ROT_NUMBER = 13

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	if err != nil {
		return n, err
	}

	var ALPHABET_NUMBER byte = 26

	for i := 0; i < n; i++ {
		var OFFSET byte = 0

		if p[i] >= 65 && p[i] <= 90 {
			// A-Z: 65 - 90
			OFFSET = 65
			p[i] = ((p[i]-OFFSET)%ALPHABET_NUMBER+ROT_NUMBER)%ALPHABET_NUMBER + OFFSET
		} else if p[i] >= 97 && p[i] <= 122 {
			// a-z: 97 - 122
			OFFSET = 97
			p[i] = ((p[i]-OFFSET)%ALPHABET_NUMBER+ROT_NUMBER)%ALPHABET_NUMBER + OFFSET
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
