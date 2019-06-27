/* This is free and unencumbered software released into the public domain. */

package gdk

// Player
type Player interface {
	IsAlive() bool
	IsDead() bool
}

// PlayerPredicate
type PlayerPredicate func(Player) bool

// PlayerPredicates
func PlayerPredicates() map[string]PlayerPredicate {
	return map[string]PlayerPredicate{
		"is_alive": func(player Player) bool {
			return player.IsAlive()
		},
		"is_dead": func(player Player) bool {
			return player.IsDead()
		},
	}
}
