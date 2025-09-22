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

func (h *Handlers) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	input := dto.RegisterInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")

		return nil, errorResponse(codes.InvalidArgument, errorPayload{Type: base_errors.Validation, Details: invalidFields})
	}

	output, err := h.uc.Register(ctx, input)
	if err != nil {
		logger.Error(err, "h.uc.Register")

		return nil, errorResponse(codes.Internal, errorPayload{Type: base_errors.InternalServer})
	}

	return output.ToRegisterResponse(), nil
}
