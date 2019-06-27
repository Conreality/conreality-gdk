/* This is free and unencumbered software released into the public domain. */

package gdk

// Unit
type Unit interface {
	IsAlive() bool
	IsDead() bool
}

// UnitPredicate
type UnitPredicate func(Unit) bool

// UnitPredicates
func UnitPredicates() map[string]UnitPredicate {
	return map[string]UnitPredicate{
		"is_alive": func(unit Unit) bool {
			return unit.IsAlive()
		},
		"is_dead": func(unit Unit) bool {
			return unit.IsDead()
		},
	}
}
