package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junmocsq/bookstore/api/apiv1"
	_ "github.com/junmocsq/bookstore/api/models"
	"github.com/sirupsen/logrus"

	_ "github.com/junmocsq/bookstore/api/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			图书阅读器 API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.WarnLevel)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	unlogin := r.Group("/apiv1")

	unlogin.Use(apiv1.Prepare())
	unlogin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := unlogin.Group("/")
	auth.Use(apiv1.CheckUser())
	{
		auth.GET("/user", apiv1.NewUserController().User)
		auth.POST("/user/update", apiv1.NewUserController().Update)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
