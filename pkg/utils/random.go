package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomUUID(num uint32) string {
	day := time.Now().Format("20060102")
	return fmt.Sprintf("%s%s", day, RandomString(int(num)))
}

func GenUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
