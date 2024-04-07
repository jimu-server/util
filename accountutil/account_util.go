package accountutil

import (
	"crypto/sha512"
	"encoding/base64"
)

// Password VerifyPasswd 验证密码是否正确
// @param password 输入的密码
func Password(password string) (base64Hash string) {
	salt := password[:]
	hash := sha512.Sum512([]byte(salt + password))
	base64Hash = base64.StdEncoding.EncodeToString(hash[:])
	return
}

// VerifyPasswd 验证密码是否正确
// @param source 数据库密码
// @param passwd 用户输入的密码
func VerifyPasswd(source, passwd string) bool {
	hash := sha512.Sum512([]byte(passwd[:] + passwd))
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])
	return base64Hash == source
}
