package rcon

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

func (r *Rcon) GetTempBans() ([]hll.ServerBan, error) {
	resp, err := r.runListCommand("get tempbans")
	if err != nil {
		return []hll.ServerBan{}, err
	}
	return ParseBans(resp), nil
}

func (r *Rcon) GetPermaBans() ([]hll.ServerBan, error) {
	resp, err := r.runListCommand("get permabans")
	if err != nil {
		return []hll.ServerBan{}, err
	}
	return ParseBans(resp), nil
}

func ParseBans(banLogs []string) []hll.ServerBan {
	bans := []hll.ServerBan{}
	for _, line := range banLogs {
		if line == "" {
			continue
		}

		regex := regexp.MustCompile(`^(?P<ID>[\w-]+) :(?: nickname "(?P<Nickname>.+?)")? banned(?: for (?P<Duration>\d+ hours))? on (?P<Date>\d{4}\.\d{2}\.\d{2}-\d{2}\.\d{2}\.\d{2}) for "(?P<Reason>(?s:.+?))" by admin "(?P<Admin>.+?)"$`)

		matches := regex.FindStringSubmatch(line)
		if matches == nil {
			logger.Error("no match found for line:", line)
			continue
		}

		// Extract named groups into a map
		result := make(map[string]string)
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = matches[i]
			}
		}

		// Parse the timestamp
		timestamp, err := time.Parse("2006.01.02-15.04.05", result["Date"])
		if err != nil {
			logger.Error("error parsing date:", err)
			continue
		}

		// Parse the duration if present
		var duration time.Duration
		var banType hll.BanType

		if result["Duration"] != "" {
			hours := strings.Split(result["Duration"], " ")[0]
			durationHours, err := time.ParseDuration(fmt.Sprintf("%sh", hours))
			if err != nil {
				logger.Error("error parsing duration:", err)
				continue
			}
			duration = durationHours
			banType = hll.TempBan
		} else {
			banType = hll.PermaBan
		}

		serverBan := hll.ServerBan{
			Type: banType,
			Player: hll.PlayerInfo{
				ID:   result["ID"],
				Name: result["Nickname"],
			},
			Timestamp: timestamp,
			Duration:  duration,
			Reason:    result["Reason"],
			AdminName: result["Admin"],
			RawLog:    line,
		}

		bans = append(bans, serverBan)
	}
	return bans
}

func (r *Rcon) SendMessage(playerID string, message string) error {
	return runSetCommand(r, fmt.Sprintf("message %s %s", playerID, message))
}

func (r *Rcon) PunishPlayer(player, reason string) error {
	reason = strings.ReplaceAll(reason, config.LIST_DELIMITER, "") // WARN: this would break the response
	return runSetCommand(r, fmt.Sprintf("punish %s %s", player, reason))
}

func (r *Rcon) SwitchPlayerOnDeath(player string) error {
	return runSetCommand(r, fmt.Sprintf("switchteamondeath %s", player))
}

func (r *Rcon) SwitchPlayerNow(player string) error {
	return runSetCommand(r, fmt.Sprintf("switchteamnow %s", player))
}

func (r *Rcon) KickPlayer(player, reason string) error {
	reason = strings.ReplaceAll(reason, config.LIST_DELIMITER, "") // WARN: this would break the response
	return runSetCommand(r, fmt.Sprintf("kick %s %s", player, reason))
}

func (r *Rcon) TempBanPlayer(player string, duration int, reason, admin string) error {
	duration = min(1, duration)
	reason = strings.ReplaceAll(reason, config.LIST_DELIMITER, "") // WARN: this would break the response
	return runSetCommand(r, fmt.Sprintf("tempban %s %d %s %s", player, duration, reason, admin))
}

func (r *Rcon) PardonTempBanPlayer(ban hll.ServerBan) error {
	return runSetCommand(r, fmt.Sprintf("pardontempban %s", ban.RawLog))
}

func (r *Rcon) PermaBanPlayer(player, reason, admin string) error {
	reason = strings.ReplaceAll(reason, config.LIST_DELIMITER, "") // WARN: this would break the response
	return runSetCommand(r, fmt.Sprintf("tempban %s %s %s", player, reason, admin))
}

func (r *Rcon) PardonPermaBanPlayer(ban hll.ServerBan) error {
	return runSetCommand(r, fmt.Sprintf("pardonpermaban %s", ban.RawLog))
}
