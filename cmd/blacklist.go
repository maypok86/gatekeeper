package cmd

import (
	"context"
	"log"
	"time"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootBlacklistCmd = &cobra.Command{
	Use:   "blacklist",
	Short: "Add subnet to blacklist or remove subnet from blacklist",
	Long:  `Add ip/subnet to blacklist or remove ip/subnet from blacklist`,
}

var blacklistAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add ip/subnet to blacklist",
	Long:    `Add ip/subnet to blacklist`,
	Example: "gk blacklist add <ip/subnet>",
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
		subnet := args[0]
		response, err := client.BlacklistAdd(ctx, &apipb.BlacklistAddRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not add to blacklist: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

var blacklistRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove ip/subnet from blacklist",
	Long:    `Remove ip/subnet from blacklist`,
	Example: "gk blacklist remove <ip/subnet>",
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
		subnet := args[0]
		response, err := client.BlacklistRemove(ctx, &apipb.BlacklistRemoveRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not remove from blacklist: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

func init() {
	rootBlacklistCmd.AddCommand(blacklistAddCmd)
	rootBlacklistCmd.AddCommand(blacklistRemoveCmd)
	rootCmd.AddCommand(rootBlacklistCmd)
}
