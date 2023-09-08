package cmd

import (
	"anki/google"
	"anki/google/google_tts"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/texttospeech/v1"
)

var langCmd = &cobra.Command{
	Use: "lang",
	RunE: func(cmd *cobra.Command, args []string) error {
		return langFn(credentialFilePath)
	},
}

func langFn(credential string) error {
	ctx := context.Background()
	c, err := google.AuthFromServiceAccount(context.TODO(), credential, texttospeech.CloudPlatformScope)
	if err != nil {
		return err
	}

	client, err := google_tts.GoogleTTS(ctx, c)
	if err != nil {
		return err
	}

	langs, err := client.LanguageCodes(ctx)
	if err != nil {
		return err
	}

	for _, lang := range langs {
		fmt.Println(lang)
	}
	return nil
}
