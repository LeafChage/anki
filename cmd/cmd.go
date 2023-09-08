package cmd

import (
	"os"
	"path"

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
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	rootCmd.PersistentFlags().StringVarP(&credentialFilePath, "file", "f", path.Join(home, ".anki", "google_credential.json"), "google credential file")
	// rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(langCmd)
	rootCmd.AddCommand(voiceCmd)
	rootCmd.AddCommand(speakCmd)

	return rootCmd.Execute()
}
