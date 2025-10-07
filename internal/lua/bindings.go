package lua

import (
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func RegisterBindings(L *lua.LState) {
	for _, goName := range rconFunctions() {
		// Skip event handler functions as they're handled separately
		if strings.HasPrefix(goName, callbackPrefix) {
			continue
		}

		luaName := pascalToCamel(goName)
		L.SetGlobal(luaName, L.NewFunction(luaRconRunner(goName)))
	}
}
