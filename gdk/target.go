/* This is free and unencumbered software released into the public domain. */

package gdk

// Target
type Target interface{}

// TargetPredicate
type TargetPredicate func(Target) bool

// TargetPredicates
func TargetPredicates() map[string]TargetPredicate {
	return map[string]TargetPredicate{}
}
