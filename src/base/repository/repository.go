package baserepository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"myclass/src/config"
)

func GetDBFromContext(ctx context.Context) *sqlx.DB {
	val, ok := ctx.Value("db").(*sqlx.DB)
	if !ok {
		return config.GetConfig().Db
	}

	return val
}
