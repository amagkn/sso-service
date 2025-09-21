package ver1

import (
	"context"

	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
)

func (h *Handlers) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	return &ssov1.LoginResponse{Token: req.GetEmail()}, nil
}
