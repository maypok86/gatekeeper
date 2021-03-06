package grpc

import (
	"context"

	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/maypok86/gatekeeper/internal/config"
	"github.com/maypok86/gatekeeper/internal/service"
	"github.com/maypok86/gatekeeper/pkg/password"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service struct {
	apipb.UnimplementedGatekeeperServiceServer
	loginService     service.Bucket
	passwordService  service.Bucket
	ipService        service.Bucket
	whitelistService service.IP
	blacklistService service.IP
}

func NewService(ctx context.Context) (*Service, error) {
	cfg := config.Get()

	loginService, err := service.NewBucket(ctx, cfg, cfg.RateLogin)
	if err != nil {
		return nil, err
	}

	passwordService, err := service.NewBucket(ctx, cfg, cfg.RatePassword)
	if err != nil {
		return nil, err
	}

	ipService, err := service.NewBucket(ctx, cfg, cfg.RateIP)
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

	return &Service{
		loginService:     loginService,
		passwordService:  passwordService,
		ipService:        ipService,
		whitelistService: whitelistService,
		blacklistService: blacklistService,
	}, nil
}

func (s *Service) Allow(ctx context.Context, request *apipb.AllowRequest) (*apipb.AllowResponse, error) {
	login := request.GetLogin()
	pwd, err := password.HashPassword(request.GetPassword())
	if err != nil {
		return &apipb.AllowResponse{Ok: false}, err
	}
	ip := request.GetIp()

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

func (s *Service) ResetLogin(ctx context.Context, request *apipb.ResetLoginRequest) (*apipb.ResetLoginResponse, error) {
	if login := request.GetLogin(); s.loginService.HasBucket(ctx, login) {
		zap.L().Info("Reset", zap.String("login", login))
		s.loginService.Remove(ctx, login)
		return &apipb.ResetLoginResponse{Ok: true}, nil
	}
	return &apipb.ResetLoginResponse{Ok: false}, errors.New("Not found login")
}

func (s *Service) ResetIP(ctx context.Context, request *apipb.ResetIPRequest) (*apipb.ResetIPResponse, error) {
	if ip := request.GetIp(); s.ipService.HasBucket(ctx, ip) {
		zap.L().Info("Reset", zap.String("ip", ip))
		s.ipService.Remove(ctx, ip)
		return &apipb.ResetIPResponse{Ok: true}, nil
	}
	return &apipb.ResetIPResponse{Ok: false}, errors.New("Not found ip")
}

func (s *Service) WhitelistAdd(ctx context.Context, request *apipb.WhitelistAddRequest) (*apipb.WhitelistAddResponse, error) {
	subnet := request.GetSubnet()
	if err := s.whitelistService.Add(ctx, subnet); err != nil {
		return &apipb.WhitelistAddResponse{Ok: false}, err
	}
	zap.L().Info("Add to whitelist", zap.String("subnet", subnet))
	return &apipb.WhitelistAddResponse{Ok: true}, nil
}

func (s *Service) WhitelistRemove(ctx context.Context, request *apipb.WhitelistRemoveRequest) (*apipb.WhitelistRemoveResponse, error) {
	subnet := request.GetSubnet()
	s.whitelistService.Remove(ctx, subnet)
	zap.L().Info("Remove from whitelist", zap.String("subnet", subnet))
	return &apipb.WhitelistRemoveResponse{Ok: true}, nil
}

func (s *Service) BlacklistAdd(ctx context.Context, request *apipb.BlacklistAddRequest) (*apipb.BlacklistAddResponse, error) {
	subnet := request.GetSubnet()
	if err := s.blacklistService.Add(ctx, subnet); err != nil {
		return &apipb.BlacklistAddResponse{Ok: false}, err
	}
	zap.L().Info("Add to blacklist", zap.String("subnet", subnet))
	return &apipb.BlacklistAddResponse{Ok: true}, nil
}

func (s *Service) BlacklistRemove(ctx context.Context, request *apipb.BlacklistRemoveRequest) (*apipb.BlacklistRemoveResponse, error) {
	subnet := request.GetSubnet()
	s.blacklistService.Remove(ctx, subnet)
	zap.L().Info("Remove from blacklist", zap.String("subnet", subnet))
	return &apipb.BlacklistRemoveResponse{Ok: true}, nil
}
