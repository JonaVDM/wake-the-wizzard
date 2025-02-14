package cmd

import (
	"fmt"

	"github.com/JonaVDM/wake-the-wizzard/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "open a webserver to control stuff",

	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		return server.Serve(fmt.Sprintf("%s:%d", host, port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 3080, "Port on which the server runs on")
	serveCmd.Flags().String("host", "0.0.0.0", "Ip address to bind to")
}
