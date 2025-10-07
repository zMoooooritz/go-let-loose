package lua

import (
	lua "github.com/yuin/gopher-lua"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

var eventHandlers = make(map[string]*lua.LFunction)
var forwarder Forwarder

type Forwarder struct {
	L *lua.LState
}

func (f *Forwarder) Notify(e hll.Event) {
	handler, ok := eventHandlers[string(e.Type())]
	if !ok {
		return
	}

	_ = f.L.CallByParam(lua.P{
		Fn:      handler,
		NRet:    0,
		Protect: true,
	}, GoToLua(f.L, e))
}

func RegisterEvents(L *lua.LState) {
	L.SetGlobal("registerEvent", L.NewFunction(registerEvent))
	L.SetGlobal("unregisterEvent", L.NewFunction(unregisterEvent))

	forwarder = Forwarder{L}
	rcn := GetRconInstance()
	rcn.Events.Register(&forwarder)
}

func UnregisterEvents() {
	rcn := GetRconInstance()
	rcn.Events.Unregister(&forwarder)
}

func registerEvent(L *lua.LState) int {
	eventName := L.CheckString(1)
	handler := L.CheckFunction(2)
	eventHandlers[eventName] = handler
	return 0
}

func unregisterEvent(L *lua.LState) int {
	eventName := L.CheckString(1)
	delete(eventHandlers, eventName)
	return 0
}
