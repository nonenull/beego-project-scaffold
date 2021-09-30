package sign

import (
	"fmt"
	"testing"
)

var s *Signer

func init() {
	s = NewSignerWithDefaultOption()
}

func TestGenerateSalt(t *testing.T) {
	salt := s.generateSalt()
	fmt.Println("生成 salt ", salt)
}

func TestSign(t *testing.T) {
	a := s.Sign("fucker")
	fmt.Println("生成 加密文本:  ", a)
}

func TestUnSign(t *testing.T) {
	a := s.Sign("fucker")
	fmt.Println("生成 加密文本:  ", a)

	b := s.Verify("fucker", a)
	fmt.Println("验签结果:  ", b)
}
