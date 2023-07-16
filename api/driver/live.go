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
	ImageBase64 string `json:"image_base64" binding:"required"`
}

type RpcOutput struct {
	User *schema.UserOutput `json:"user"`
}

func InitLiveRouter(gr *gin.RouterGroup, conn *gorm.DB, client pb.ImageServiceClient) {
	// @Summary 顔識別
	// @Tags live
	// @Produce  json
	// @securityDefinitions.basic BasicAuth
	// @Param       body body     RpcInput false "postの新規作成"
	// @Success 200 {object} RcpOutput
	// @Failure 400 {object} error
	// @Router /live/stream [post]
	gr.POST("/stream",
		func(ctx *gin.Context) {
			input := RpcInput{}
			if err := ctx.ShouldBindJSON(&input); err != nil {
				log.Print("bad reqwest")
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return

			}
			req := &pb.ImageBase64{Base: input.ImageBase64}
			res, err := client.ImageReqBase64(ctx, req, grpc.WaitForReady(true))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Msg),
			})
		})
}
