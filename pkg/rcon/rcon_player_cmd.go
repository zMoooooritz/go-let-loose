package rcon

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zMoooooritz/go-let-loose/internal/util"
	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	caser = cases.Title(language.AmericanEnglish)
)

func (r *Rcon) GetPlayers() ([]hll.PlayerInfo, error) {
	players := []hll.PlayerInfo{}
	data, err := r.runListCommand("get playerids")
	if err != nil {
		return players, err
	}
	nameIdSeparator := " : "
	for _, entry := range data {
		splitIndex := strings.LastIndex(entry, nameIdSeparator)
		if splitIndex == -1 {
			continue
		}
		players = append(players, hll.PlayerInfo{
			Name: entry[:splitIndex],
			ID:   entry[splitIndex+len(nameIdSeparator):],
		})
	}
	return players, nil
}

func (r *Rcon) GetPlayerNames() ([]string, error) {
	data, err := r.runListCommand("get players")
	if err != nil {
		return []string{}, err
	}
	return data, nil
}

func (r *Rcon) GetPlayerIDs() ([]string, error) {
	playerIDs := []string{}
	players, err := r.GetPlayers()
	if err != nil {
		return playerIDs, err
	}
	for _, player := range players {
		playerIDs = append(playerIDs, player.ID)
	}
	return playerIDs, nil
}

func (r *Rcon) GetAdmins() ([]hll.Admin, error) {
	admins := []hll.Admin{}
	data, err := r.runListCommand("get adminids")
	if err != nil {
		return admins, err
	}
	for _, entry := range data {
		split := strings.SplitN(entry, " ", 3)
		if len(split) < 3 {
			continue
		}
		admins = append(admins, hll.Admin{
			PlayerInfo: hll.PlayerInfo{
				Name: split[2],
				ID:   split[0],
			},
			Role: hll.AdminRole(split[1]),
		})
	}
	return admins, nil
}

func (r *Rcon) GetAdminRoles() ([]hll.AdminRole, error) {
	roles := []hll.AdminRole{}
	data, err := r.runListCommand("get admingroups")
	if err != nil {
		return roles, err
	}
	for _, entry := range data {
		roles = append(roles, hll.AdminRole(entry))
	}
	return roles, nil
}

func (r *Rcon) GetVIPs() ([]hll.PlayerInfo, error) {
	players := []hll.PlayerInfo{}
	data, err := r.runListCommand("get vipids")
	if err != nil {
		return players, err
	}
	for _, entry := range data {
		split := strings.SplitN(entry, " ", 2)
		if len(split) == 2 {
			players = append(players, hll.PlayerInfo{
				Name: split[1],
				ID:   split[0],
			})
		}
	}
	return players, nil
}

func (r *Rcon) GetPlayerInfo(playerName string) (hll.DetailedPlayerInfo, error) {
	detailedPlayer := hll.DetailedPlayerInfo{}
	data, err := r.runBasicCommand("playerinfo " + playerName)
	if err != nil {
		return detailedPlayer, err
	}
	return ParsePlayerInfo(data)
}

func ParsePlayerInfo(data string) (hll.DetailedPlayerInfo, error) {
	var detailedPlayer hll.DetailedPlayerInfo

	lines := strings.Split(data, config.NEWLINE)
	valueMap := make(map[string]string)
	for _, line := range lines {
		split := strings.SplitN(line, ": ", 2)
		if len(split) >= 2 {
			valueMap[strings.ToLower(split[0])] = split[1]
		}
	}

	detailedPlayer.Name = valueMap["name"]
	detailedPlayer.ID = valueMap["steamid64"]
	detailedPlayer.Team = hll.TeamFromString(valueMap["team"])
	detailedPlayer.Role = hll.RoleFromString(valueMap["role"])
	if detailedPlayer.Role == hll.ArmyCommander {
		detailedPlayer.Unit = hll.CommandUnit
	} else {
		if value, ok := valueMap["unit"]; ok {
			unitSplit := strings.Split(value, " - ")
			if len(unitSplit) == 2 {
				detailedPlayer.Unit = hll.Unit{
					Name: caser.String(unitSplit[1]),
					ID:   util.ToInt(unitSplit[0]),
				}
			} else {
				return detailedPlayer, errors.New("unit data invalid")
			}
		} else {
			detailedPlayer.Unit = hll.NoUnit
		}
	}

	if value, ok := valueMap["loadout"]; ok {
		detailedPlayer.Loadout = value
	} else {
		detailedPlayer.Loadout = "none"
	}

	kd := strings.Split(valueMap["kills"], " - Deaths: ")
	if len(kd) == 2 {
		detailedPlayer.Kills = util.ToInt(kd[0])
		detailedPlayer.Deaths = util.ToInt(kd[1])
	} else {
		return detailedPlayer, errors.New("k/d data invalid")
	}

	score := strings.Split(valueMap["score"], ", ")
	if len(score) < 4 {
		return detailedPlayer, errors.New("score data invalid")
	}
	if len(strings.Split(score[0], " ")) < 2 || len(strings.Split(score[1], " ")) < 2 || len(strings.Split(score[2], " ")) < 2 || len(strings.Split(score[3], " ")) < 2 {
		return detailedPlayer, errors.New("score data invalid")
	}
	detailedPlayer.Score = hll.Score{
		Combat:  util.ToInt(strings.Split(score[0], " ")[1]),
		Offense: util.ToInt(strings.Split(score[1], " ")[1]),
		Defense: util.ToInt(strings.Split(score[2], " ")[1]),
		Support: util.ToInt(strings.Split(score[3], " ")[1]),
	}

	detailedPlayer.Level = util.ToInt(valueMap["level"])

	return detailedPlayer, nil
}

func (r *Rcon) AddAdmin(id, name string, role hll.AdminRole) error {
	if id == "" {
		return errors.New("invaild id")
	}
	name = strings.ReplaceAll(name, config.LIST_DELIMITER, "") // WARN: this would break the response
	_, err := r.runBasicCommand(fmt.Sprintf("adminadd %s %s \"%s\"", id, role, name))
	return err
}

func (r *Rcon) RemoveAdmin(id string) error {
	return runSetCommand(r, fmt.Sprintf("admindel %s", id))
}

func (r *Rcon) AddVip(id, name string) error {
	if id == "" {
		return errors.New("invaild id")
	}
	name = strings.ReplaceAll(name, config.LIST_DELIMITER, "") // WARN: this would break the response
	_, err := r.runBasicCommand(fmt.Sprintf("vipadd %s \"%s\"", id, name))
	return err
}

func (r *Rcon) RemoveVip(id string) error {
	return runSetCommand(r, fmt.Sprintf("vipdel %s", id))
}
