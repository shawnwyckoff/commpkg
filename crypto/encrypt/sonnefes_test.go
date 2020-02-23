package encrypt

import "testing"

func TestSonnefesEncrypt(t *testing.T) {
	in := "hello 你好 こんにちは"
	key := "this is a very complex key!"
	cipher, err := SonnefesEncrypt(in, key)
	if err != nil {
		t.Error(err)
		return
	}
	dec, err := SonnefesDecrypt(cipher, key)
	if err != nil {
		t.Error(err)
		return
	}
	if dec != in {
		t.Errorf("SonnefesEncrypt/SonnefesDecrypt error")
		return
	}
}
