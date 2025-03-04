package routers

// import (
// 	"fmt"

// 	controller "github.com/TheCodeGhinux/TaskHub/taskhub/pkg/controllers/Media"
// 	"github.com/gin-gonic/gin"
// )

// func Media(router *gin.Engine, ApiVersion string) *gin.Engine {
// 	MediaController := controller.Media

// 	MediaGroup := router.Group(fmt.Sprintf("%v", ApiVersion))
// 	{
// 		MediaGroup.GET("/", MediaController)
// 		MediaGroup.GET("/genrate-upload-url/:id", MediaController.UploadUrl)
// 		MediaGroup.GET("/streamline-videos", MediaController.Streamline)
// 		MediaGroup.POST("/saveupload", MediaController.SaveUpload)
// 		MediaGroup.GET("/media-by-user/:id", MediaController.SaveUpload)
// 	}

// 	return router
// }
