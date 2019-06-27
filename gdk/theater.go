/* This is free and unencumbered software released into the public domain. */

package gdk

// Theater
type Theater interface{}

// TheaterPredicate
type TheaterPredicate func(Theater) bool

// TheaterPredicates
func TheaterPredicates() map[string]TheaterPredicate {
	return map[string]TheaterPredicate{}
}
