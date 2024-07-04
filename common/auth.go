package common

import (
	"errors"
	"fmt"
)

const accessSecret string = "123456"

func Auth(token string) (int, error) {
	Claims, err := ParseToken(token, accessSecret)
	if err != nil {
		// fmt.Println(err)
		return -1, errors.New("解析token失败")
	}

	return Claims.UserID, nil
}

func GetToken(userid int) (string, error) {
	token, err := GenToken(JwtPayLoad{userid}, accessSecret, 6)
	if err != nil {
		fmt.Println("获取token失败")
		return "", errors.New("获取token失败")
	}
	return token, nil
}
