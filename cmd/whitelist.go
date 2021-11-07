package cmd

import (
	"context"
	"fmt"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
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
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.WhitelistAdd(ctx, &apipb.WhitelistAddRequest{
			Subnet: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not add to whitelist: %w", err)
		}
		return response.GetOk(), nil
	}),
}

var whitelistRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove ip/subnet from whitelist",
	Long:    `Remove ip/subnet from whitelist`,
	Example: "gk whitelist remove <ip/subnet>",
	Args:    cobra.ExactArgs(1),
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.WhitelistRemove(ctx, &apipb.WhitelistRemoveRequest{
			Subnet: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not remove from whitelist: %w", err)
		}
		return ok, nil
	}),
}

func init() {
	rootWhitelistCmd.AddCommand(whitelistAddCmd)
	rootWhitelistCmd.AddCommand(whitelistRemoveCmd)
	rootCmd.AddCommand(rootWhitelistCmd)
}
