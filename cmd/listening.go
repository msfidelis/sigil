package cmd

import (
	"log"
	"main/pkg/lsof"
	"main/pkg/tables"

	"github.com/spf13/cobra"
)

var ListeningCmd = &cobra.Command{
	Use:   "listening",
	Short: "Show host ports that are active and accepting connections",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tcpFlag, _ := cmd.Flags().GetBool("tcp")
		udpFlag, _ := cmd.Flags().GetBool("udp")

		if udpFlag {
			oup, err := lsof.GetUDPConnections(0)
			if err != nil {
				log.Fatal(err)
			}
			tables.ShowConnectionsTable(oup)
			return
		}

		if tcpFlag {
			otp, err := lsof.GetTCPConnections(0)
			if err != nil {
				log.Fatal(err)
			}
			tables.ShowConnectionsTable(otp)
		}

	},
}
