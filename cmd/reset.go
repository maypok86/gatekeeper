package cmd //nolint:dupl

import (
	"context"
	"fmt"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
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
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.ResetLogin(ctx, &apipb.ResetLoginRequest{
			Login: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not reset login: %w", err)
		}
		return ok, nil
	}),
}

var resetIPCmd = &cobra.Command{
	Use:     "ip",
	Short:   "Reset bucket by ip",
	Long:    `Reset bucket by ip`,
	Example: "gk reset ip <ip>",
	Args:    cobra.ExactArgs(1),
	Run: runGrpcClient(func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error) {
		response, err := client.ResetIP(ctx, &apipb.ResetIPRequest{
			Ip: arg,
		})
		ok := response.GetOk()
		if err != nil {
			return ok, fmt.Errorf("could not reset ip: %w", err)
		}
		return ok, nil
	}),
}

func init() {
	rootResetCmd.AddCommand(resetLoginCmd)
	rootResetCmd.AddCommand(resetIPCmd)
	rootCmd.AddCommand(rootResetCmd)
}
