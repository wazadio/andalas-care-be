package main

import (
	"andalas-care/configs/database"
	"andalas-care/configs/env"
	"andalas-care/configs/redis"
	"andalas-care/configs/twilio_config"
	"andalas-care/internal/entity/user_entity"
	"andalas-care/internal/router"
	"log"
)

func main() {
	envConfig := env.NewEnvConfig(".env")
	envConfig.LoadEnv()

	db := database.NewDBconnection()
	db.Debug().AutoMigrate(user_entity.User{})

	rdb := redis.NewRedisClient()

	myTwilio := twilio_config.NewTwilioClient()

	r := router.NewRouter(db, rdb, myTwilio)

	log.Fatal(r.Run(":8080"))
}
