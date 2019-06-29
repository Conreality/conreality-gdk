/* This is free and unencumbered software released into the public domain. */

package rt

import (
	"bytes"
	"os"

	"github.com/Azure/golua/lua"
	"github.com/Azure/golua/std"
	"github.com/conreality/conreality-gdk/gdk"
)

// Thread
type Thread struct {
	Model *Model
	state *lua.State
}

// NewThread
func NewThread(model *Model, usePrelude bool) (*Thread, error) {
	var opts = []lua.Option{
		lua.WithTrace(false),
		lua.WithVerbose(false),
		lua.WithChecks(true),
	}

	thread := &Thread{Model: model, state: lua.NewState(opts...)}

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

	gdkRegisterGlobal(thread.state, "self", thread.Model.Self, "Agent", gdkAgentMethods())
	gdkRegisterGlobal(thread.state, "unit", thread.Model.Unit, "Unit", gdkUnitMethods())
	gdkRegisterGlobal(thread.state, "theater", thread.Model.Theater, "Theater", gdkTheaterMethods())
	gdkRegisterGlobal(thread.state, "game", thread.Model.Game, "Game", gdkGameMethods())
	// TODO: here, now, targets

	if usePrelude {
		err := thread.state.ExecText(gdk.Prelude)
		if err != nil {
			return nil, err
		}
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

// EvalChunk
func (thread *Thread) EvalChunk(chunk []byte) error {
	return thread.state.ExecFrom(bytes.NewReader(chunk))
}

// EvalScript
func (thread *Thread) EvalScript(script string) error {
	return thread.state.ExecText(script)
}

// EvalFile
func (thread *Thread) EvalFile(filePath string) error {
	return thread.state.ExecFile(filePath)
}
