package cmd

import (
	"context"
	"log"
	"time"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverURL string

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverURL, "url", "u", "localhost:50051", "Gatekeeper server url")
}

var rootCmd = &cobra.Command{
	Use:   "gk",
	Short: "Gatekeeper service and cli",
	Long:  `Gatekeeper is storage limiting service and cli tool`,
}

func runGrpcClient(
	getOk func(client apipb.GatekeeperServiceClient, ctx context.Context, arg string) (bool, error),
) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
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
		ok, err := getOk(client, ctx, args[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Success: %t", ok)
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
