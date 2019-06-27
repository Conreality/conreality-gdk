/* This is free and unencumbered software released into the public domain. */

package gdk

// Game
type Game interface {
	IsPaused() bool
	IsPrivate() bool
	IsPublic() bool
	IsStarted() bool
	IsStopped() bool
}

// GamePredicate
type GamePredicate func(Game) bool

// GamePredicates
func GamePredicates() map[string]GamePredicate {
	return map[string]GamePredicate{
		"is_paused": func(game Game) bool {
			return game.IsPaused()
		},
		"is_private": func(game Game) bool {
			return game.IsPrivate()
		},
		"is_public": func(game Game) bool {
			return game.IsPublic()
		},
		"is_started": func(game Game) bool {
			return game.IsStarted()
		},
		"is_stopped": func(game Game) bool {
			return game.IsStopped()
		},
	}
}
