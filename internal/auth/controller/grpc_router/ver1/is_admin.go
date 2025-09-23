package ver1

import (
	"context"
	"errors"

	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
	"github.com/amagkn/sso-service/internal/auth/dto"
	"github.com/amagkn/sso-service/internal/auth/entity"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/amagkn/sso-service/pkg/validation"
	"google.golang.org/grpc/codes"
)

func (h *Handlers) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	input := dto.IsAdminInput{
		UserId: req.GetUserId(),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")

		return nil, errorResponse(codes.InvalidArgument, errorPayload{Type: base_errors.Validation, Details: invalidFields})
	}

	output, err := h.uc.IsAdmin(ctx, input)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			return nil, errorResponse(codes.NotFound, errorPayload{Type: entity.ErrUserNotFound})
		}

		return nil, errorResponse(codes.Internal, errorPayload{Type: base_errors.InternalServer})
	}

	return output.ToIsAdminResponse(), nil
}
