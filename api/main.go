package main

import (
	"rujukan/controller/aplikasi/user"
	"rujukan/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	model.DBConnection()
	router.GET("/aplikasi/user", user.Index)
	//http.ListenAndServe(":8080", nil)
	router.Run(":8081") // listen and serve on 0.0.0.0:8080
}
