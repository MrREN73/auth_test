package main

import (
	"auth/src/common"
	"auth/src/common/healthcheck"
	"auth/src/common/utils"
	"auth/src/handler"
	"auth/src/service"
	"log"
	"time"
)

const secondsToRoughStop = 5

func main() {
	dbConfig := common.Config{}
	db := common.Connect(dbConfig)

	s := service.New(db)
	healthcheck.Serve(s.GetDB().DB)

	srvStop := handler.Serve(s)

	log.Printf("\nStart listening\n")

	utils.GracefulStop(secondsToRoughStop*time.Second, srvStop)
}
