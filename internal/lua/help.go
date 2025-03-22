package lua

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	lua "github.com/yuin/gopher-lua"
	"github.com/zMoooooritz/go-let-loose/pkg/event"
)

var allPublicEvents = []event.EventType{
	event.EVENT_CONNECTED,
	event.EVENT_DISCONNECTED,
	event.EVENT_KILL,
	event.EVENT_DEATH,
	event.EVENT_TEAMKILL,
	event.EVENT_TEAMDEATH,
	event.EVENT_CHAT,
	event.EVENT_BAN,
	event.EVENT_KICK,
	event.EVENT_MESSAGE,
	event.EVENT_MATCHSTART,
	event.EVENT_MATCHEND,
	event.EVENT_ENTER_ADMINCAM,
	event.EVENT_LEAVE_ADMINCAM,
	event.EVENT_VOTE_KICK_STARTED,
	event.EVENT_VOTE_SUBMITTED,
	event.EVENT_VOTE_KICK_COMPLETED,
	event.EVENT_TEAM_SWITCHED,
	event.EVENT_SQUAD_SWITCHED,
	event.EVENT_SCORE_UPDATE,
	event.EVENT_ROLE_CHANGED,
	event.EVENT_LOADOUT_CHANGED,
	event.EVENT_OBJECTIVE_CAPPED,
}

func RegisterHelp(L *lua.LState) {
	L.SetGlobal("listCommands", L.NewFunction(func(L *lua.LState) int {
		L.Push(GoToLua(L, luaApiSignatures()))
		return 1
	}))
	L.SetGlobal("listEvents", L.NewFunction(func(l *lua.LState) int {
		L.Push(GoToLua(L, allPublicEvents))
		return 1
	}))
}

func luaApiSignatures() []string {
	signatures := []string{
		"registerEvent(string, function)",
		"unregisterEvent(string)",
		"getServerView() -> (error, table)",
		"listCommands() -> (table)",
		"listEvents() -> (table)",
	}
	signatures = append(signatures, rconFunctionSignatures()...)

	sort.Strings(signatures)

	return signatures
}

func rconFunctionSignatures() []string {
	signatures := []string{}
	rcn := GetRconInstance()
	for _, rconFunc := range rconFunctions() {
		method, found := reflect.ValueOf(rcn).Type().MethodByName(rconFunc)
		if !found {
			continue
		}

		params := []string{}
		for i := 1; i < method.Type.NumIn(); i++ {
			params = append(params, luaType(method.Type.In(i)))
		}

		results := []string{}
		otherResults := []string{}
		for i := range method.Type.NumOut() {
			outType := method.Type.Out(i)
			if outType.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
				results = append([]string{"error"}, results...)
			} else {
				otherResults = append(otherResults, luaType(outType))
			}
		}
		results = append(results, otherResults...)

		signatures = append(signatures, fmt.Sprintf("%s(%s) -> (%s)",
			pascalToCamel(method.Name),
			strings.Join(params[:], ", "),
			strings.Join(results[:], ", "),
		))
	}
	return signatures
}
