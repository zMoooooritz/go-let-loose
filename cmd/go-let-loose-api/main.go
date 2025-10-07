package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

const workerCount = 10

func main() {
	logger.DefaultLogger()

	logger.Info("Starting go-let-loose-api...")

	var cfg rcon.ServerConfig

	flag.StringVar(&cfg.Host, "host", "", "hostname of server")
	flag.StringVar(&cfg.Port, "port", "", "port on the server")
	flag.StringVar(&cfg.Password, "password", "", "password of the rcon")
	flag.Parse()

	rcn, err := rcon.NewRcon(cfg, workerCount, rcon.WithVerification())
	if err != nil {
		logger.Fatal(err)
		os.Exit(0)
	}

	cmds, err := rcn.GetCommands()
	if err != nil {
		logger.Fatal(err)
		os.Exit(0)
	}

	cmdDetails := []hll.CommandDetails{}
	for _, c := range cmds {
		if c.ClientSupported {
			commandDetails, err := rcn.GetCommandDetails(c.ID)
			if err == nil {
				cmdDetails = append(cmdDetails, commandDetails)
			} else {
				cmdDetails = append(cmdDetails, hll.CommandDetails{
					Name:        c.ID,
					Text:        c.Name,
					Description: c.Name,
				})
			}
		} else {
			cmdDetails = append(cmdDetails, hll.CommandDetails{
				Name:        c.ID,
				Text:        c.Name,
				Description: c.Name,
			})
		}
	}

	f, err := os.Create("hll_api.json")
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cmdDetails); err != nil {
		logger.Fatal(err)
		_ = f.Close()
		os.Exit(1)
	}
	_ = f.Close()

	rcn.Close()
}
