package cmd

import (
	"anki/google"
	"anki/google/google_tts"
	"context"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/texttospeech/v1"
)

type speakCommandOption struct {
	languageCode string
	text         string
}

var (
	speakCmdOpt = speakCommandOption{}
	speakCmd    = &cobra.Command{
		Use: "speak",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			speakCmdOpt.text = strings.Join(args, "")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return speakFn(credentialFilePath, speakCmdOpt)
		},
	}
)

func init() {
	speakCmd.Flags().StringVar(&speakCmdOpt.languageCode, "lang", "ja-JP", "output sound speaked person in this language")
}

func speakFn(credential string, option speakCommandOption) error {
	ctx := context.Background()
	c, err := google.AuthFromServiceAccount(ctx, credential, texttospeech.CloudPlatformScope)
	if err != nil {
		return err
	}

	client, err := google_tts.GoogleTTS(ctx, c)
	if err != nil {
		return err
	}

	err = client.Synthesize(
		ctx,
		google_tts.SynthesisText(option.text),
		google_tts.SynthesisLanguageCode(option.languageCode),
	)
	if err != nil {
		return err
	}
	return nil
}
