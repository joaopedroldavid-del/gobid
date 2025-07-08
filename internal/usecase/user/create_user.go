package user

import (
	"context"

	"github.com/joaopedroldavid-del/gobid/internal/validator"
)

type CreateUserRequest struct {
	UserName     string	`json:"user_name"`
	Email        string	`json:"email"`
	Password	 string	`json:"password"`
	Bio          string	`json:"bio"`
}

func (req CreateUserRequest) Valid (ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Email), "email", "this field must be a valid email address")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "this field cannot be empty")
	eval.CheckField(
		validator.MinChars(req.Bio, 10) &&
		validator.MaxChars(req.Bio, 255),
		"bio", 
		"this field must have a length between 10 and 255 characters",
	)
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "this field must have a minimum length of 8 characters")
	eval.CheckField(validator.Matchers(req.Email, validator.EmailRX), "email", "this field must be a valid email address")


	return eval
}