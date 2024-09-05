package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zMoooooritz/go-let-loose/internal/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/event"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

const workerCount = 10

type Printer struct{}

func (p *Printer) IsSubscribed(e event.Event) bool {
	return true
}

func (p *Printer) Notify(e event.Event) {
	fmt.Println(e)
}

func main() {
	logger.DefaultLogger()

	cfg := rcon.ServerConfig{
		Host:     "123.123.123.123",
		Port:     "12345",
		Password: "password",
	}

	rcn, err := rcon.NewRcon(cfg, workerCount)
	if err != nil {
		log.Fatal(err)
	}

	serverName, err := rcn.GetServerName()
	if err == nil {
		fmt.Printf("Conntected to the Server: %s\n", serverName)
	} else {
		fmt.Println(err)
	}

	infoCache := event.NewCache()
	eventListener := event.NewEventListener(rcn, infoCache)
	printer := Printer{}

	eventListener.Register(&printer)

	time.Sleep(time.Second)

	fmt.Printf("%+v\n", infoCache.GetGameState())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	eventListener.Close()
	rcn.Close()
}
