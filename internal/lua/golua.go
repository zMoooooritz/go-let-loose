package lua

import (
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

func luaToGo(value lua.LValue, targetType reflect.Type) reflect.Value {
	switch targetType.Kind() {
	case reflect.String:
		return reflect.ValueOf(value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(int(value.(lua.LNumber)))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(uint(value.(lua.LNumber)))
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(float64(value.(lua.LNumber)))
	case reflect.Bool:
		return reflect.ValueOf(value == lua.LTrue)
	case reflect.Struct:
		structValue := reflect.New(targetType).Elem()
		table := value.(*lua.LTable)
		for i := range targetType.NumField() {
			field := targetType.Field(i)
			if field.PkgPath != "" {
				continue // Skip unexported fields
			}
			luaValue := table.RawGetString(field.Name)
			if luaValue != lua.LNil {
				structValue.Field(i).Set(luaToGo(luaValue, field.Type))
			}
		}
		return structValue
	case reflect.Ptr:
		elemType := targetType.Elem()
		ptrValue := reflect.New(elemType)
		ptrValue.Elem().Set(luaToGo(value, elemType))
		return ptrValue
	case reflect.Interface:
		return reflect.ValueOf(value)
	case reflect.Slice, reflect.Array:
		length := value.(*lua.LTable).Len()
		slice := reflect.MakeSlice(targetType, length, length)
		for i := 1; i <= length; i++ {
			slice.Index(i - 1).Set(luaToGo(value.(*lua.LTable).RawGetInt(i), targetType.Elem()))
		}
		return slice
	case reflect.Map:
		mapValue := reflect.MakeMap(targetType)
		value.(*lua.LTable).ForEach(func(key lua.LValue, val lua.LValue) {
			mapValue.SetMapIndex(luaToGo(key, targetType.Key()), luaToGo(val, targetType.Elem()))
		})
		return mapValue
	default:
		return reflect.ValueOf(nil)
	}
}

func GoToLua(L *lua.LState, value any) lua.LValue {
	return goToLua(L, reflect.ValueOf(value))
}

func goToLua(L *lua.LState, value reflect.Value) lua.LValue {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return lua.LNumber(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return lua.LNumber(value.Uint())
	case reflect.Float32, reflect.Float64:
		return lua.LNumber(value.Float())
	case reflect.String:
		return lua.LString(value.String())
	case reflect.Bool:
		if value.Bool() {
			return lua.LTrue
		}
		return lua.LFalse
	case reflect.Struct:
		v := value
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		table := L.NewTable()
		t := v.Type()
		for i := range v.NumField() {
			field := t.Field(i)
			if field.PkgPath != "" {
				continue
			}
			luaValue := goToLua(L, v.Field(i))
			table.RawSetString(field.Name, luaValue)
		}
		return table
	case reflect.Ptr:
		if value.IsNil() {
			return lua.LNil
		}
		return goToLua(L, value.Elem())
	case reflect.Interface:
		if value.IsNil() {
			return lua.LNil
		}
		return goToLua(L, value.Elem())
	case reflect.Slice, reflect.Array:
		list := L.NewTable()
		for i := range value.Len() {
			list.Append(goToLua(L, value.Index(i)))
		}
		return list
	case reflect.Map:
		mapTable := L.NewTable()
		for _, key := range value.MapKeys() {
			mapTable.RawSet(goToLua(L, key), goToLua(L, value.MapIndex(key)))
		}
		return mapTable
	default:
		return lua.LNil
	}
}

func luaType(goType reflect.Type) string {
	if goType.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return "error"
	}

	switch goType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.String:
		return "string"
	case reflect.Struct:
		return "table" // Structs map to Lua tables
	case reflect.Slice, reflect.Array:
		return "table" // Arrays and slices become tables
	case reflect.Map:
		return "table" // Maps are tables in Lua
	case reflect.Ptr:
		return luaType(goType.Elem()) // Dereference pointers
	default:
		return "unknown" // Fallback for unsupported types
	}
}
