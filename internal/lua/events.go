package lua

import (
	lua "github.com/yuin/gopher-lua"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

var observer *luaEventObserver

type luaEventObserver struct {
	callbacks map[hll.EventType]*lua.LFunction
	L         *lua.LState
}

func (o *luaEventObserver) Notify(event hll.Event) {
	if callback, exists := o.callbacks[event.Type()]; exists {
		o.L.Push(callback)
		o.L.Push(GoToLua(o.L, event))
		if err := o.L.PCall(1, 0, nil); err != nil {
			println("Error in Lua event callback:", err.Error())
		}
	}
}

func RegisterEvents(L *lua.LState) {
	rcon := GetRconInstance()
	if rcon == nil {
		return
	}

	callbacks := make(map[hll.EventType]*lua.LFunction)

	observer = &luaEventObserver{callbacks: callbacks, L: L}
	rcon.Events.Register(observer)

	registerHandler := func(eventType hll.EventType, handlerName string) {
		L.SetGlobal(handlerName, L.NewFunction(func(L *lua.LState) int {
			callback := L.Get(1)
			if callback.Type() != lua.LTFunction {
				L.Push(lua.LString("Error: " + handlerName + " expects a function as argument"))
				return 1
			}

			callbacks[eventType] = callback.(*lua.LFunction)
			L.Push(lua.LNil) // No error
			return 1
		}))
	}

	registerHandler(hll.EVENT_KILL, "onKill")
	registerHandler(hll.EVENT_DEATH, "onDeath")
	registerHandler(hll.EVENT_TEAMKILL, "onTeamKill")
	registerHandler(hll.EVENT_TEAMDEATH, "onTeamDeath")
	registerHandler(hll.EVENT_CONNECTED, "onConnected")
	registerHandler(hll.EVENT_DISCONNECTED, "onDisconnected")
	registerHandler(hll.EVENT_CHAT, "onChat")
	registerHandler(hll.EVENT_BAN, "onBan")
	registerHandler(hll.EVENT_KICK, "onKick")
	registerHandler(hll.EVENT_MESSAGE, "onMessage")
	registerHandler(hll.EVENT_MATCHSTART, "onMatchStart")
	registerHandler(hll.EVENT_MATCHEND, "onMatchEnd")
	registerHandler(hll.EVENT_ENTER_ADMINCAM, "onEnterAdminCam")
	registerHandler(hll.EVENT_LEAVE_ADMINCAM, "onLeaveAdminCam")
	registerHandler(hll.EVENT_VOTE_KICK_STARTED, "onVoteKickStarted")
	registerHandler(hll.EVENT_VOTE_SUBMITTED, "onVoteSubmitted")
	registerHandler(hll.EVENT_VOTE_KICK_COMPLETED, "onVoteKickCompleted")
	registerHandler(hll.EVENT_TEAM_SWITCHED, "onTeamSwitched")
	registerHandler(hll.EVENT_SQUAD_SWITCHED, "onSquadSwitched")
	registerHandler(hll.EVENT_SCORE_UPDATE, "onScoreUpdate")
	registerHandler(hll.EVENT_ROLE_CHANGED, "onRoleChanged")
	registerHandler(hll.EVENT_LOADOUT_CHANGED, "onLoadoutChanged")
	registerHandler(hll.EVENT_OBJECTIVE_CAPPED, "onObjectiveCapped")
	registerHandler(hll.EVENT_POSITION_CHANGED, "onPositionChanged")
	registerHandler(hll.EVENT_CLAN_TAG_CHANGED, "onClanTagChanged")
}

func UnregisterEvents() {
	rcon := GetRconInstance()
	if rcon == nil || observer == nil {
		return
	}
	rcon.Events.Unregister(observer)
	observer = nil
}
