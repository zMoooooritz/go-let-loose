package rcon

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetServerName() (string, error) {
	data, err := r.runBasicCommand("get name")
	if err != nil {
		return "", err
	}
	return data, nil
}

func (r *Rcon) GetSlots() (int, int, error) {
	resp, err := r.runBasicCommand("get slots")
	if err != nil {
		return 0, 0, err
	}
	split := strings.Split(resp, "/")
	if len(split) < 2 {
		return 0, 0, nil
	}
	return util.ToInt(split[0]), util.ToInt(split[1]), nil
}

func (r *Rcon) GetGameState() (hll.GameState, error) {
	gameState := hll.GameState{}
	data, err := r.runBasicCommand("get gamestate")
	if err != nil {
		return gameState, err
	}
	return ParseGameState(data)
}

func ParseGameState(data string) (hll.GameState, error) {
	var gameState hll.GameState
	lines := strings.Split(data, config.NEWLINE)
	if len(lines) < 5 {
		return gameState, fmt.Errorf("invalid data %s", data)
	}
	r := regexp.MustCompile(`Players: Allied: (\d+) - Axis: (\d+)`)
	match := r.FindStringSubmatch(lines[0])
	if len(match) < 3 {
		return gameState, fmt.Errorf("invalid data %s", data)
	}
	gameState.PlayerCount = hll.TeamData{
		Allies: util.ToInt(match[1]),
		Axis:   util.ToInt(match[2]),
	}

	r = regexp.MustCompile(`Score: Allied: (\d+) - Axis: (\d+)`)
	match = r.FindStringSubmatch(lines[1])
	if len(match) < 3 {
		return gameState, fmt.Errorf("invalid data %s", data)
	}
	gameState.GameScore = hll.TeamData{
		Allies: util.ToInt(match[1]),
		Axis:   util.ToInt(match[2]),
	}

	r = regexp.MustCompile(`Remaining Time: (\d):(\d{2}):(\d{2})`)
	match = r.FindStringSubmatch(lines[2])
	if len(match) < 4 {
		return gameState, fmt.Errorf("invalid data %s", data)
	}
	gameState.RemainingSeconds = util.ToInt(match[1])*60*60 + util.ToInt(match[2])*60 + util.ToInt(match[3])

	currMapName := strings.Split(lines[3], ": ")[1]
	nextMapName := strings.Split(lines[4], ": ")[1]

	gameState.CurrentMap = hll.ParseLayer(currMapName)
	gameState.NextMap = hll.ParseLayer(nextMapName)
	return gameState, nil
}
func (r *Rcon) GetMaxQueuedPlayers() (int, error) {
	return getNumVal(r, "get maxqueuedplayers")
}

func (r *Rcon) GetNumVipSlots() (int, error) {
	return getNumVal(r, "get numvipslots")
}

func (r *Rcon) SetMaxQueuedPlayers(size int) error {
	return runSetCommand(r, fmt.Sprintf("setmaxqueuedplayers %d", min(0, max(size, 6))))
}

func (r *Rcon) SetNumVipSlots(amount int) error {
	return runSetCommand(r, fmt.Sprintf("setnumvipslots %d", min(0, max(amount, 100))))
}

func (r *Rcon) SetWelcomeMessage(message string) error {
	return runSetCommand(r, fmt.Sprintf("say %s", message))
}

func (r *Rcon) SetBroadcastMessage(message string) error {
	return runSetCommand(r, fmt.Sprintf("broadcast %s", message))
}

func (r *Rcon) ClearBroadcastMessage(message string) error {
	return r.SetBroadcastMessage("")
}

func (r *Rcon) GetLogs(spanMins int) ([]string, error) {
	return r.runUnindexedListCommand(fmt.Sprintf("showlog %d", spanMins))
}
