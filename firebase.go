package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type Service struct {
	*firebase.App
	*auth.Client
}

func (s *Service) Dial(ctx context.Context, cfg Config) error {
	var err error
	s.App, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(cfg.Credentials))

	if err != nil {
		return err
	}

	s.Client, err = s.App.Auth(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Close(ctx context.Context) error {
	return nil
}
