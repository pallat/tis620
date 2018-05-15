package tis620

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

func ToUTF8(s string) (string, error) {
	tis620Reader := bytes.NewBufferString(s)

	reader, err := charset.NewReaderLabel("TIS-620", tis620Reader)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
