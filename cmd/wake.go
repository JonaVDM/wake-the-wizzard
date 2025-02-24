package cmd

import (
	"fmt"

	"github.com/JonaVDM/wake-the-wizzard/wol"
	"github.com/spf13/cobra"
)

// wakeCmd represents the wake command
var wakeCmd = &cobra.Command{
	Use:   "wake",
	Short: "Just a quick wake",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("wake called", args[0])
		return wol.SendWol(args[0])
	},
}

func init() {
	rootCmd.AddCommand(wakeCmd)
}
