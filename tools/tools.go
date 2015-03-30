package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	. "gdcpc_system/models"
)

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

const _EXP_USERNAME = `^[A-Za-z0-9]+$`

func CheckUserName(username string) (string, error) {
	length := len(username)
	fmt.Println(length)
	if length < 3 || length > 20 {
		return username, errors.New("Invalid username length!")
	}
	exp := regexp.MustCompile(_EXP_USERNAME)
	if !exp.MatchString(username) {
		return username, errors.New("Invalid username!")
	}
	return username, nil
}

const _EXP_EMAIL = `^[A-Za-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

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

const _EXP_STRING = "^[^'\"`]+$"

func CheckNotEmpty(st string) (string, error) {
	length := len(st)
	if length == 0 {
		return st, errors.New("Please fill in all input fields!")
	}
	exp := regexp.MustCompile(_EXP_STRING)
	if !exp.MatchString(st) {
		return st, errors.New("Invalid character!")
	}
	return st, nil
}

func GetMD5Pwd(password string) (string, error) {
	length := len(password)
	data := md5.New()
	data.Write([]byte(password))
	Md5Pwd := hex.EncodeToString(data.Sum(nil))
	if length < 6 || length > 20 {
		return Md5Pwd, errors.New("Password's length must be more than 5 and less then 21!")
	}
	return Md5Pwd, nil
}

const _EXP_ENGLISH = `^[a-zA-Z0-9._\-]+$`

func CheckEnglish(st string) (string, error) {
	length := len(st)
	if length == 0 {
		return st, errors.New("Please fill in all input fields!")
	}
	exp := regexp.MustCompile(_EXP_ENGLISH)
	if !exp.MatchString(st) {
		return st, errors.New("Invalid English team name!")
	}
	return st, nil
}

const _EXP_PT = `^[ a-zA-Z]+$`

func CheckPT(st string) (string, error) {
	length := len(st)
	if length == 0 {
		return st, errors.New("Please fill in all input fields!")
	}
	exp := regexp.MustCompile(_EXP_ENGLISH)
	if !exp.MatchString(st) {
		return st, errors.New("Invalid English member name!")
	}
	return st, nil
}

func UpdateTeamsByCoach(coach Coach) error {
	var (
			team  Teams
			teams []Teams
			_     error
	)
	team.Coach = &coach
	teams, _ = team.GetTeamsByCoach("ZhuHai")
	for _, val := range teams {
		val.School = coach.School
			val.Coachnamech = coach.Chname
			val.Coachnameen = coach.Enname
			val.Update()
	}
	teams, _ = team.GetTeamsByCoach("GuangZhou")
	for _, val := range teams {
		val.School = coach.School
			val.Coachnamech = coach.Chname
			val.Coachnameen = coach.Enname
			val.Update()
	}
	return nil
}
