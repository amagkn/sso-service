package dto

import ssov1 "github.com/amagkn/sso-protos/gen/go/sso"

type LoginInput struct {
	AppID    int32  `json:"app_id" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginOutput struct {
	Token string
}

func (lo *LoginOutput) ToLoginResponse() *ssov1.LoginResponse {
	return &ssov1.LoginResponse{
		Token: lo.Token,
	}
}
