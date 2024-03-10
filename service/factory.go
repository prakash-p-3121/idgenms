package service

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/idgenms/service/impl"
)

func NewIDGenService(ginCtx *gin.Context) IDGenService {
	return &impl.IDGenServiceImpl{Ctx: ginCtx}
}
