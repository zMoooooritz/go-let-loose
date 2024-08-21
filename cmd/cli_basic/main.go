package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

var (
	logsCmd = "showlog"

	indexedListCommands = []string{
		"get mapsforrotation",
		"get players",
		"get playerids",
		"get adminids",
		"get admingrups",
		"get vipids",
		"get tempbans",
		"get permabans",
		"get profanities",
		"banprofanity",
		"unbanprofanity",
	}

	unindexedListCommands = []string{
		logsCmd,
	}
)

func main() {
	cfg := promptCredentials()
	rcn, err := rcon.NewRcon(cfg, 1)
	if err != nil {
		fmt.Println("Unable to establish connection to the server")
		return
	} else {
		fmt.Println("Successfully connected")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")
		response := executeCommand(rcn, cmd)
		for _, line := range response {
			fmt.Println(line)
		}
	}
}

func executeCommand(rcn *rcon.Rcon, cmd string) []string {
	responseFormat := config.RF_DIRECT
	for _, listCmd := range indexedListCommands {
		if strings.HasPrefix(strings.ToLower(cmd), listCmd) {
			responseFormat = config.RF_INDEXEDLIST
		}
	}
	for _, listCmd := range unindexedListCommands {
		if strings.HasPrefix(strings.ToLower(cmd), listCmd) {
			responseFormat = config.RF_UNINDEXEDLIST
		}
	}

	response := []string{}
	err := func() error { return nil }()

	switch responseFormat {
	case config.RF_DIRECT:
		fallthrough
	case config.RF_INDEXEDLIST:
		response, err = rcn.RunCommand(cmd, responseFormat)
	case config.RF_UNINDEXEDLIST:
		if strings.HasPrefix(strings.ToLower(cmd), logsCmd) {
			cmd, _ = strings.CutPrefix(strings.ToLower(cmd), logsCmd)
			minutes := util.ToInt(strings.TrimSpace(cmd))
			response, err = rcn.GetLogs(minutes)
		}
	}

	if err != nil {
		response = []string{fmt.Sprintln(err)}
	}

	return response
}

func promptCredentials() rcon.ServerConfig {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter hostname: ")
	host, _ := reader.ReadString('\n')
	host = host[:len(host)-1]

	fmt.Print("Enter port: ")
	port, _ := reader.ReadString('\n')
	port = port[:len(port)-1]

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1]

	return rcon.ServerConfig{
		Host:     host,
		Port:     port,
		Password: password,
	}
}
