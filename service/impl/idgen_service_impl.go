package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/idgenms/repository"
)

type IDGenServiceImpl struct {
	Ctx *gin.Context
}

func (service *IDGenServiceImpl) NextID(req *idgenmodel.IDGenReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}
	idGenRepo := repository.NewIDGenRepository()
	nextIDBytes, err := idGenRepo.NextIDGet(service.Ctx.Request.Context(), *req.TableName)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &idgenmodel.IDGenResp{
		ID:       nextIDBytes,
		BitCount: int64(len(nextIDBytes)),
	}, nil
}
