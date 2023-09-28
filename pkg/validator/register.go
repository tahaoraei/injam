package validator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"injam/param"
	"regexp"
)

func (v Validator) Register(request param.UserRegisterRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&request.Password, validation.Required, validation.Match(regexp.MustCompile(`[A-z0-9#!*%@]{5,}`))),
		validation.Field(&request.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09[0-9]{9}$`)), validation.By(v.phoneNumberUniqueness)),
	)

}

func (v Validator) phoneNumberUniqueness(value interface{}) error {
	phone := value.(string)
	b, err := v.repo.IsPhoneNumberUnique(phone)

	if b {
		return nil
	}
	if !b && err == nil {
		return errors.New("phone is not unique")
	}
	return err
}
