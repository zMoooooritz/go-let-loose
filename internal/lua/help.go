package lua

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func RegisterHelp(L *lua.LState) {
	L.SetGlobal("listCommands", L.NewFunction(func(L *lua.LState) int {
		L.Push(GoToLua(L, luaApiSignatures()))
		return 1
	}))
}

func luaApiSignatures() []string {
	signatures := []string{
		"listCommands() -> (table)",
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

		if strings.HasPrefix(method.Name, callbackPrefix) {
			signatures = append(signatures, fmt.Sprintf("%s(function) -> (error)", pascalToCamel(method.Name)))
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
