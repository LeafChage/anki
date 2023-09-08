package google_tts

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/texttospeech/v1"
)

// ref: https://pkg.go.dev/google.golang.org/api@v0.138.0/texttospeech/v1
type googleTTS struct {
	service *texttospeech.Service
}

func GoogleTTS(ctx context.Context, credential *google.Credentials) (*googleTTS, error) {
	service, err := texttospeech.NewService(ctx, option.WithCredentials(credential))
	if err != nil {
		return nil, err
	}
	return &googleTTS{service}, nil
}
