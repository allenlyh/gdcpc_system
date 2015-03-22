package tools

import (
	"regexp"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
)

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

const _EXP_USERNAME = `^[A-Za-z0-9]$`

func CheckUserName(username string) (string, error) {
	length := len(username)
	if length < 3 || length > 20 {
		return username, errors.New("Username cannot be empty!")
	}
	exp := regexp.MustCompile(_EXP_USERNAME)
	if !exp.MatchString(username) {
		return username, errors.New("Invalid username!")
	}
	return username, nil
}

const _EXP_EMAIL = `^[A-Za-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2-4}$`

func CheckEmail(email string) (string, error) {
	length := len(email)
	if length == 0 {
		return email, errors.New("Email cannot be empty!")
	}
	exp := regexp.MustCompile(_EXP_EMAIL)
	if !exp.MatchString(email) {
		return email, errors.New("Invalid email")
	}
	return email, nil
}

func CheckNotEmpty(st string) (string, err) {
	length := len(st)
	if length == 0 {
		return st, errors.New("The string is empty!")
	}
	return st, nil
}

func GetMD5Pwd(password string) (string, err) {
	length := len(username)
	if length < 6 || length > 20 {
		return username, errors.New("Password length must be more than 5 and less then 21!")
	}
	data := []byte(password)
	return hex.EncodeToString(md5.Sum(data)), nil
}
