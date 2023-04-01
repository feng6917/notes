package hex

import (
	"encoding/hex"
	"testing"
)

func Test_AesCBC(t *testing.T) {
	origin := "qwews"
	key := "zxcvbnmlkjhgfdsa"
	r := "697c3f2624036b2c056260ac542c414c"
	rbt, err := hex.DecodeString(r)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Run("Encrypt", func(t *testing.T) {
		encrypted, err := AesEncryptCBC([]byte(origin), []byte(key))
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		ss := hex.EncodeToString(encrypted)
		if ss != r {
			t.Fail()
		}
	})
	t.Run("Decrypt", func(t *testing.T) {
		decrypted, err := AesDecryptCBC(rbt, []byte(key))
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		if string(decrypted) != origin {
			t.Fail()
		}
	})
}
