package google

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/texttospeech/v1"
	"io/ioutil"
	"os"
)

type Scope = string

const GOOGLE_TTS Scope = texttospeech.CloudPlatformScope

func AuthFromServiceAccount(
	ctx context.Context,
	path string,
	scopes ...Scope,
) (*google.Credentials, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return google.CredentialsFromJSON(
		ctx,
		bytes,
		scopes...,
	)
}
