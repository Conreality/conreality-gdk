/* This is free and unencumbered software released into the public domain. */

package rt

import (
	"github.com/Azure/golua/lua"
	"github.com/conreality/conreality-gdk/gdk"
)

func gdkLibraryOpen(vm *lua.State) int {
	funcs := gdkLibraryFunctions()
	vm.NewTableSize(0, len(funcs))
	vm.SetFuncs(funcs, 0)
	return 1
}

func gdkLibraryFunctions() map[string]lua.Func {
	return map[string]lua.Func{
		"version": func(vm *lua.State) int {
			vm.Push(gdk.Version)
			return 1
		},
	}
}

func gdkAgentMethods() map[string]lua.Func {
	result := make(map[string]lua.Func)
	for k, v := range gdk.AgentPredicates() {
		result[k] = gdkAgentPredicate(v)
	}
	return result
}

func gdkAgentPredicate(f gdk.AgentPredicate) lua.Func {
	return func(vm *lua.State) int {
		self := vm.CheckUserData(1, "Agent").(*gdk.Agent)
		vm.Push(f(self))
		return 1 // [] => [result]
	}
}
