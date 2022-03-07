package validation

import (
	"DRCache/global"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UserVerify(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		if data != "" {
			return true
		}
	}
	return false
}

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func PasswordVerify(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		var level = levelD
		patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
		for _, pattern := range patternList {
			match, _ := regexp.MatchString(pattern, data)
			if match {
				level++
			}
		}

		if level < global.Settings.PasswordLevel {
			return false
		}
		return true
	}
	return false
}
