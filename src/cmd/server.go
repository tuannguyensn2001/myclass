package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"myclass/src/config"
	enumconfig "myclass/src/enums/config"
	"myclass/src/middlewares"
	"myclass/src/services/auth/authprovider"
)

type Provider = func(r *gin.Engine) error

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {

			cfg := config.GetConfig()

			if cfg.Server.Mode == enumconfig.Production {
				gin.SetMode(gin.ReleaseMode)
			}

			r := gin.Default()

			r.Use(middlewares.Recover)

			providers := []Provider{authprovider.Boot}

			for _, item := range providers {
				err := item(r)
				if err != nil {
					log.Fatalln(err)
				}
			}

			r.Run(":" + cfg.Server.Port)
		},
	}
}
