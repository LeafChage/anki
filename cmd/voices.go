package cmd

import (
	"anki/google"
	"anki/google/google_tts"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/api/texttospeech/v1"
)

type voiceCommandOption struct {
	languageCode string
}

var (
	voiceCmdOpt = voiceCommandOption{}
	voiceCmd    = &cobra.Command{
		Use: "voice",
		RunE: func(cmd *cobra.Command, args []string) error {
			return voiceFn(credentialFilePath, voiceCmdOpt.languageCode)
		},
	}
)

func init() {
	voiceCmd.Flags().StringVar(&voiceCmdOpt.languageCode, "lang", "ja-JP", "filter by this language code")
}

func voiceFn(credential string, langCode string) error {
	ctx := context.Background()
	c, err := google.AuthFromServiceAccount(ctx, credential, texttospeech.CloudPlatformScope)
	if err != nil {
		return err
	}

	client, err := google_tts.GoogleTTS(ctx, c)
	if err != nil {
		return err
	}

	codes, err := client.VoiceName(
		ctx,
		google_tts.LanguageCode(langCode),
	)
	if err != nil {
		return err
	}

	for _, code := range codes {
		fmt.Println(code)
	}
	return nil
}
