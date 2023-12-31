package client

import (
	"context"
	"fmt"

	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/client/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/service"
	"github.com/stebinsabu13/ecommerce_microservice/api_gateway/pkg/utils"
)

type authclient struct {
	Server pb.AuthServiceClient
}

func NewauthClient(service service.Clients) interfaces.AuthClient {
	return &authclient{
		Server: service.Authcli,
	}
}

func (cr *authclient) UserSignup(ctx context.Context, body utils.BodySignUpuser) (*pb.RegisterResponse, error) {
	res, err := cr.Server.Register(ctx, &pb.RegisterRequest{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.MobileNum,
		Password:  body.Password,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authclient) UserSignupOtpVerify(ctx context.Context, body utils.Otpverify) (*pb.OtpVerifyResponse, error) {
	res, err := cr.Server.RegisterOtpVerify(ctx, &pb.OtpVerifyRequest{
		Otp:   body.Otp,
		OtpId: body.OtpID,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authclient) UserLogin(ctx context.Context, body utils.BodyLogin) (*pb.LoginResponse, error) {
	res, err := cr.Server.Login(ctx, &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authclient) AuthorizationMiddleware(tokenstr string) (*pb.ValidateResponse, error) {
	res, err := cr.Server.Validate(context.Background(), &pb.ValidateRequest{
		Token: tokenstr,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authclient) AdminSignup(ctx context.Context, body utils.BodySignUpuser) (*pb.OtpVerifyResponse, error) {
	res, err := cr.Server.AdminRegister(ctx, &pb.RegisterRequest{
		FirstName: body.FirstName,
		Email:     body.Email,
		Phone:     body.MobileNum,
		Password:  body.Password,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authclient) AdminLogin(ctx context.Context, body utils.BodyLogin) (*pb.LoginResponse, error) {
	res, err := cr.Server.AdminLogin(ctx, &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *authclient) ShowAddress(ctx context.Context, id uint) (*pb.ShowUserAddressResponse, error) {
	address, err := c.Server.ShowUserAddress(ctx, &pb.ShowUserAddressRequest{
		Userid: uint32(id),
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (c *authclient) AddAddress(ctx context.Context, body utils.AddAddress, id uint) (*pb.AddUserAddressResponse, error) {
	fmt.Println(id)
	address, err := c.Server.AddUserAddress(ctx, &pb.AddUserAddressRequest{
		Userid:    uint32(id),
		HouseName: body.HouseName,
		Street:    body.Street,
		City:      body.City,
		State:     body.State,
		Country:   body.Country,
		Pincode:   body.Pincode,
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}
