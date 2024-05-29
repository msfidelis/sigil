package tables

import (
	"fmt"
	"main/models"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ShowConnectionsTable(connections []models.Connection) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PID", "Port", "Process", "Protocol"})

	for _, conn := range connections {
		table.Append([]string{conn.Pid, conn.Port, conn.ProcessName, conn.Protocol})
	}
	fmt.Println()
	table.SetRowLine(true)
	table.Render()
}
