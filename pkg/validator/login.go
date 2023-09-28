package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"injam/param"
	"regexp"
)

func (v Validator) Login(request param.UserLoginRequest) error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09[0-9]{9}$`)), validation.By(v.isPhoneNumberExist)),
		validation.Field(&request.Password, validation.Required),
	)
}

func (v Validator) isPhoneNumberExist(value interface{}) error {
	number := value.(string)
	exist, err := v.repo.IsPhoneNumberUnique(number)
	if !exist && err == nil {
		return nil
	}
	return err
}
