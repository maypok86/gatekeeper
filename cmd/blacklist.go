package cmd //nolint:dupl

import (
	"context"
	"fmt"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
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
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.BlacklistAdd(ctx, &apipb.BlacklistAddRequest{
			Subnet: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not add to blacklist: %w", err)
		}
		return ok, nil
	}),
}

var blacklistRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove ip/subnet from blacklist",
	Long:    `Remove ip/subnet from blacklist`,
	Example: "gk blacklist remove <ip/subnet>",
	Args:    cobra.ExactArgs(1),
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.BlacklistRemove(ctx, &apipb.BlacklistRemoveRequest{
			Subnet: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not remove from blacklist: %w", err)
		}
		return ok, nil
	}),
}

func init() {
	rootBlacklistCmd.AddCommand(blacklistAddCmd)
	rootBlacklistCmd.AddCommand(blacklistRemoveCmd)
	rootCmd.AddCommand(rootBlacklistCmd)
}
