package grpc

import (
	"context"
	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/maypok86/gatekeeper/pkg/password"
	"go.uber.org/zap"

	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/service"
)

type Server struct {
	loginService     service.Bucket
	passwordService  service.Bucket
	ipService        service.Bucket
	whitelistService service.IP
	blacklistService service.IP
}

func NewServer(ctx context.Context) (*Server, error) {
	cfg := config.Get()

	loginService, err := service.NewBucket(ctx, cfg)
	if err != nil {
		return nil, err
	}

	passwordService, err := service.NewBucket(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ipService, err := service.NewBucket(ctx, cfg)
	if err != nil {
		return nil, err
	}

	whitelistService, err := service.NewIP(ctx, cfg)
	if err != nil {
		return nil, err
	}

	blacklistService, err := service.NewIP(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		loginService:     loginService,
		passwordService:  passwordService,
		ipService:        ipService,
		whitelistService: whitelistService,
		blacklistService: blacklistService,
	}, nil
}

func (s *Server) Allow(ctx context.Context, request *apipb.AllowRequest) (*apipb.AllowResponse, error) {
	login := request.Login
	pwd, err := password.HashPassword(request.Password)
	if err != nil {
		return &apipb.AllowResponse{Ok: false}, err
	}
	ip := request.Ip

	isAllow, err := s.whitelistService.Contains(ctx, ip)
	if err != nil {
		return &apipb.AllowResponse{Ok: false}, err
	}
	if isAllow {
		return &apipb.AllowResponse{Ok: true}, nil
	}

	disallow, err := s.blacklistService.Contains(ctx, ip)
	if err != nil {
		return &apipb.AllowResponse{Ok: false}, err
	}
	if disallow {
		zap.L().Info("Block allow", zap.String("login", login), zap.String("ip", ip))
		return &apipb.AllowResponse{Ok: false}, nil
	}

	if s.loginService.Allow(ctx, login) && s.passwordService.Allow(ctx, pwd) && s.ipService.Allow(ctx, ip) {
		return &apipb.AllowResponse{Ok: true}, nil
	}

	zap.L().Info("Block allow", zap.String("login", login), zap.String("ip", ip))
	return &apipb.AllowResponse{Ok: false}, nil
}
