package dto

import ssov1 "github.com/amagkn/sso-protos/gen/go/sso"

type RegisterInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterOutput struct {
	UserId int64
}

func (ro *RegisterOutput) ToRegisterResponse() *ssov1.RegisterResponse {
	return &ssov1.RegisterResponse{UserId: ro.UserId}
}
