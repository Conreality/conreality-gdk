/* This is free and unencumbered software released into the public domain. */

package gdk

// Game
type Game struct{}

// GamePredicate
type GamePredicate func(*Game) bool

// GamePredicates
func GamePredicates() map[string]GamePredicate {
	return map[string]GamePredicate{
		"is_paused":  (*Game).IsPaused,
		"is_private": (*Game).IsPrivate,
		"is_public":  (*Game).IsPublic,
		"is_started": (*Game).IsStarted,
		"is_stopped": (*Game).IsStopped,
	}
}

// IsPaused
func (game *Game) IsPaused() bool {
	return false // TODO
}

// IsPrivate
func (game *Game) IsPrivate() bool {
	return false // TODO
}

// IsPublic
func (game *Game) IsPublic() bool {
	return !game.IsPrivate()
}

// IsStarted
func (game *Game) IsStarted() bool {
	return false // TODO
}

// IsStopped
func (game *Game) IsStopped() bool {
	return false // TODO
}
