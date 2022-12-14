package cmd

import (
	"log"

	"server/internal/web"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	WebCmd.Flags().IntP("port", "p", 8080, "Port to listen on")
}

// VersionCmd is the command to print the version of the server
var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run:   runWebCmd,
}

func runWebCmd(cmd *cobra.Command, args []string) {
	log.Println("Starting web server")

	r := gin.Default()
	r.GET("/", web.Index)

	r.Run()
}
