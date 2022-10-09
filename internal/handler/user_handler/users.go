package user_handler

import (
	"andalas-care/configs/twilio_config"
	"andalas-care/internal/service/user_service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userHandler struct {
	DB       *gorm.DB
	RDB      *redis.Client
	MyTwilio *twilio_config.TwilioClient
}

func NewUserHandler(db *gorm.DB, rdb *redis.Client, myTwilio *twilio_config.TwilioClient) *userHandler {
	return &userHandler{
		DB:       db,
		RDB:      rdb,
		MyTwilio: myTwilio,
	}
}

func (u *userHandler) LoginWithPhoneNumber(ctx *gin.Context) {
	payload := loginWithPhoneNumberRequest{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println("==================", payload)

	userService := user_service.NewUserService(u.DB, ctx, u.MyTwilio)

	err = userService.LoginWithPhoneNumber(payload.PhoneNumber)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (u *userHandler) VerifyOTPSms(ctx *gin.Context) {
	payload := verifyOTSSmsRequest{}

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println("==================", payload)

	userService := user_service.NewUserService(u.DB, ctx, u.MyTwilio)

	token, err := userService.VerifyOTPSms(payload.PhoneNumber, payload.Code, u.RDB)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
