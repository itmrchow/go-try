package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

)

func main() {
	start := time.Now()

	PassWordHashingHandler()

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
	return cost
}

func PassWordHashingHandler() {
	password := "password"
	hash, _ := HashPassword(password) // 为了简单忽略错误处理

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:", match)

	cost := GetHashingCost([]byte(hash))
	fmt.Println("Cost:", cost)

}
