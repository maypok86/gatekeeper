package cmd

import (
	"context"
	"log"
	"time"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset bucket by login or ip",
	Long:  `Reset bucket by login or ip`,
}

var resetLoginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Reset bucket by login",
	Long:    `Reset bucket by login`,
	Example: "gk reset login <login>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(serverURL, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer func() {
			_ = conn.Close()
		}()
		client := apipb.NewGatekeeperServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		login := args[0]
		response, err := client.ResetLogin(ctx, &apipb.ResetLoginRequest{
			Login: login,
		})
		if err != nil {
			log.Fatalf("could not reset: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

var resetIPCmd = &cobra.Command{
	Use:     "ip",
	Short:   "Reset bucket by ip",
	Long:    `Reset bucket by ip`,
	Example: "gk reset ip <ip>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(serverURL, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer func() {
			_ = conn.Close()
		}()
		client := apipb.NewGatekeeperServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		ip := args[0]
		response, err := client.ResetIP(ctx, &apipb.ResetIPRequest{
			Ip: ip,
		})
		if err != nil {
			log.Fatalf("could not reset: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

func init() {
	rootResetCmd.AddCommand(resetLoginCmd)
	rootResetCmd.AddCommand(resetIPCmd)
	rootCmd.AddCommand(rootResetCmd)
}
