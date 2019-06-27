/* This is free and unencumbered software released into the public domain. */

package gdk

// Player
type Player struct {
	Agent Agent
}

// PlayerPredicate
type PlayerPredicate func(*Player) bool

// PlayerPredicates
func PlayerPredicates() map[string]PlayerPredicate {
	return map[string]PlayerPredicate{
		"is_alive": (*Player).IsAlive,
		"is_dead":  (*Player).IsDead,
	}
}

// IsAlive
func (player *Player) IsAlive() bool {
	return true // TODO
}

// IsDead
func (player *Player) IsDead() bool {
	return !player.IsAlive()
}
