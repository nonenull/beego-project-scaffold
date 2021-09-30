package sign

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

var logger = logs.GetBeeLogger()

type Signer struct {
	keyLength  int
	iter       int
	saltLength int
}

func (s *Signer) generateSalt() []byte {
	b := make([]byte, s.saltLength)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		logger.Error("生成随机salt发生错误: %s", err)
		return []byte("")
	}
	return b
}

func (s *Signer) Sign(text string) string {
	salt := s.generateSalt()
	return s.sign(text, salt)
}

func (s *Signer) sign(text string, salt []byte) string {
	textBytes := []byte(text)
	logger.Debug("salt: %s", string(salt))
	logger.Debug("未加密文本: %s", []byte(text))
	cipherText := pbkdf2.Key(textBytes, salt, s.iter, s.keyLength, sha256.New)
	logger.Debug("加密文本内容: %s", string(cipherText))
	saltCipherText := append(salt, cipherText...)
	logger.Debug("加密文本加盐内容: %s", string(saltCipherText))
	return base64.StdEncoding.EncodeToString(saltCipherText)
}

func (s *Signer) Verify(text, encryptText string) bool {
	decodeText, err := base64.StdEncoding.DecodeString(encryptText)
	logger.Debug("解码文本加盐内容: %s", string(decodeText))
	if err != nil {
		logger.Error("加密文本 base64 解码失败: %s", err)
		return false
	}
	// 2、截取加密串 固定长度
	salt := decodeText[:s.saltLength]
	logger.Debug("解码 salt ===: %s", string(salt))

	reCryptText := s.sign(text, salt)
	logger.Debug("reCryptText ===: %s", string(reCryptText))
	return reCryptText == encryptText
}

func NewSignerWithDefaultOption() *Signer {
	return NewSigner(0, 0, 0)
}

func NewSigner(saltLength int, iter int, keyLength int) *Signer {
	if saltLength == 0 {
		saltLength = 16
	}
	if iter == 0 {
		iter = 1000
	}
	if keyLength == 0 {
		keyLength = 32
	}
	signer := &Signer{}
	signer.iter = iter
	signer.saltLength = saltLength
	signer.keyLength = keyLength

	return signer
}
