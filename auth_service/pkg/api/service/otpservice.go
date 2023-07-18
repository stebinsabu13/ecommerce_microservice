package service

import (
	"context"

	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/config"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

func (c *authService) TwilioSendOTP(ctx context.Context, phoneNumber string) (string, error) {
	password := config.GetCofig().Auth_token
	userName := config.GetCofig().Account_sid
	seviceSid := config.GetCofig().Service_sid

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: password,
		Username: userName,
	})
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo("+91" + phoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(seviceSid, params)
	if err != nil {
		return *resp.Sid, err
	}
	otpsession := domain.OtpSession{
		OtpID:     *resp.Sid,
		MobileNum: phoneNumber,
	}
	err1 := c.Repo.SaveOtp(ctx, otpsession)
	if err1 != nil {
		return *resp.Sid, err1
	}
	return *resp.Sid, nil
}

func (c *authService) TwilioVerifyOTP(ctx context.Context, req *pb.OtpVerifyRequest) (domain.OtpSession, error) {
	//create a twilio client with twilio details
	password := config.GetCofig().Auth_token
	userName := config.GetCofig().Account_sid
	seviceSid := config.GetCofig().Service_sid
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: password,
		Username: userName,
	})
	session, err := c.Repo.RetrieveSession(ctx, req.OtpId)
	if err != nil {
		return session, err
	}
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + session.MobileNum)
	params.SetCode(req.Otp)
	resp, err1 := client.VerifyV2.CreateVerificationCheck(seviceSid, params)
	if err1 != nil {
		return session, err1
	} else if *resp.Status == "approved" {
		return session, nil
	}
	return session, nil
}
