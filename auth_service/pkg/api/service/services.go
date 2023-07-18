package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/repository/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/support"
)

type authService struct {
	Repo interfaces.AuthRepo
	pb.UnimplementedAuthServiceServer
}

func NewauthUseCase(repo interfaces.AuthRepo) pb.AuthServiceServer {
	return &authService{
		Repo: repo,
	}
}

func (cr *authService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if _, err := cr.Repo.FindbyEmail(ctx, req.Email); err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
		}, errors.New("user already exists")
	}
	referalcode := support.ReferalCodeGenerator()
	hash, err := support.HashPassword(req.Password)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	req.Password = hash
	req.Referalcode = referalcode
	if err := cr.Repo.UserSignup(context.Background(), req); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	respSid, err1 := cr.TwilioSendOTP(ctx, req.Phone)
	if err1 != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.RegisterResponse{
		Status:     http.StatusOK,
		Responseid: respSid,
	}, nil
}

func (cr *authService) RegisterOtpVerify(ctx context.Context, req *pb.OtpVerifyRequest) (*pb.OtpVerifyResponse, error) {
	session, err := cr.TwilioVerifyOTP(ctx, req)
	if err != nil {
		return &pb.OtpVerifyResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	err1 := cr.Repo.UpdateVerify(session.MobileNum)
	if err1 != nil {
		return &pb.OtpVerifyResponse{
			Status: http.StatusInternalServerError,
		}, err1
	}
	return &pb.OtpVerifyResponse{
		Status:   http.StatusOK,
		Response: "user signin success",
	}, nil
}

func (cr *authService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := cr.Repo.FindbyEmail(ctx, req.Email)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	ok := support.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, errors.New("wrong password")
	}
	tokenString, err1 := support.GenerateJWT(user.ID)
	if err1 != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, err1
	}
	data := &pb.User{
		Id:          uint32(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Phone:       user.MobileNum,
		Referalcode: user.ReferalCode,
		Block:       user.Block,
		Verified:    user.Verified,
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		User:   data,
		Token:  tokenString,
	}, nil
}

func (cr *authService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := support.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(claims.ID),
	}, nil
}

func (cr *authService) AdminRegister(ctx context.Context, req *pb.RegisterRequest) (*pb.OtpVerifyResponse, error) {
	if _, err := cr.Repo.AdminFindbyEmail(ctx, req.Email); err == nil {
		return &pb.OtpVerifyResponse{
			Status: http.StatusBadRequest,
		}, errors.New("user already exists")
	}
	hash, err := support.HashPassword(req.Password)
	if err != nil {
		return &pb.OtpVerifyResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	req.Password = hash
	data := domain.Admin{
		Name:      req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
		MobileNum: req.Phone,
	}
	if err := cr.Repo.SignUpAdmin(context.Background(), data); err != nil {
		return &pb.OtpVerifyResponse{
			Status: http.StatusInternalServerError,
		}, err
	}
	return &pb.OtpVerifyResponse{
		Status:   http.StatusOK,
		Response: "admin registered",
	}, nil
}

func (cr *authService) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := cr.Repo.AdminFindbyEmail(ctx, req.Email)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, err
	}
	ok := support.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, errors.New("wrong password")
	}
	tokenString, err1 := support.GenerateJWT(user.ID)
	if err1 != nil {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, err1
	}
	data := &pb.User{
		Id:        uint32(user.ID),
		FirstName: user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.MobileNum,
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		User:   data,
		Token:  tokenString,
	}, nil
}
