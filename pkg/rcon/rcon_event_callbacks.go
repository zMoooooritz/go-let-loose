package rcon

import "github.com/zMoooooritz/go-let-loose/pkg/hll"

type callbackObserver[T hll.Event] struct {
	callback func(T)
}

func (cb callbackObserver[T]) Notify(e hll.Event) {
	if typedEvent, ok := e.(T); ok {
		cb.callback(typedEvent)
	}
}

func (r *Rcon) OnConnected(callback func(hll.ConnectEvent)) {
	r.Events.registerEvent(hll.EVENT_CONNECTED, callbackObserver[hll.ConnectEvent]{callback: callback})
}

func (r *Rcon) OnDisconnected(callback func(hll.DisconnectEvent)) {
	r.Events.registerEvent(hll.EVENT_DISCONNECTED, callbackObserver[hll.DisconnectEvent]{callback: callback})
}

func (r *Rcon) OnKill(callback func(hll.KillEvent)) {
	r.Events.registerEvent(hll.EVENT_KILL, callbackObserver[hll.KillEvent]{callback: callback})
}

func (r *Rcon) OnDeath(callback func(hll.DeathEvent)) {
	r.Events.registerEvent(hll.EVENT_DEATH, callbackObserver[hll.DeathEvent]{callback: callback})
}

func (r *Rcon) OnTeamKill(callback func(hll.TeamKillEvent)) {
	r.Events.registerEvent(hll.EVENT_TEAMKILL, callbackObserver[hll.TeamKillEvent]{callback: callback})
}

func (r *Rcon) OnTeamDeath(callback func(hll.TeamDeathEvent)) {
	r.Events.registerEvent(hll.EVENT_TEAMDEATH, callbackObserver[hll.TeamDeathEvent]{callback: callback})
}

func (r *Rcon) OnChat(callback func(hll.ChatEvent)) {
	r.Events.registerEvent(hll.EVENT_CHAT, callbackObserver[hll.ChatEvent]{callback: callback})
}

func (r *Rcon) OnBan(callback func(hll.BanEvent)) {
	r.Events.registerEvent(hll.EVENT_BAN, callbackObserver[hll.BanEvent]{callback: callback})
}

func (r *Rcon) OnKick(callback func(hll.KickEvent)) {
	r.Events.registerEvent(hll.EVENT_KICK, callbackObserver[hll.KickEvent]{callback: callback})
}

func (r *Rcon) OnMessage(callback func(hll.MessageEvent)) {
	r.Events.registerEvent(hll.EVENT_MESSAGE, callbackObserver[hll.MessageEvent]{callback: callback})
}

func (r *Rcon) OnMatchStart(callback func(hll.MatchStartEvent)) {
	r.Events.registerEvent(hll.EVENT_MATCHSTART, callbackObserver[hll.MatchStartEvent]{callback: callback})
}

func (r *Rcon) OnMatchEnd(callback func(hll.MatchEndEvent)) {
	r.Events.registerEvent(hll.EVENT_MATCHEND, callbackObserver[hll.MatchEndEvent]{callback: callback})
}

func (r *Rcon) OnEnterAdminCam(callback func(hll.AdminCamEnteredEvent)) {
	r.Events.registerEvent(hll.EVENT_ENTER_ADMINCAM, callbackObserver[hll.AdminCamEnteredEvent]{callback: callback})
}

func (r *Rcon) OnLeaveAdminCam(callback func(hll.AdminCamLeftEvent)) {
	r.Events.registerEvent(hll.EVENT_LEAVE_ADMINCAM, callbackObserver[hll.AdminCamLeftEvent]{callback: callback})
}

func (r *Rcon) OnVoteKickStarted(callback func(hll.VoteStartedEvent)) {
	r.Events.registerEvent(hll.EVENT_VOTE_KICK_STARTED, callbackObserver[hll.VoteStartedEvent]{callback: callback})
}

func (r *Rcon) OnVoteSubmitted(callback func(hll.VoteSubmittedEvent)) {
	r.Events.registerEvent(hll.EVENT_VOTE_SUBMITTED, callbackObserver[hll.VoteSubmittedEvent]{callback: callback})
}

func (r *Rcon) OnVoteKickCompleted(callback func(hll.VoteCompletedEvent)) {
	r.Events.registerEvent(hll.EVENT_VOTE_KICK_COMPLETED, callbackObserver[hll.VoteCompletedEvent]{callback: callback})
}

func (r *Rcon) OnTeamSwitched(callback func(hll.PlayerSwitchTeamEvent)) {
	r.Events.registerEvent(hll.EVENT_TEAM_SWITCHED, callbackObserver[hll.PlayerSwitchTeamEvent]{callback: callback})
}

func (r *Rcon) OnSquadSwitched(callback func(hll.PlayerSwitchSquadEvent)) {
	r.Events.registerEvent(hll.EVENT_SQUAD_SWITCHED, callbackObserver[hll.PlayerSwitchSquadEvent]{callback: callback})
}

func (r *Rcon) OnScoreUpdate(callback func(hll.PlayerScoreUpdateEvent)) {
	r.Events.registerEvent(hll.EVENT_SCORE_UPDATE, callbackObserver[hll.PlayerScoreUpdateEvent]{callback: callback})
}

func (r *Rcon) OnRoleChanged(callback func(hll.PlayerChangeRoleEvent)) {
	r.Events.registerEvent(hll.EVENT_ROLE_CHANGED, callbackObserver[hll.PlayerChangeRoleEvent]{callback: callback})
}

func (r *Rcon) OnLoadoutChanged(callback func(hll.PlayerChangeLoadoutEvent)) {
	r.Events.registerEvent(hll.EVENT_LOADOUT_CHANGED, callbackObserver[hll.PlayerChangeLoadoutEvent]{callback: callback})
}

func (r *Rcon) OnObjectiveCapped(callback func(hll.ObjectiveCaptureEvent)) {
	r.Events.registerEvent(hll.EVENT_OBJECTIVE_CAPPED, callbackObserver[hll.ObjectiveCaptureEvent]{callback: callback})
}

func (r *Rcon) OnPositionChanged(callback func(hll.PlayerPositionChangedEvent)) {
	r.Events.registerEvent(hll.EVENT_POSITION_CHANGED, callbackObserver[hll.PlayerPositionChangedEvent]{callback: callback})
}

func (r *Rcon) OnClanTagChanged(callback func(hll.PlayerClanTagChangedEvent)) {
	r.Events.registerEvent(hll.EVENT_CLAN_TAG_CHANGED, callbackObserver[hll.PlayerClanTagChangedEvent]{callback: callback})
}
