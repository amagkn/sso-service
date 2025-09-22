package ver1

import (
	"context"

	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/amagkn/sso-service/pkg/validation"
	"google.golang.org/grpc/codes"
)

func (h *Handlers) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	input := dto.LoginInput{
		AppID:    req.GetAppId(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")

		return nil, errorResponse(codes.InvalidArgument, errorPayload{Type: base_errors.Validation, Details: invalidFields})
	}

	output, err := h.uc.Login(ctx, input)
	if err != nil {
		logger.Error(err, "h.uc.Login")

		return nil, errorResponse(codes.Internal, errorPayload{Type: base_errors.InternalServer})
	}

	return output.ToLoginResponse(), nil
}
