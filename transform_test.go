package tis620

import (
	"testing"
)

func TestTis620ToUtf8(t *testing.T) {
	s := "¢Íâ·É¤èÐ ÃÐººäÁèÊÒÁÒÃ¶ãËéºÃÔ¡ÒÃä´éã¹¢³Ð¹Õé ¡ÃØ³Ò·ÓÃÒÂ¡ÒÃãËÁèÀÒÂËÅÑ§¤èÐ"
	x := "ขอโทษค่ะ ระบบไม่สามารถให้บริการได้ในขณะนี้ กรุณาทำรายการใหม่ภายหลังค่ะ"
	r, _ := Tis620ToUtf8(s)

	if r != x {
		t.Errorf("%s\nis exptected but got\n%s\n", x, r)
	}
}
