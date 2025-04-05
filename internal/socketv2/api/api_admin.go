package api

import (
	"strconv"
	"strings"
	"time"
)

type AddAdmin struct {
	PlayerId   string `json:"playerId"`
	AdminGroup string `json:"adminGroup"`
	Comment    string `json:"comment"`
}

type RemoveAdmin struct {
	PlayerId string `json:"playerId"`
}

type AdminLog struct {
	LogBackTrackTime int32  `json:"logBackTrackTime"`
	Filters          string `json:"filters"`
}

type ResponseAdminLog struct {
	Entries []AdminLogEntry `json:"entries"`
}

type AdminLogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

func (a AdminLogEntry) Time() time.Time {
	li := strings.LastIndex(a.Timestamp, ":")
	p, _ := strconv.Atoi(a.Timestamp[li+1:])
	r, _ := time.Parse("2006.01.02-15:04:05", a.Timestamp[:li])
	return time.Date(r.Year(), r.Month(), r.Day(), r.Hour(), r.Minute(), r.Second(), p*1000000, r.Location())
}
