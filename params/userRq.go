package params

const MIN_AGE_VALIDATION = 8

type RegisterUserRq struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"required~email is required,email~invalid email format"`
	Password string `json:"password" valid:"required~password is required,minstringlength(6)~Password must have a minimum of 6 characters"`
	Age      int    `json:"age" valid:"required~Age is required,minagevalidation~Age must be above 8"`
}

type UserLoginRq struct {
	Email    string `json:"email" valid:"required~email is required,email~invalid email format"`
	Password string `json:"password" valid:"required~password is required,minstringlength(6)~Password must have a minimum of 6 characters"`
}

type UpdateUserRq struct {
	Email    string `json:"email" valid:"required~email is required,email~invalid email format"`
	Username string `json:"username" valid:"required~username is required"`
}

// func ValidateRq(s interface{}) error {
// 	var errCreate error

// 	switch s.(type) {
// 	case RegisterUserRq:
// 		govalidator.CustomTypeTagMap.Set("minagevalidation", func(i, o interface{}) bool {
// 			inputAge := i.(int)
// 			return inputAge > MIN_AGE_VALIDATION
// 		})
// 		_, errCreate = govalidator.ValidateStruct(s)

// 	case UserLoginRq, UpdateUserRq:
// 		_, errCreate = govalidator.ValidateStruct(s)

// 	default:
// 		errCreate = errors.New("type is not match")
// 	}

// 	return errCreate
// }
