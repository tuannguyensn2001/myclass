package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"myclass/src/config"
	enumconfig "myclass/src/enums/config"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {

			cfg := config.GetConfig()

			if cfg.Server.Mode == enumconfig.Production {
				gin.SetMode(gin.ReleaseMode)
			}

			r := gin.Default()

			r.Run(":" + cfg.Server.Port)
		},
	}
}
