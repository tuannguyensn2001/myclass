package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"myclass/src/models"
	validatorpkg "myclass/src/packages/validators"
)

func testValidate() *cobra.Command {
	return &cobra.Command{
		Use: "test-validate",
		Run: func(cmd *cobra.Command, args []string) {
			user := models.User{
				//Username: "abbbbbbbbbbbbbbbbbbbbb",
			}

			err := validatorpkg.ExecWithKeyError(user, "auth")

			log.Fatalln(err)
		},
	}
}
