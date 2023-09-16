package cmd

import (
	"anki/google"
	"anki/google/google_tts"
	"context"
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/texttospeech/v1"
)

type speakCommandOption struct {
	languageCode string
	text         string
	gender       string
	voice        string
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
			if speakCmdOpt.voice != "NULL" && speakCmdOpt.gender != "NULL" {
				return errors.New("don't specify both gender and voice")
			}
			return speakFn(credentialFilePath, speakCmdOpt)
		},
	}
)

func init() {
	speakCmd.Flags().StringVar(&speakCmdOpt.languageCode, "lang", "Required", "speaker speaks in this language")
	speakCmd.Flags().StringVar(&speakCmdOpt.gender, "gender", "NULL", "speaker is this gender")
	speakCmd.Flags().StringVar(&speakCmdOpt.voice, "voice", "NULL", "speaker is this")
	speakCmd.MarkFlagRequired("lang")
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

	options := []google_tts.SynthesizeOptions{google_tts.SynthesisText(speakCmdOpt.text)}
	if speakCmdOpt.voice != "NULL" {
		options = append(options, google_tts.SynthesisVoiceName(speakCmdOpt.voice))
	}
	if speakCmdOpt.gender != "NULL" {
		options = append(options, google_tts.SSMGender(speakCmdOpt.gender))
	}
	if speakCmdOpt.languageCode != "NULL" {
		options = append(options, google_tts.SynthesisLanguageCode(speakCmdOpt.languageCode))
	}

	err = client.Synthesize(ctx, options...)
	if err != nil {
		return err
	}
	return nil
}
