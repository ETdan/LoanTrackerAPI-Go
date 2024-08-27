package main

import (
	"LoanAPI/LoanTrackerAPI-Go/Delivery/routers"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	DB, err := infrastructure.ConnectDB()
	if err != nil {
		panic(err)
	}

	routers.RouterSetUp(server, DB)
	server.Run()
}
