package controller

import "github.com/prakash-p-3121/restlib"

type IDGenController interface {
	NextID(restCtx restlib.RestContext)
}
