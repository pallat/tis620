package tis620

import (
	"bytes"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func Tis620ToUtf8(s string) (string, error) {
	tis620Reader := bytes.NewReader([]byte(s))

	e, _ := charset.Lookup("ISO-8859-1")
	isoReader := transform.NewReader(tis620Reader, e.NewEncoder())

	utfReader, err := charset.NewReaderLabel("TIS-620", isoReader)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(utfReader)

	return buf.String(), nil
}
