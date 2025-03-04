package greeting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// func Greeting(r *gin.Engine) {
// 	r.GET("/greeting", greet)
// }

func Greet(c *gin.Context) {
	appName := viper.GetString("App.name")
	appDesc := viper.GetString("App.desc")
	c.JSON(http.StatusOK, gin.H{
		"message": "Naija Prime API",
		"app name": `Welcome to ` + appName,
		"app description": appDesc,
	})
}