package user

import (
	"rujukan/model"
	modeluser "rujukan/model/aplikasi/user"

	"github.com/gin-gonic/gin"
)

func DataAll(c *gin.Context) {
	model.SwitchDatabase("aplikasi")
	respon, _ := modeluser.QueryAll()
	c.JSON(200, gin.H{
		"success": len(respon) > 0,
		"total":   len(respon),
		"data":    respon,
	})
}
