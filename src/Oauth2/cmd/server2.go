package cmd

import (
	src "OAUTH2/src/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(server2Cmd)
}

var server2Cmd = &cobra.Command{
	Use: "server2",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORS())
		src.InitRouter(e)
		e.Logger.Fatal(e.Start(":8080"))
	},
}
