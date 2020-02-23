package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

type HashType string

const (
	HashTypeSHA1       HashType = "sha1"
	HashTypeSHA256     HashType = "sha256"
	HashTypeSHA512     HashType = "sha512"
	HashTypeSHA512_384 HashType = "sha512_384"
	HashTypeMD5        HashType = "md5"
)

// GetSHA256 returns a SHA256 hash of a byte array
func GetSHA256(in []byte) []byte {
	digest := sha256.New()
	digest.Write(in)
	return digest.Sum(nil)
}

// GetSHA512 returns a SHA512 hash of a byte array
func GetSHA512(input []byte) []byte {
	sha := sha512.New()
	sha.Write(input)
	return sha.Sum(nil)
}

func GetSHA1(in []byte) []byte {
	digest := sha1.New()
	digest.Write(in)
	return digest.Sum(nil)
}

// GetMD5 returns a MD5 hash of a byte array
func GetMD5(in []byte) []byte {
	digest := md5.New()
	digest.Write(in)
	return digest.Sum(nil)
}

func GetHex(in []byte) string {
	return hex.EncodeToString(in)
}

func GetB64(in []byte) string {
	return base64.RawStdEncoding.EncodeToString(in)
}

// HexEncodeToString takes in a hexadecimal byte array and returns a string
func HexEncodeToString(input []byte) string {
	return hex.EncodeToString(input)
}

// GetHMAC returns a keyed-hash message authentication code using the desired
// hashtype
func GetHMAC(hashType HashType, input, key []byte) []byte {
	var hash func() hash.Hash

	switch hashType {
	case HashTypeSHA1:
		{
			hash = sha1.New
		}
	case HashTypeSHA256:
		{
			hash = sha256.New
		}
	case HashTypeSHA512:
		{
			hash = sha512.New
		}
	case HashTypeSHA512_384:
		{
			hash = sha512.New384
		}
	case HashTypeMD5:
		{
			hash = md5.New
		}
	}

	hmac := hmac.New(hash, []byte(key))
	hmac.Write(input)
	return hmac.Sum(nil)
}
