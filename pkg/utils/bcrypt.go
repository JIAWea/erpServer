package utils

import "golang.org/x/crypto/bcrypt"

const (
	DefaultPasswordRegex = "^[a-zA-Z0-9]{9,20}$"
)

// GenPasswd 密码加密 使用自适应hash算法, 不可逆
func GenPasswd(passwd string) string {
	hashPasswd, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(hashPasswd)
}

// ComparePasswd 通过比较两个字符串hash判断是否出自同一个明文
// hashPasswd 需要对比的密文
// passwd 明文
func ComparePasswd(hashPasswd string, passwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(passwd)); err != nil {
		return err
	}
	return nil
}
