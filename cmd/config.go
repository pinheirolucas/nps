package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"configure", "init", "setup"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// ps3Prompt := promptui.Prompt{
		// 	Label:   "Which PS3 firmware version are you using?",
		// 	Default: "4.89",
		// }

		// pspPrompt := promptui.Prompt{
		// 	Label:   "Which PS Vita firmware version are you using?",
		// 	Default: "6.61",
		// }

		// psvPrompt := promptui.Prompt{
		// 	Label:   "Which PSP firmware version are you using?",
		// 	Default: "3.74",
		// }

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
