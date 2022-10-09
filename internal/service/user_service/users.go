package user_service

import (
	"andalas-care/configs/twilio_config"
	"andalas-care/internal/entity/user_entity"
	"andalas-care/internal/repository/user_repository"
	"andalas-care/pkg/constant"
	"andalas-care/pkg/jwt"
	"andalas-care/pkg/twilio"
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userService struct {
	DB       *gorm.DB
	ctx      context.Context
	myTwilio *twilio_config.TwilioClient
}

func NewUserService(db *gorm.DB, ctx context.Context, myTwilio *twilio_config.TwilioClient) *userService {
	return &userService{
		DB:       db,
		ctx:      ctx,
		myTwilio: myTwilio,
	}
}

func (u *userService) LoginWithPhoneNumber(phoneNumber string) error {

	sendOTPStatus, err := twilio.SendOTPSms(phoneNumber, u.myTwilio)
	if sendOTPStatus != "pending" {
		return errors.New("unexpected status (twilio)")
	}

	return err
}

func (u *userService) VerifyOTPSms(phoneNumber, otp string, rdb *redis.Client) (token string, err error) {

	var newUser user_entity.User

	verifyStatus, err := twilio.VerifyOTPSms(phoneNumber, otp, u.myTwilio)
	if err != nil {
		return
	}
	if !verifyStatus {
		return "", errors.New("failed verifying phone number")
	}

	userRepository := user_repository.NewUserRepository(u.DB)

	where := map[string]any{
		"phone_number": phoneNumber,
		"login_type":   constant.PhoneUser,
	}

	user, err := userRepository.GetUsers(where)
	if err != nil {
		return
	}

	if len(user) < 1 {
		newUser, err = userRepository.CreateUser(user_entity.User{
			PhoneNumber: phoneNumber,
			Status:      true,
			LoginType:   constant.PhoneUser,
		})

		if err != nil {
			return "", errors.New("error creating new user <phone>")
		}
	}

	token, err = jwt.GenerateJWTToken(int(newUser.ID))
	if err != nil {
		return
	}

	err = rdb.Set(u.ctx, phoneNumber, token, time.Hour*4).Err()

	return
}
