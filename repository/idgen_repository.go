package repository

import "context"

type IDGenRepository interface {
	NextIDGet(ctx context.Context, tableName string) (string, int, error)
}
