package interfaces

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/utils"
)

type AuthRepo interface {
	UserSignup(context.Context, *pb.RegisterRequest) error
	SaveOtp(context.Context, domain.OtpSession) error
	RetrieveSession(context.Context, string) (domain.OtpSession, error)
	UpdateVerify(string) error
	FindbyEmail(context.Context, string) (utils.ResponseUsers, error)

	//admin
	AdminFindbyEmail(ctx context.Context, email string) (domain.Admin, error)
	SignUpAdmin(ctx context.Context, admin domain.Admin) error
}
