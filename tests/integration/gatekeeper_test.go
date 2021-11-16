package integration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cucumber/godog"
	apipb "github.com/maypok86/gatekeeper/internal/api/grpc/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var grpcAddress = os.Getenv("GRPC_ADDRESS")

func init() {
	if grpcAddress == "" {
		grpcAddress = "localhost:50051"
	}
}

type apiTest struct {
	login         string
	password      string
	ip            string
	subnet        string
	responseError error
}

func (at *apiTest) loginIs(login string) error {
	at.login = login
	return nil
}

func (at *apiTest) passwordIs(pass string) error {
	at.password = pass
	return nil
}

func (at *apiTest) ipIs(ip string) error {
	at.ip = ip
	return nil
}

func (at *apiTest) subnetIs(subnet string) error {
	at.subnet = subnet
	return nil
}

func (at *apiTest) iCallGrpcMethod(method string) error {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := apipb.NewGatekeeperServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	switch method {
	case "Allow":
		_, err = client.Allow(ctx, &apipb.AllowRequest{
			Login:    at.login,
			Password: at.password,
			Ip:       at.ip,
		})
		at.responseError = err
	case "ResetLogin":
		_, err = client.ResetLogin(ctx, &apipb.ResetLoginRequest{
			Login: at.login,
		})
		at.responseError = err
	case "ResetIP":
		_, err = client.ResetIP(ctx, &apipb.ResetIPRequest{
			Ip: at.ip,
		})
		at.responseError = err
	case "BlacklistAdd":
		_, err = client.BlacklistAdd(ctx, &apipb.BlacklistAddRequest{
			Subnet: at.subnet,
		})
		at.responseError = err
	case "BlacklistRemove":
		_, err = client.BlacklistRemove(ctx, &apipb.BlacklistRemoveRequest{
			Subnet: at.subnet,
		})
		at.responseError = err
	case "WhitelistAdd":
		_, err = client.WhitelistAdd(ctx, &apipb.WhitelistAddRequest{
			Subnet: at.subnet,
		})
		at.responseError = err
	case "WhitelistRemove":
		_, err = client.WhitelistRemove(ctx, &apipb.WhitelistRemoveRequest{
			Subnet: at.subnet,
		})
		at.responseError = err
	default:
		return errors.New("unexpected method: " + method)
	}

	return nil
}

func (at *apiTest) responseErrorShouldBe(err string) error {
	const errNil = "nil"
	if err != errNil {
		err = "rpc err: code = Unknown desc = " + err
	}
	if err == errNil && at.responseError != nil {
		return fmt.Errorf("unexpected err, expected %s, got %w", err, at.responseError)
	}
	if err != errNil && at.responseError == nil {
		return fmt.Errorf("unexpected err, expected %s, got %w", err, nil)
	}
	if at.responseError != nil && err != at.responseError.Error() {
		return fmt.Errorf("unexpected err, expected %s, got %w", err, at.responseError)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	var test apiTest
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		test.login = ""
		test.password = ""
		test.ip = ""
		test.responseError = nil

		return ctx, nil
	})

	ctx.Step(`^login is "([^"]*)"$`, test.loginIs)
	ctx.Step(`^password is "([^"]*)"$`, test.passwordIs)
	ctx.Step(`^ip is "([^"]*)"$`, test.ipIs)
	ctx.Step(`^subnet is "([^"]*)"$`, test.subnetIs)

	ctx.Step(`^I call grpc method "([^"]*)"$`, test.iCallGrpcMethod)
	ctx.Step(`^response error should be "([^"]*)"$`, test.responseErrorShouldBe)
}
