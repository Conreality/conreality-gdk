/* This is free and unencumbered software released into the public domain. */

package rt

import (
	"github.com/Azure/golua/lua"
	"github.com/conreality/conreality-gdk/gdk"
)

func gdkLibraryOpen(state *lua.State) int {
	funcs := gdkLibraryFunctions()
	state.NewTableSize(0, len(funcs))
	state.SetFuncs(funcs, 0)
	return 1
}

func gdkLibraryFunctions() map[string]lua.Func {
	return map[string]lua.Func{
		"version": func(state *lua.State) int {
			state.Push(gdk.Version)
			return 1
		},
	}
}

func gdkAgentMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	for id, fn := range gdk.AgentPredicates() {
		result[id] = func(state *lua.State) int {
			obj := state.CheckUserData(1, "Agent").(gdk.Agent)
			state.Push(fn(obj))
			return 1 // [] => [result]
		}
	}
	return result
}

func gdkGameMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	for id, fn := range gdk.GamePredicates() {
		result[id] = func(state *lua.State) int {
			obj := state.CheckUserData(1, "Game").(gdk.Game)
			state.Push(fn(obj))
			return 1 // [] => [result]
		}
	}
	return result
}

func gdkTheaterMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	// TODO
	return result
}

func gdkUnitMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	for id, fn := range gdk.UnitPredicates() {
		result[id] = func(state *lua.State) int {
			obj := state.CheckUserData(1, "Unit").(gdk.Unit)
			state.Push(fn(obj))
			return 1 // [] => [result]
		}
	}
	return result
}

func gdkRegisterGlobal(state *lua.State, globalVar string, globalVal interface{}, metatableID string, metatableFuncs map[string]lua.Func) {
	state.Push(globalVal)             // []          => [val] <-TOS
	state.PushIndex(-1)               // [val]       => [val val]
	state.SetGlobal(globalVar)        // [val val]   => [val]
	state.NewMetaTable(metatableID)   // [val]       => [val mt]
	state.PushIndex(-1)               // [val mt]    => [val mt mt]
	state.SetField(-2, "__index")     // [val mt mt] => [val mt] // metatable.__index = metatable
	state.SetFuncs(metatableFuncs, 0) // ...
	state.SetMetaTableAt(-2)          // [val mt]    => [val]
	state.Pop()                       // [val]       => []
}
