package encrypt

// string to string encrypt/decrypt, used to protect password in PlainText like json/toml config file.

import (
	"github.com/jbenet/go-base58"
	"github.com/shawnwyckoff/gpkg/crypto/hash"
)

func SonnefesEncrypt(plainText, key string) (string, error) {
	keySHA256 := hash.GetSHA256([]byte(key))
	cipherAES256, err := SymmetricEncrypt(ALG_AES_256_CBC, []byte(plainText), keySHA256)
	if err != nil {
		return "", err
	}
	return base58.Encode(cipherAES256), nil
}

func SonnefesDecrypt(cipher string, key string) (string, error) {
	keySHA256 := hash.GetSHA256([]byte(key))
	bin := base58.Decode(cipher)
	plainText, err := SymmetricDecrypt(ALG_AES_256_CBC, bin, []byte(keySHA256))
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}
