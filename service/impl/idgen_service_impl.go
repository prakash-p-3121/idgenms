package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/idgenms/repository"
	"log"
)

type IDGenServiceImpl struct {
	Ctx *gin.Context
}

func (service *IDGenServiceImpl) NextID(req *idgenmodel.IDGenReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}
	log.Println("req=%d", *req.TableName)
	idGenRepo := repository.NewIDGenRepository()
	nextIDStr, bitCount, err := idGenRepo.NextIDGet(service.Ctx.Request.Context(), *req.TableName)
	if err != nil {
		log.Println("err=" + err.Error())
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	log.Println(nextIDStr)
	return &idgenmodel.IDGenResp{
		ID:       nextIDStr,
		BitCount: int64(bitCount),
	}, nil
}
