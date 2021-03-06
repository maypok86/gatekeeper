package cmd

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_service "github.com/maypok86/gatekeeper/internal/api/grpc"
	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start Gatekeeper api",
	Long:  `Start grpc api for Gatekeeper service`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := config.Get()
		logger.Configure(cfg.Logger)
		zap.L().Info("Logger init success")

		promRegistry := prometheus.NewRegistry()
		grpcMetrics := grpc_prometheus.NewServerMetrics()
		promRegistry.MustRegister(grpcMetrics)

		grpcServer := grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
				grpcMetrics.StreamServerInterceptor(),
				grpc_zap.StreamServerInterceptor(zap.L()),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				grpcMetrics.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zap.L()),
			)),
		)
		if cfg.IsDev() {
			reflection.Register(grpcServer)
		}

		grpcService, err := grpc_service.NewService(ctx)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to get grpc service %v", err))
		}
		apipb.RegisterGatekeeperServiceServer(grpcServer, grpcService)

		errChan := make(chan error)
		quitChan := make(chan os.Signal, 1)

		grpc_prometheus.Register(grpcServer)
		grpc_prometheus.EnableHandlingTimeHistogram()
		grpcMetrics.InitializeMetrics(grpcServer)
		zap.L().Info("Monitoring export listen ", zap.String("port", cfg.PrometheusPort))
		go func() {
			err = http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.PrometheusPort), promhttp.Handler())
			if err != nil {
				zap.L().Error(err.Error())
				errChan <- err
			}
			http.Handle("/metrics", promhttp.Handler())
		}()

		zap.L().Info("grpc server starting")
		go func() {
			lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
			if err != nil {
				zap.L().Error(err.Error())
				errChan <- err
			}
			if err := grpcServer.Serve(lis); err != nil {
				zap.L().Error(err.Error())
				errChan <- err
			}
		}()

		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-errChan:
		case <-quitChan:
		}
		grpcServer.GracefulStop()
	},
}
