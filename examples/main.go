package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/event"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
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
		Host:     "176.57.153.41",
		Port:     "28116",
		Password: "nixzuessen",
	}
	// cfg := rcon.ServerConfig{
	// 	Host:     "208.115.226.2",
	// 	Port:     "7819",
	// 	Password: "61534hrehdfh",
	// }

	rcn, err := rcon.NewRcon(cfg, workerCount)
	if err != nil {
		logger.Fatal(err)
	}

	serverName, err := rcn.GetServerName()
	if err == nil {
		fmt.Printf("Conntected to the Server: %s\n", serverName)
	} else {
		fmt.Println(err)
	}

	infoCache := event.NewCache()
	eventListener := event.NewEventListener(rcn, infoCache)
	// printer := Printer{}
	//
	// eventListener.Register(&printer)

	time.Sleep(5 * time.Second)

	sv := infoCache.GetServerView()
	fmt.Printf("%#v\n", sv)
	// Open or create the file
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Convert the structure to JSON with indentation for readability
	jsonData, err := json.MarshalIndent(sv, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// playerinfo, _ := rcn.GetPlayerInfo("EV3NT")
	// fmt.Printf("%#v\n", playerinfo)

	// objs, err := rcn.GetCurrentMapObjectives()
	// if err == nil {
	// 	for _, row := range objs {
	// 		fmt.Println(strings.Join(row, " | "))
	// 	}
	// }

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	eventListener.Close()
	rcn.Close()
}
