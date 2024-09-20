package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s is required (type: %s)", name, typ)
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Salary   int64  `json:"salary"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Salary   int64  `json:"salary"`
	Password string `json:"password"`
}

func (schema *CreateUserRequest) Validate() error {
	if schema.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if schema.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if schema.Password == "" {
		return errParamIsRequired("password", "string")
	}

	if schema.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}

	return nil
}

func (schema *UpdateUserRequest) Validate() error {
	if schema.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if schema.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if schema.Password == "" {
		return errParamIsRequired("password", "string")
	}

	if schema.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}

	return nil
}
