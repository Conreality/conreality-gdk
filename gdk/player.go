/* This is free and unencumbered software released into the public domain. */

package gdk

// Player
type Player struct {
	Agent Agent
}

// PlayerPredicate
type PlayerPredicate func(*Player) bool
