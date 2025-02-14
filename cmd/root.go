package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"tailscale.com/tsnet"
)

var rootCmd = &cobra.Command{
	Use:   "wol",
	Short: "a simple wake on lan application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("tsnet", false, "Run the app on a tailnet")
	rootCmd.PersistentFlags().String("tsnet-authkey", "", "The key to connect to the tailnet")
	rootCmd.PersistentFlags().String("tsnet-control", "", "The control server")
	rootCmd.PersistentFlags().String("tsnet-hostname", "wizzard", "The hostname on tailnet")
}

func tsConnect(cmd *cobra.Command) (*tsnet.Server, error) {
	s := new(tsnet.Server)

	control, err := cmd.Flags().GetString("tsnet-control")
	if err != nil {
		return nil, err
	}
	authkey, err := cmd.Flags().GetString("tsnet-authkey")
	if err != nil {
		return nil, err
	}
	hostname, err := cmd.Flags().GetString("tsnet-hostname")
	if err != nil {
		return nil, err
	}
	s.ControlURL = control
	s.AuthKey = authkey
	s.Hostname = hostname

	_, err = s.Up(context.TODO())
	if err != nil {
		return nil, err
	}

	return s, nil
}
