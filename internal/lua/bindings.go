package lua

import (
	lua "github.com/yuin/gopher-lua"
)

func RegisterBindings(L *lua.LState) {
	for _, goName := range rconFunctions() {
		luaName := pascalToCamel(goName)
		L.SetGlobal(luaName, L.NewFunction(luaRconRunner(goName)))
	}
}
