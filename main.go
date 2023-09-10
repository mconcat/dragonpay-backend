package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLGlob("templates/*")

	//	r.POST("/telegram-oauth", telegramOAuth)
	//r.POST("/payment-request", handlePaymentRequest)
	r.GET("/login", login)
	r.GET("/login/callback", loginCallback)
	r.POST("/register", register)


	r.Run() // listen and serve on
}
