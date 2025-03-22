package lua

import (
	lua "github.com/yuin/gopher-lua"
)

func RegisterBindings(L *lua.LState) {
	cacheFunctions := map[string]lua.LGFunction{
		"getServerView": func(l *lua.LState) int {
			L.Push(lua.LNil)
			L.Push(GoToLua(L, cacheInstance.GetServerView()))
			return 2
		},
	}

	for name, fn := range cacheFunctions {
		L.SetGlobal(name, L.NewFunction(fn))
	}

	for _, goName := range rconFunctions() {
		luaName := pascalToCamel(goName)
		L.SetGlobal(luaName, L.NewFunction(luaRconRunner(goName)))
	}
}
