package cmd

import (
	"fmt"
	"log"

	"github.com/JonaVDM/wake-the-wizzard/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "open a webserver to control stuff",

	RunE: func(cmd *cobra.Command, args []string) error {
		tsnet, err := cmd.Flags().GetBool("tsnet")
		if err != nil {
			return err
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		listen := fmt.Sprintf("%s:%d", host, port)
		if !tsnet {
			return server.Serve(listen)
		}

		s, err := tsConnect(cmd)
		if err != nil {
			return err
		}
		defer s.Close()

		ln, err := s.Listen("tcp", listen)
		if err != nil {
			log.Fatal(err)
		}
		defer ln.Close()

		return server.ServeListner(ln)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 3080, "Port on which the server runs on")
	serveCmd.Flags().String("host", "", "Ip address to bind to")
}
