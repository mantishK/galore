package validate

import (
	"regexp"
	"strings"
)

func Password(password string) bool {
	if len(password) < 8 {
		return false
	}
	if !strings.ContainsAny(password, "1234567890") {
		return false
	}
	if strings.ContainsAny(password, " ") {
		return false
	}
	return true
}

func UserName(username string) bool {
	if strings.Contains(username, " ") {
		return false
	}
	exp, err := regexp.Compile(`<?\S+@\S+?>?`)
	if err != nil {
		return false
	}
	if !exp.MatchString(username) {
		return false
	}
	return true
}

func Url(url string) bool {
	return true
}
