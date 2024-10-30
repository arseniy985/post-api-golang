package database

import "github.com/asaskevich/govalidator"

func ValidatePostData(title, content string) bool {
	if !govalidator.MinStringLength(title, "4") {
		return false
	}
	if !govalidator.MinStringLength(content, "4") {
		return false
	}
	if !govalidator.MaxStringLength(title, "80") {
		return false
	}
	if !govalidator.MaxStringLength(content, "20000") {
		return false
	}
	return true
}
