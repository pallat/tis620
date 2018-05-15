package tis620

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func TestWriteTis620CharSet(t *testing.T) {
	s := "ขอโทษค่ะ ระบบไม่สามารถให้บริการได้ในขณะนี้ กรุณาทำรายการใหม่ภายหลังค่ะ"
	r := bytes.NewBufferString(s)
	nr := NewReaderEncoder("tis-620", r)

	b, err := ioutil.ReadAll(nr)
	if err != nil {
		return
	}

	ioutil.WriteFile("tis620.txt", b, 0666)
}

func TestTis620ToUtf8(t *testing.T) {
	b, _ := ioutil.ReadFile("tis620.txt")
	s := string(b)
	x := "ขอโทษค่ะ ระบบไม่สามารถให้บริการได้ในขณะนี้ กรุณาทำรายการใหม่ภายหลังค่ะ"
	r, _ := ToUTF8(s)

	if r != x {
		t.Errorf("%s\nis exptected but got\n%s\n", x, r)
	}
}

func NewReaderEncoder(label string, input io.Reader) (reader io.Reader) {
	e, _ := charset.Lookup(label)
	return transform.NewReader(input, e.NewEncoder())
}
