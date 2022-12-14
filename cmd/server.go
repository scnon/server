package cmd

import (
	"fmt"
	"log"
	"net"

	pb "server/internal/proto"
	"server/internal/server"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run:   runServerCmd,
}

func init() {
	ServerCmd.Flags().IntP("port", "p", 9090, "Port to listen on")
}

func runServerCmd(cmd *cobra.Command, args []string) {
	port := cmd.Flag("port").Value.String()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server.Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
