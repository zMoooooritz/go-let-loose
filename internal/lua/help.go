package lua

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	lua "github.com/yuin/gopher-lua"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

var allPublicEvents = []hll.EventType{
	hll.EVENT_CONNECTED,
	hll.EVENT_DISCONNECTED,
	hll.EVENT_KILL,
	hll.EVENT_DEATH,
	hll.EVENT_TEAMKILL,
	hll.EVENT_TEAMDEATH,
	hll.EVENT_CHAT,
	hll.EVENT_BAN,
	hll.EVENT_KICK,
	hll.EVENT_MESSAGE,
	hll.EVENT_MATCHSTART,
	hll.EVENT_MATCHEND,
	hll.EVENT_ENTER_ADMINCAM,
	hll.EVENT_LEAVE_ADMINCAM,
	hll.EVENT_VOTE_KICK_STARTED,
	hll.EVENT_VOTE_SUBMITTED,
	hll.EVENT_VOTE_KICK_COMPLETED,
	hll.EVENT_TEAM_SWITCHED,
	hll.EVENT_SQUAD_SWITCHED,
	hll.EVENT_SCORE_UPDATE,
	hll.EVENT_ROLE_CHANGED,
	hll.EVENT_LOADOUT_CHANGED,
	hll.EVENT_OBJECTIVE_CAPPED,
	hll.EVENT_POSITION_CHANGED,
	hll.EVENT_CLAN_TAG_CHANGED,
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
