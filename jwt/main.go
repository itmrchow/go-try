package main

import (
	"github.com/golang-jwt/jwt"
)

var (
	t *jwt.Token
)

func main() {
	// H256
	JwtH256([]byte("secret")) // 移到env

}

func JwtH256(key []byte) string {
	// HS256 對稱加密演算法
	// 使用相同的金鑰來加密和解密資料

	t = jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(key)

	if err != nil {
		println(err.Error())
		return ""
	} else {
		println("token:" + s)
	}

	return s
}

// func JwtECDSA(key *ecdsa.PrivateKey) string {
// 	// 非對稱加密

// 	t = jwt.New(jwt.SigningMethodES256)
// 	s, _ := t.SignedString(key)
// 	return s
// }
