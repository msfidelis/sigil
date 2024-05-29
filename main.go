package main

import (
	"main/cmd"

	"github.com/spf13/cobra"
)

type connection struct {
	protocol    string
	port        string
	processName string
	pid         string
}

func main() {

	var rootCmd = &cobra.Command{Use: "sigil"}

	rootCmd.AddCommand(
		cmd.ListeningCmd,
	)

	cmd.ListeningCmd.Flags().Int32P("port", "p", 0, "Show only active connections on specified port")
	cmd.ListeningCmd.Flags().BoolP("tcp", "t", true, "Show ports that are accepting inbound TCP connections")
	cmd.ListeningCmd.Flags().BoolP("udp", "u", false, "Show ports that are accepting inbound UDP connections")
	cmd.ListeningCmd.Flags().BoolP("all", "a", false, "Show ports that are accepting inbound connections, all of TCP and UDP")

	rootCmd.Execute()
}
