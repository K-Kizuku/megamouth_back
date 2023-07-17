package driver

import (
	"fmt"
	"log"
	"net/http"

	"megamouth/api/driver/db"
	"megamouth/api/utils/config"

	pb "megamouth/grpc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

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
	rpc, err := grpc.Dial(config.GRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer rpc.Close()
	clientS := pb.NewImageServiceClient(rpc)
	clientR := pb.NewImageRegistorClient(rpc)
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
			gr := v1.Group("/grpc")

			InitUserRouter(ur, conn)
			InitPostRouter(pr, conn)
			InitLiveRouter(gr, conn, clientS, clientR)
		}
	}
	if err := r.Run(fmt.Sprintf(":%s", config.ApiPort)); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
