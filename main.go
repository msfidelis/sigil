package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type connection struct {
	protocol    string
	port        string
	processName string
	pid         string
}

func main() {
	port := flag.Int("port", 0, "target port, default: 0")
	flag.Parse()

	if *port != 0 {
		return
	}

	connections, err := getActiveConnections()
	if err != nil {
		log.Fatal(err)
	}

	output(connections)
}

func getActiveConnections() ([]connection, error) {
	output, err := lsofListeningTCP()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	scanner.Scan()

	var connections []connection

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		protocol := fields[7]
		address := fields[8]
		addressParts := strings.Split(address, ":")
		port := addressParts[len(addressParts)-1]

		connections = append(connections, connection{
			protocol:    protocol,
			port:        port,
			processName: fields[0],
			pid:         fields[1],
		})
	}

	return connections, nil
}

func output(connections []connection) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PID", "Port", "Process", "Protocol"})

	for _, conn := range connections {
		table.Append([]string{conn.pid, conn.port, conn.processName, conn.protocol})
	}

	table.SetRowLine(true)
	table.Render()
}

func lsofListeningTCP() ([]byte, error) {
	cmd := exec.Command("lsof", "-iTCP", "-P", "-n", "-sTCP:LISTEN")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}
