package ver1

import (
	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
	"github.com/amagkn/sso-service/internal/auth/usecase"
)

type Handlers struct {
	uc *usecase.UseCase
	ssov1.UnimplementedAuthServer
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{uc: uc}
}
