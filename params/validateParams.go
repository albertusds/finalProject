package params

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func ValidateRq(s interface{}) error {
	var errCreate error

	switch s.(type) {
	case RegisterUserRq:
		govalidator.CustomTypeTagMap.Set("minagevalidation", func(i, o interface{}) bool {
			inputAge := i.(int)
			return inputAge > MIN_AGE_VALIDATION
		})
		_, errCreate = govalidator.ValidateStruct(s)

	case UserLoginRq, UpdateUserRq, SavePhotoRq, CreateCommentRq, UpdateCommentRq, SocMedRq:
		_, errCreate = govalidator.ValidateStruct(s)

	default:
		errCreate = errors.New("type is not match")
	}

	return errCreate
}
