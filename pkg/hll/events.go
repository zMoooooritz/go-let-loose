package hll

import (
	"time"
)

type EventType string

const (
	EVENT_CONNECTED    EventType = "CONNECTED"
	EVENT_DISCONNECTED EventType = "DISCONNECTED"
	EVENT_KILL         EventType = "KILL"
	EVENT_DEATH        EventType = "DEATH"
	EVENT_TEAMKILL     EventType = "TEAM KILL"
	EVENT_TEAMDEATH    EventType = "TEAM DEATH"
	EVENT_CHAT         EventType = "CHAT"
	EVENT_BAN          EventType = "BAN"
	EVENT_KICK         EventType = "KICK"
	EVENT_MESSAGE      EventType = "MESSAGE"
	EVENT_MATCHSTART   EventType = "MATCH START"
	EVENT_MATCHEND     EventType = "MATCH END"
	// event_teamswitch   EventType = "TEAMSWITCH" // WARN: generic hll event; this event does not have a player id
	// event_admincam     EventType = "Player"
	// event_vote         EventType = "VOTESYS"

	// native log events above; custom events below

	EVENT_ENTER_ADMINCAM      EventType = "ADMINCAM ENTERED"
	EVENT_LEAVE_ADMINCAM      EventType = "ADMINCAM LEFT"
	EVENT_VOTE_KICK_STARTED   EventType = "VOTE KICK STARTED"
	EVENT_VOTE_SUBMITTED      EventType = "VOTE SUBMITTED"
	EVENT_VOTE_KICK_COMPLETED EventType = "VOTE KICK COMPLETED"
	EVENT_TEAM_SWITCHED       EventType = "TEAM SWITCHED"
	EVENT_SQUAD_SWITCHED      EventType = "SQUAD SWITCHED"
	EVENT_SCORE_UPDATE        EventType = "SCORE UPDATE"
	EVENT_ROLE_CHANGED        EventType = "ROLE CHANGED"
	EVENT_LOADOUT_CHANGED     EventType = "LOADOUT CHANGED"
	EVENT_OBJECTIVE_CAPPED    EventType = "OBJECTIVE CAPPED"
	EVENT_POSITION_CHANGED    EventType = "POSITION CHANGED"
	EVENT_CLAN_TAG_CHANGED    EventType = "CLAN TAG CHANGED"
	EVENT_GENERIC             EventType = "GENERIC"
)

type Event interface {
	Type() EventType
	Time() time.Time
	AffectedPlayers() []PlayerInfo
}

type GenericEvent struct {
	EventType EventType
	EventTime time.Time
}

func (ge GenericEvent) Type() EventType {
	return ge.EventType
}

func (ge GenericEvent) Time() time.Time {
	return ge.EventTime
}

func (ge GenericEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{}
}

type ConnectEvent struct {
	GenericEvent
	Player PlayerInfo
}

func (ce ConnectEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{ce.Player}
}

type DisconnectEvent struct {
	GenericEvent
	Player PlayerInfo
}

func (de DisconnectEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{de.Player}
}

type KillEvent struct {
	GenericEvent
	Killer PlayerInfo
	Victim PlayerInfo
	Weapon Weapon
}

func (ke KillEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{ke.Killer, ke.Victim}
}

type DeathEvent struct {
	GenericEvent
	Victim PlayerInfo
	Killer PlayerInfo
	Weapon Weapon
}

func (de DeathEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{de.Victim, de.Killer}
}

type TeamKillEvent struct {
	GenericEvent
	Killer PlayerInfo
	Victim PlayerInfo
	Weapon Weapon
}

func (tke TeamKillEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{tke.Killer, tke.Victim}
}

type TeamDeathEvent struct {
	GenericEvent
	Victim PlayerInfo
	Killer PlayerInfo
	Weapon Weapon
}

func (tde TeamDeathEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{tde.Victim, tde.Killer}
}

type TeamSwitchEvent struct {
	GenericEvent
	Player PlayerInfo
	From   Team
	To     Team
}

func (tse TeamSwitchEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{tse.Player}
}

type ChatEvent struct {
	GenericEvent
	Player  PlayerInfo
	Team    Team
	Scope   ChatScope
	Message string
}

func (ce ChatEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{ce.Player}
}

type BanEvent struct {
	GenericEvent
	Player PlayerInfo
	Reason string
}

func (be BanEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{be.Player}
}

type KickEvent struct {
	GenericEvent
	Player PlayerInfo
	Reason string
}

func (ke KickEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{ke.Player}
}

type MessageEvent struct {
	GenericEvent
	Player  PlayerInfo
	Message string
}

func (me MessageEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{me.Player}
}

type MatchStartEvent struct {
	GenericEvent
	Map GameMap
}

func (mse MatchStartEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{}
}

type MatchEndEvent struct {
	GenericEvent
	Map   GameMap
	Score TeamData
}

func (mee MatchEndEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{}
}

type AdminCamEnteredEvent struct {
	GenericEvent
	Player PlayerInfo
}

func (acee AdminCamEnteredEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{acee.Player}
}

type AdminCamLeftEvent struct {
	GenericEvent
	Player PlayerInfo
}

func (acle AdminCamLeftEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{acle.Player}
}

type VoteStartedEvent struct {
	GenericEvent
	Reason    string
	ID        int
	Initiator PlayerInfo
	Target    PlayerInfo
}

func (vse VoteStartedEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{vse.Initiator, vse.Target}
}

type VoteSubmittedEvent struct {
	GenericEvent
	Submitter PlayerInfo
	ID        int
	Vote      string
}

func (vse VoteSubmittedEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{vse.Submitter}
}

type VoteCompletedEvent struct {
	GenericEvent
	Reason    string
	Result    string
	ID        int
	Initiator PlayerInfo
	Target    PlayerInfo
}

func (vce VoteCompletedEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{vce.Initiator, vce.Target}
}

type ObjectiveCaptureEvent struct {
	GenericEvent
	OldScore TeamData
	NewScore TeamData
}

func (oce ObjectiveCaptureEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{}
}

type PlayerScoreUpdateEvent struct {
	GenericEvent
	Player   PlayerInfo
	OldScore Score
	NewScore Score
}

func (psue PlayerScoreUpdateEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{psue.Player}
}

type PlayerSwitchTeamEvent struct {
	GenericEvent
	Player  PlayerInfo
	OldTeam Team
	NewTeam Team
}

func (pstee PlayerSwitchTeamEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{pstee.Player}
}

type PlayerSwitchSquadEvent struct {
	GenericEvent
	Player   PlayerInfo
	OldSquad Unit
	NewSquad Unit
}

func (psse PlayerSwitchSquadEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{psse.Player}
}

type PlayerChangeRoleEvent struct {
	GenericEvent
	Player  PlayerInfo
	OldRole Role
	NewRole Role
}

func (pcre PlayerChangeRoleEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{pcre.Player}
}

type PlayerChangeLoadoutEvent struct {
	GenericEvent
	Player     PlayerInfo
	OldLoadout string
	NewLoadout string
}

func (pcle PlayerChangeLoadoutEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{pcle.Player}
}

type PlayerPositionChangedEvent struct {
	GenericEvent
	Player PlayerInfo
	OldPos Position
	NewPos Position
}

func (ppce PlayerPositionChangedEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{ppce.Player}
}

type PlayerClanTagChangedEvent struct {
	GenericEvent
	Player     PlayerInfo
	OldClanTag string
	NewClanTag string
}

func (pctce PlayerClanTagChangedEvent) AffectedPlayers() []PlayerInfo {
	return []PlayerInfo{pctce.Player}
}
