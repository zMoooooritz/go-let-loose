go-let-loose
======

[![Go](https://img.shields.io/badge/Go-blue.svg?style=for-the-badge&logo=go)](https://go.dev/)
[![Latest Release](https://img.shields.io/github/release/zMoooooritz/go-let-loose.svg?style=for-the-badge)](https://github.com/zMoooooritz/go-let-loose/releases)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://pkg.go.dev/github.com/zMoooooritz/go-let-loose)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/zMoooooritz/go-let-loose/build.yml?branch=master&style=for-the-badge)](https://github.com/zMoooooritz/go-let-loose/actions)
[![Go ReportCard](https://goreportcard.com/badge/github.com/zMoooooritz/go-let-loose?style=for-the-badge)](https://goreportcard.com/report/zMoooooritz/go-let-loose)

Go bindings and interface for the remote console of **Hell Let Loose**.

---

## 🚀 Features

- Bindings for all supported RCON operations.
- Proper typing for maps, armaments, squads, and more.
- Event system that expands on the HLL provided logs.
- Support for Lua plugins.

---

## 📦 Installation

### Install the Go package:

To install the Go package, simply run:

```bash
go get github.com/zMoooooritz/go-let-loose
```

### Install the CLI:

For those who want to install and use the **CLI** version of `go-let-loose`, download the latest release from the [releases page](https://github.com/zMoooooritz/go-let-loose/releases) or use the following command:

```bash
go install github.com/zMoooooritz/go-let-loose/cmd/go-let-loose-cli@latest
```

This will install `go-let-loose-cli` globally on your system, allowing you to interact with the HLL server directly from the command line.

---

## 📖 Usage

Below is an example of how to use the `go-let-loose` module in a Go project:

```go
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

    const workerCount = 10
	rcn, err := rcon.NewRcon(cfg, workerCount)
	if err != nil {
		log.Fatal(err)
	}

	serverName, err := rcn.GetServerName()
	if err == nil {
		fmt.Printf("Connected to the Server: %s\n", serverName)
	} else {
		fmt.Println(err)
	}

	printer := Printer{}
	infoCache := event.NewCache()
	eventListener := event.NewEventListener(rcn, infoCache)
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

### Lua Plugins

To use Lua plugins, place your Lua scripts in the `plugins` directory. The system will automatically detect and load them.

Example:

```lua
function Init()
	print("Initializing Plugin")
end

function Run()
	local err, name = getServerName()
	if name then
		print("Connected to the Server: " .. name)
	else
		print("Error: " .. err)
	end
end

function Stop()
	print("Exiting Plugin")
	exit()
end
```

For a more in-depth example, see [hello_world.lua](https://github.com/zMoooooritz/go-let-loose/blob/master/plugins/hello_world.lua).

To run the plugins, use the following command:

```bash
go run cmd/go-let-loose-lua/main.go
```

---

## 🔧 Built with

- [ttlcache](https://github.com/jellydator/ttlcache)
- TUI powered by [bubbletea](https://github.com/charmbracelet/bubbletea) and its awesome ecosystem.

---

## 📄 License

This project is licensed under the **MIT License**. You can view the full license [here](https://github.com/zMoooooritz/go-let-loose/blob/master/LICENSE).

