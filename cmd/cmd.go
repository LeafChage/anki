package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "anki",
		Long: "anki utility",
	}
	credentialFilePath string
)

func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&credentialFilePath, "file", "f", "~/.anki/google_credential.json", "google credential file")
	// rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(langCmd)
	rootCmd.AddCommand(voiceCmd)
	rootCmd.AddCommand(speakCmd)

	return rootCmd.Execute()
}
