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

- Bindings for all supported RCONv2 operations.
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

---

## 🚨 Warning

> [!WARNING]
> **Important API Changes:**
>
> - **Library update required**: Since the official implementation of the legacy RCON protocol got removed, everyone needs to update to the new library version
> - **Module rename**: The `rconv2` implementation has been renamed to the `rcon` module
>   - Users currently using `rconv2` will need to update their import paths and code
>   - Users who remained on the original `rcon` should experience minimal disruption
> - **Unified event system**: The event system has been consolidated into the `rcon` module, providing a more elegant, uniform, and easy-to-use interface

---

## 📖 Usage

### Using the `rcon` API

Below is an example of how to use the `go-let-loose/rcon` module in a Go project:

```go
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
    fmt.Printf("Connected to the Server: %s\n", serverName)
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
    onKill(function(event)
      print("Kill: " .. event.Killer.Name .. " -> " .. event.Victim.Name .. " (" .. event.Weapon.Name .. ")")
    end)
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
- [gopher-lua](https://github.com/yuin/gopher-lua)

---

## 📄 License

This project is licensed under the **MIT License**. You can view the full license [here](https://github.com/zMoooooritz/go-let-loose/blob/master/LICENSE).
