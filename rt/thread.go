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
	state   *lua.State
	Self    *gdk.Agent
	Unit    *gdk.Unit
	Theater *gdk.Theater
	Game    *gdk.Game
}

// NewThread
func NewThread() (*Thread, error) {
	var opts = []lua.Option{
		lua.WithTrace(false),
		lua.WithVerbose(false),
		lua.WithChecks(true),
	}

	thread := &Thread{state: lua.NewState(opts...)}
	thread.Self = &gdk.Agent{}
	thread.Unit = &gdk.Unit{}
	thread.Theater = &gdk.Theater{}
	thread.Game = &gdk.Game{}

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

	gdkRegisterGlobal(thread.state, "self", thread.Self, "Agent", gdkAgentMethods())
	gdkRegisterGlobal(thread.state, "unit", thread.Unit, "Unit", gdkUnitMethods())
	gdkRegisterGlobal(thread.state, "theater", thread.Theater, "Theater", gdkTheaterMethods())
	gdkRegisterGlobal(thread.state, "game", thread.Game, "Game", gdkGameMethods())
	// TODO: here, now, targets

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
