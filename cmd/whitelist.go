package cmd

import (
	"context"
	"log"
	"time"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootWhitelistCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Add subnet to whitelist or remove subnet from whitelist",
	Long:  `Add ip/subnet to whitelist or remove ip/subnet from whitelist`,
}

var whitelistAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add ip/subnet to whitelist",
	Long:    `Add ip/subnet to whitelist`,
	Example: "gk whitelist add <ip/subnet>",
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
		response, err := client.WhitelistAdd(ctx, &apipb.WhitelistAddRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not add to whitelist: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

var whitelistRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove ip/subnet from whitelist",
	Long:    `Remove ip/subnet from whitelist`,
	Example: "gk whitelist remove <ip/subnet>",
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
		response, err := client.WhitelistRemove(ctx, &apipb.WhitelistRemoveRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not add to whitelist: %v", err)
		}
		log.Printf("Success: %t", response.Ok)
	},
}

func init() {
	rootWhitelistCmd.AddCommand(whitelistAddCmd)
	rootWhitelistCmd.AddCommand(whitelistRemoveCmd)
	rootCmd.AddCommand(rootWhitelistCmd)
}
