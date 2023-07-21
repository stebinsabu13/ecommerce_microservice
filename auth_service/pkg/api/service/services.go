package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	client "github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/repository/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/support"
)

type authService struct {
	Repo   interfaces.AuthRepo
	Client client.CartClient
	pb.UnimplementedAuthServiceServer
}

func NewauthUseCase(repo interfaces.AuthRepo, client client.CartClient) pb.AuthServiceServer {
	return &authService{
		Repo:   repo,
		Client: client,
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
	userid, err2 := cr.Repo.UserSignup(context.Background(), req)
	if err2 != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
		}, err2
	}
	if err := cr.Client.CreateCart(userid); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadGateway,
		}, errors.New("failed to create cart")
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
		}, err
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

func (cr *authService) AddUserAddress(ctx context.Context, req *pb.AddUserAddressRequest) (*pb.AddUserAddressResponse, error) {
	address := domain.Address{
		HouseName: req.HouseName,
		Street:    req.Street,
		City:      req.City,
		State:     req.State,
		Country:   req.Country,
		Pincode:   req.Pincode,
		UserID:    uint(req.Userid),
	}
	fmt.Println(address)
	if err := cr.Repo.AddAddress(ctx, address); err != nil {
		return nil, err
	}
	return &pb.AddUserAddressResponse{
		Status:   http.StatusOK,
		Response: "Address addedd",
	}, nil
}

func (cr *authService) ShowUserAddress(ctx context.Context, req *pb.ShowUserAddressRequest) (*pb.ShowUserAddressResponse, error) {
	res, err := cr.Repo.ShowAddress(ctx, uint(req.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.ShowUserAddressResponse{
		Status:  http.StatusOK,
		Address: res,
	}, nil
}

func (cr *authService) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.ShowUserAddressResponse, error) {
	fmt.Println("inside service")
	res, err := cr.Repo.Getaddress(uint(req.Addressid))
	if err != nil {
		return nil, err
	}
	return &pb.ShowUserAddressResponse{
		Status:  http.StatusOK,
		Address: res,
	}, nil
}
