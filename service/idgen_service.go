package service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
)

type IDGenService interface {
	NextID(req *idgenmodel.IDGenReq) (*idgenmodel.IDGenResp, errorlib.AppError)
}
