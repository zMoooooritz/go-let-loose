
go-let-loose
======

[![Go](https://img.shields.io/badge/Go-blue.svg?style=for-the-badge&logo=go)](https://go.dev/)
[![Latest Release](https://img.shields.io/github/release/zMoooooritz/go-let-loose.svg?style=for-the-badge)](https://github.com/zMoooooritz/go-let-loose/releases)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/zMoooooritz/go-let-loose)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/zMoooooritz/go-let-loose/build.yml?branch=master&style=for-the-badge)](https://github.com/zMoooooritz/go-let-loose/actions)
[![Go ReportCard](https://goreportcard.com/badge/github.com/zMoooooritz/go-let-loose?style=for-the-badge)](https://goreportcard.com/report/zMoooooritz/go-let-loose)


Go bindings and interface for the remote console of Hell Let Loose

## ⇁ Features

- Bindings for all supported RCON opeartions
- Proper typing for all maps, armaments, squads and more
- Event system that expands on the provided logs

## ⇁ Installation 
```bash
go get github.com/zMoooooritz/go-let-loose
```

## ⇁ Usage
```go
const workerCount = 10

type Printer struct{}

func (p *Printer) IsSubscribed(e event.Event) bool {
	return true
}

func (p *Printer) Notify(e event.Event) {
	fmt.Println(e)
}

func main() {
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
```

## ⇁ Built with
- [ttlcache](https://github.com/jellydator/ttlcache)
- TUI [bubbletea](https://github.com/charmbracelet/bubbletea) and its awesome ecosystem

