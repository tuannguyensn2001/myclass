package baserepository

import (
	"context"
	"gorm.io/gorm"
	"myclass/src/config"
)

func GetDBFromContext(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return config.GetConfig().Db
	}

	return val
}

func SetDBWithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}
