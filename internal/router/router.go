package router

import (
	"andalas-care/configs/twilio_config"
	"andalas-care/internal/handler/user_handler"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, rdb *redis.Client, myTwilio *twilio_config.TwilioClient) *gin.Engine {

	userHandler := user_handler.NewUserHandler(db, rdb, myTwilio)

	r := gin.Default()
	r.Use(middleware())
	r.POST("/user/phone/login", userHandler.LoginWithPhoneNumber)
	r.POST("/user/phone/verify", userHandler.VerifyOTPSms)

	return r
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
