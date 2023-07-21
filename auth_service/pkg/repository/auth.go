package repository

import (
	"context"
	"errors"
	"time"

	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/repository/interfaces"
	"github.com/stebinsabu13/ecommerce_microservice/auth_service/pkg/utils"
	"gorm.io/gorm"
)

type authRepo struct {
	DB *gorm.DB
}

func NewauthRepo(db *gorm.DB) interfaces.AuthRepo {
	return &authRepo{
		DB: db,
	}
}

func (c *authRepo) UserSignup(ctx context.Context, req *pb.RegisterRequest) (uint, error) {
	var userid uint
	query1 := `insert into users(created_at,updated_at,first_name,last_name,email,mobile_num,password,referal_code)values($1,$2,$3,$4,$5,$6,$7,$8) returning id`
	if err := c.DB.Raw(query1, time.Now(), time.Now(), req.FirstName, req.LastName, req.Email, req.Phone, req.Password, req.Referalcode).Scan(&userid).Error; err != nil {
		return userid, err
	}
	return userid, nil
}

func (c *authRepo) SaveOtp(ctx context.Context, otpsession domain.OtpSession) error {
	err := c.DB.Create(&otpsession).Error
	return err
}

func (c *authRepo) RetrieveSession(ctx context.Context, otpid string) (domain.OtpSession, error) {
	var session domain.OtpSession
	err := c.DB.Where("otp_id=?", otpid).Find(&session).Error
	if err != nil {
		return session, err
	}
	return session, nil
}

func (c *authRepo) UpdateVerify(number string) error {
	if err := c.DB.Model(&domain.User{}).Where("mobile_num=?", number).UpdateColumn("verified", true).Error; err != nil {
		return err
	}
	return nil
}

func (c *authRepo) FindbyEmail(ctx context.Context, email string) (utils.ResponseUsers, error) {
	var user utils.ResponseUsers
	query := `SELECT * from users where email=$1`
	c.DB.Raw(query, email).Scan(&user)
	if user.ID == 0 {
		return user, errors.New("invalid email")
	}
	if user.Block {
		return user, errors.New("you are blocked")
	}
	if !user.Verified {
		return user, errors.New("you need to verify your mobile number")
	}
	return user, nil
}

func (c *authRepo) AdminFindbyEmail(ctx context.Context, email string) (domain.Admin, error) {
	var admin domain.Admin
	if err := c.DB.Where("email=?", email).Find(&admin).Error; err != nil {
		return admin, err
	}
	if admin.ID == 0 {
		return domain.Admin{}, errors.New("invalid email")
	}
	return admin, nil
}

func (c *authRepo) SignUpAdmin(ctx context.Context, admin domain.Admin) error {
	if err := c.DB.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func (c *authRepo) AddAddress(ctx context.Context, address domain.Address) error {
	err := c.DB.Create(&address).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *authRepo) ShowAddress(ctx context.Context, id uint) ([]*pb.UserAddress, error) {
	var address []*pb.UserAddress
	query := `select id,house_name,street,city,state,country,pincode from addresses where user_id=?`
	if err := c.DB.Raw(query, id).Scan(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (c *authRepo) Getaddress(addresid uint) ([]*pb.UserAddress, error) {
	var address []*pb.UserAddress
	query := `select id,house_name,street,city,state,country,pincode from addresses where id=?`
	if err := c.DB.Raw(query, addresid).Scan(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}
