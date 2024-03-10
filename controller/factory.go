package controller

import "github.com/prakash-p-3121/idgenms/controller/impl"

func NewIDGenController() IDGenController {
	return &impl.IDGenControllerImpl{}
}
