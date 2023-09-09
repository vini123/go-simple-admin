package utils

import "regexp"

func IsPhone(value string) bool {

	reg := regexp.MustCompile(`^1[3456789]{1}\d{9}$`)

	return reg.MatchString(value)
}

func IsEmail(value string) bool {

	reg := regexp.MustCompile(`^([a-zA-Z\\d][\w-]{2,})@(\w{2,})\.([a-z]{2,})(\.[a-z]{2,})?$`)

	return reg.MatchString(value)
}
