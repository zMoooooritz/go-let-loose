package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zMoooooritz/go-let-loose/internal/lua"
	"github.com/zMoooooritz/go-let-loose/pkg/event"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
)

const workerCount = 10

func main() {
	logger.NOPLogger()

	var cfg rcon.ServerConfig

	flag.StringVar(&cfg.Host, "host", "", "hostname of server")
	flag.StringVar(&cfg.Port, "port", "", "port on the server")
	flag.StringVar(&cfg.Password, "password", "", "password of the rcon")
	flag.Parse()

	rcn, err := rcon.NewRcon(cfg, workerCount)
	if err != nil {
		logger.Fatal(err)
		os.Exit(0)
	}

	cache := event.NewCache()
	eventListener := event.NewEventListener(rcn, cache)

	lua.InitLua(rcn, cache, eventListener)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	lua.DeinitLua()

	eventListener.Close()
	rcn.Close()
}
