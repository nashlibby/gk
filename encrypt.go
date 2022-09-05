package gk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

// 生成密钥
func SecretMake(length int) string {
	var secret strings.Builder
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	for i := 0; i < length; i++ {
		secret.WriteRune(chars[rand.Intn(len(chars))])
	}
	return secret.String()
}

// 生成密码
func HashMake(password string) string {
	ret, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(ret)
}

// 校验密码
func HashCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if length-unpadding < 0 {
		return []byte("")
	}
	return src[:(length - unpadding)]
}

// open_ssl aes-128-cbc 解码
func DecryptData(str, key, iv string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(str)
	i := []byte(iv)
	k := []byte(key)
	cipherBlock, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	cipher.NewCBCDecrypter(cipherBlock, i).CryptBlocks(data, data)

	return string(PKCS5UnPadding(data)), nil
}

// md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// sha1加密
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
