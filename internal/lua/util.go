package lua

import (
	"fmt"
	"reflect"
	"slices"
	"unicode"

	lua "github.com/yuin/gopher-lua"
)

func rconFunctions() []string {
	excludedFuncs := []string{
		"Close",
		"QueueJob",
		"RunCommand",
	}

	funcs := []string{}
	rcn := GetRconInstance()
	rcnMethods := reflect.ValueOf(rcn).NumMethod()
	for i := range rcnMethods {
		method := reflect.ValueOf(rcn).Type().Method(i)

		if slices.Contains(excludedFuncs, method.Name) {
			continue
		}

		funcs = append(funcs, method.Name)
	}

	return funcs
}

func pascalToCamel(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func luaRconRunner(methodName string) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		rcon := GetRconInstance()
		if rcon == nil {
			L.Push(lua.LString("RCON not initialized"))
			return 1
		}

		// Use reflection to call the appropriate method
		method := reflect.ValueOf(rcon).MethodByName(methodName)
		if !method.IsValid() {
			L.Push(lua.LString("Invalid method: " + methodName))
			return 1
		}

		argCount := L.GetTop()
		if method.Type().NumIn() != argCount {
			L.Push(lua.LString("Unexpected method signature"))
			return 1
		}

		// Prepare arguments
		args := make([]reflect.Value, argCount)
		for i := range argCount {
			arg := L.Get(i + 1)
			args[i] = luaToGo(arg, method.Type().In(i))
		}

		// Call the method
		results := method.Call(args)
		retCount := len(results)
		if retCount == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// Extract error if present
		var err error
		if results[retCount-1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			err, _ = results[retCount-1].Interface().(error)
			retCount--
		}

		// Push error first if present
		if err != nil {
			L.Push(lua.LString(fmt.Sprintf("Error: %s", err)))
		} else {
			L.Push(lua.LNil)
		}

		// Push other results to Lua stack
		for i := range retCount {
			L.Push(goToLua(L, results[i]))
		}

		return retCount + 1
	}
}
