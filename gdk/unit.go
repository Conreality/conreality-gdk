/* This is free and unencumbered software released into the public domain. */

package gdk

// Unit
type Unit struct{}

// UnitPredicate
type UnitPredicate func(*Unit) bool

// UnitPredicates
func UnitPredicates() map[string]UnitPredicate {
	return map[string]UnitPredicate{
		"is_alive": (*Unit).IsAlive,
		"is_dead":  (*Unit).IsDead,
	}
}

// IsAlive
func (unit *Unit) IsAlive() bool {
	return true // TODO
}

// IsDead
func (unit *Unit) IsDead() bool {
	return !unit.IsAlive()
}
