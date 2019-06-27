/* This is free and unencumbered software released into the public domain. */

package rt

import (
	"os"

	"github.com/Azure/golua/lua"
	"github.com/Azure/golua/std"
	"github.com/conreality/conreality-gdk/gdk"
)

// Thread
type Thread struct {
	state *lua.State
}

// NewThread
func NewThread() (*Thread, error) {
	var opts = []lua.Option{
		lua.WithTrace(false),
		lua.WithVerbose(false),
		lua.WithChecks(true),
	}
	thread := &Thread{state: lua.NewState(opts...)}

	std.Open(thread.state)
	thread.state.Push(true)
	thread.state.SetGlobal("_U")
	//thread.state.Debug(false)

	var libs = []struct {
		Name string
		Open lua.Func
	}{
		{"gdk", lua.Func(gdkLibraryOpen)},
	}
	for _, lib := range libs {
		thread.state.Logf("opening stdlib mode %q", lib.Name)
		thread.state.Require(lib.Name, lib.Open, true)
		thread.state.Pop()
	}

	self := gdk.Agent{}
	thread.state.Push(&self)                    // []          => [self] <-TOS
	thread.state.PushIndex(-1)                  // [self]      => [self self]
	thread.state.SetGlobal("self")              // [self self] => [self]
	thread.state.NewMetaTable("Agent")          // [self] => [self mt]
	thread.state.PushIndex(-1)                  // [self mt] => [self mt mt]
	thread.state.SetField(-2, "__index")        // [self mt mt] => [self mt] // metatable.__index = metatable
	thread.state.SetFuncs(gdkAgentMethods(), 0) // ...
	thread.state.SetMetaTableAt(-2)             // [self mt] => [self]
	thread.state.Pop()                          // [self] => []

	err := thread.state.ExecText(gdk.Prelude)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

// Destroy
func (thread *Thread) Destroy() {
	thread.state.Close()
}

// DumpStack
func (thread *Thread) DumpStack() {
	thread.state.DumpStack(os.Stderr)
}

// EvalFile
func (thread *Thread) EvalFile(filePath string) error {
	return thread.state.ExecFile(filePath)
}

// EvalScript
func (thread *Thread) EvalScript(script string) error {
	return thread.state.ExecText(script)
}
