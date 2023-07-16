package driver

import (
	"fmt"
	"log"
	pb "megamouth/grpc"
	"net/http"

	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type RpcInput struct {
	ImageURL string `json:"image_url" binding:"required"`
}

type RpcOutput struct {
	User *schema.UserOutput `json:"user"`
}

func InitLiveRouter(lr *gin.RouterGroup, conn *gorm.DB, client pb.ImageServiceClient) {
	// @Summary 顔識別
	// @Tags live
	// @Produce  json
	// @securityDefinitions.basic BasicAuth
	// @Param       body body     RpcInput false "postの新規作成"
	// @Success 200 {object} RcpOutput
	// @Failure 400 {object} error
	// @Router /live/stream [post]
	lr.POST("/stream",
		func(ctx *gin.Context) {
			input := RpcInput{}
			if err := ctx.ShouldBindJSON(&input); err != nil {
				log.Print("bad reqwest")
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return

			}
			req := &pb.ImageURL{Url: input.ImageURL}
			log.Print(input.ImageURL)
			res, err := client.ImageReq(ctx, req, grpc.WaitForReady(true))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error1": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Msg),
			})
		})
}
