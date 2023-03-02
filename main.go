package main

import (
	"os"
	"strconv"
	"swagger-gin/module/auth"
	"swagger-gin/module/config"
	"swagger-gin/module/controller"
	"swagger-gin/module/middlewares"

	docs "swagger-gin/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @securityDefinitions.apikey token
// @in header
// @name Authorization

func main() {
	config.ConnectDataBase()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{

		m := v1.Group("/m_user")
		{
			m.Use(middlewares.JwtAuthMiddleware())
			m.GET("/get-user", controller.UserGet)
			m.GET("/get-user/:id_user", controller.UserGetById)
			m.POST("/insert-user", controller.User)
			m.PUT("/update-user", controller.UserUpdate)
			m.DELETE("/delete-user/:id_user", controller.UserDeleteById)
		}

		i := v1.Group("/item")
		{
			i.Use(middlewares.JwtAuthMiddleware())
			i.GET("/get-item", controller.GetItem)
			i.POST("/insert-item", controller.InsertItem)
			i.PUT("/update-item", controller.UpdateItem)
			i.DELETE("/delete-item/:id", controller.DeleteItem)
		}

		l := v1.Group("/login")
		{
			l.POST("/login", auth.Login)
		}
	}

	persisauth, _ := strconv.ParseBool(os.Getenv("PERSISTAUTHORIZATION"))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.PersistAuthorization(persisauth),
		ginSwagger.DocExpansion(os.Getenv("DOCEXPANSION"))))
	r.Run(":8080")

}
