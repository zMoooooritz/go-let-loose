package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

func onKill(e hll.KillEvent) {
	fmt.Printf("Kill: %s -> %s (%s)\n", e.Killer.Name, e.Victim.Name, e.Weapon.Name)
}

func main() {
	logger.DefaultLogger()

	cfg := rcon.ServerConfig{
		Host:     "123.123.123.123",
		Port:     "12345",
		Password: "password",
	}

	const workerCount = 10
	rcn, err := rcon.NewRcon(cfg, workerCount, rcon.WithCache(), rcon.WithEvents())
	if err != nil {
		log.Fatal(err)
	}

	serverName, err := rcn.GetServerName()
	if err == nil {
		fmt.Printf("Conntected to the Server: %s\n", serverName)
	} else {
		fmt.Println(err)
	}

	rcn.OnKill(onKill)

	time.Sleep(time.Second)

	gameState, err := rcn.GetGameState()
	if err == nil {
		fmt.Printf("Current game state: %+v\n", gameState)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	rcn.Close()
}
