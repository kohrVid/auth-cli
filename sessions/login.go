package sessions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kohrVid/auth/proto"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

func Login(sessionParams map[string]string) string {
	err := validateSessionParams(sessionParams)
	if err != nil {
		log.Fatalf("invalid credentials: %v", err)
	}

	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	authClient := proto.NewAuthenticationServiceClient(conn)
	r := sessionAuth(authClient, sessionParams)
	return r.Result
}

func validateSessionParams(sessionParams map[string]string) error {
	var errs []error

	for k, v := range sessionParams {
		if v == "" {
			msg := fmt.Sprintf("missing %v", k)
			errs = append(errs, errors.New(msg))
		}

	}

	if len(errs) > 0 {
		return errs[0]
	} else {
		return nil
	}
}

func sessionAuth(authClient proto.AuthenticationServiceClient, sp map[string]string) *proto.AuthenticationResponse {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := authClient.CredentialCheck(
		ctx,
		&proto.AuthenticationRequest{
			Username: sp["username"],
			Password: sp["password"],
		},
	)

	if err != nil {
		log.Fatalf("could not log in user: %v", err)
	}
	return r
}
