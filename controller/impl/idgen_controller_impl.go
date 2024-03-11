package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/idgenms/service"
	"github.com/prakash-p-3121/restlib"
)

type IDGenControllerImpl struct {
}

func (controller *IDGenControllerImpl) NextID(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	ctx := ginRestCtx.CtxGet()
	if !ok {
		err := errorlib.NewInternalServerError("expected-gin-context")
		err.SendRestResponse(ctx)
		return
	}
	var req *idgenmodel.IDGenReq
	err := ctx.BindJSON(&req)
	if err != nil {
		appErr := errorlib.NewInternalServerError("payload-serialization")
		appErr.SendRestResponse(ctx)
		return
	}

	idGenService := service.NewIDGenService(ginRestCtx.CtxGet())
	resp, appErr := idGenService.NextID(req)
	if err != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, resp)
}
