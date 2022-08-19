package logpkg

import "go.uber.org/zap"

var z *zap.Logger

func Init() error {
	log, err := zap.NewProduction()
	if err != nil {
		return err
	}

	z = log
	return nil
}

func SyncLog() {
	z.Sync()
}

func Sugar() *zap.SugaredLogger {
	return z.Sugar()
}
