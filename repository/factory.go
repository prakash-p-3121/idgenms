package repository

import "github.com/prakash-p-3121/idgenms/repository/impl"

func NewIDGenRepository() IDGenRepository {
	return &impl.IDGenRepositoryImpl{}
}
