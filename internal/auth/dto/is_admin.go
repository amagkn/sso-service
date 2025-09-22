package dto

import ssov1 "github.com/amagkn/sso-protos/gen/go/sso"

type IsAdminInput struct {
	UserId int64 `json:"user_id" validate:"required"`
}

type IsAdminOutput struct {
	IsAdmin bool
}

func (ia *IsAdminOutput) ToIsAdminResponse() *ssov1.IsAdminResponse {
	return &ssov1.IsAdminResponse{
		IsAdmin: ia.IsAdmin,
	}
}
