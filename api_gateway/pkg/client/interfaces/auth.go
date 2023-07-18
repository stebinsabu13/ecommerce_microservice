package interfaces

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type AuthClient interface {
	UserSignup(context.Context, utils.BodySignUpuser) (*pb.RegisterResponse, error)
	UserSignupOtpVerify(context.Context, utils.Otpverify) (*pb.OtpVerifyResponse, error)
	UserLogin(context.Context, utils.BodyLogin) (*pb.LoginResponse, error)

	//token validation
	AuthorizationMiddleware(string) (*pb.ValidateResponse, error)

	//Admin
	AdminSignup(context.Context, utils.BodySignUpuser) (*pb.OtpVerifyResponse, error)
	AdminLogin(context.Context, utils.BodyLogin) (*pb.LoginResponse, error)
}
