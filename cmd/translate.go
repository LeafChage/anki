package cmd

import (
	"anki/deepl"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

type translateCmdOption struct {
	apiKeyFile string
	text       string
	target     string
	source     string
}

var (
	translateCmdOpt = translateCmdOption{}
	translateCmd    = &cobra.Command{
		Use: "translate",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			translateCmdOpt.text = strings.Join(args, "")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return translateFn(translateCmdOpt)
		},
	}
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	voiceCmd.Flags().StringVar(&translateCmdOpt.apiKeyFile, "f", path.Join(home, ".anki", "deepl.json"), "deepL api key file")
	voiceCmd.Flags().StringVar(&translateCmdOpt.target, "target", "Required", "target language ID")
	voiceCmd.Flags().StringVar(&translateCmdOpt.source, "source", "NULL", "source language ID")
	voiceCmd.MarkFlagRequired("target")
}

func apiKey(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	apiKey := strings.TrimSuffix(string(f), "\n")
	return apiKey, nil
}

func translateFn(option translateCmdOption) error {
	key, err := apiKey(option.apiKeyFile)
	if err != nil {
		return err
	}

	client := deepl.DeepLClient(key)

	options := []deepl.TranslateRequest{}
	if option.source != "NULL" {
		options = append(options, deepl.SourceLang(option.source))
	}

	res, err := client.Translate(
		deepl.TranslateText(option.text),
		deepl.TargetLang(option.target),
		options...,
	)
	if err != nil {
		return err
	}

	for _, val := range res.Translations {
		fmt.Println(val.Text)
	}
	return nil
}
