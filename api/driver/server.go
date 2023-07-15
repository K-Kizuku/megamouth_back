package driver

import (
	"fmt"
	"log"
	"net/http"

	"megamouth/api/driver/db"
	"megamouth/api/utils/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// OpenAPI
	_ "megamouth/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve(addr string) {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(cors.Default())

	//health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello hackz!!",
		})
	})

	//OpenAPI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			ur := v1.Group("/user")
			pr := v1.Group("/post")

			InitUserRouter(ur, conn)
			InitPostRouter(pr, conn)
		}
	}
	if err := r.Run(fmt.Sprintf(":%s", config.ApiPort)); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
