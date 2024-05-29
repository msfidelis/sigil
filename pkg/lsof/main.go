package lsof

import (
	"bufio"
	"main/models"
	"os/exec"
	"strings"
)

func GetTCPConnections(port int32) ([]models.Connection, error) {
	cmd := exec.Command("lsof", "-iTCP", "-P", "-n", "-sTCP:LISTEN")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	connections, err := buildConnectionInfo(output)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

func GetUDPConnections(port int32) ([]models.Connection, error) {
	cmd := exec.Command("lsof", "-iUDP", "-P", "-n")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	connections, err := buildConnectionInfo(output)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

func buildConnectionInfo(buff []byte) ([]models.Connection, error) {

	var connections []models.Connection

	scanner := bufio.NewScanner(strings.NewReader(string(buff)))
	scanner.Scan() // Skip da primeira linha

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		protocol := fields[7]
		address := fields[8]
		addressParts := strings.Split(address, ":")
		port := addressParts[len(addressParts)-1]

		if port != "*" && !strings.Contains(address, "->") {
			connections = append(connections, models.Connection{
				Protocol:    protocol,
				Port:        port,
				ProcessName: fields[0],
				Pid:         fields[1],
			})
		}
	}

	return connections, nil
}
