/* This is free and unencumbered software released into the public domain. */

package gdk

// Target
type Target struct {
	Object Object
}

// TargetPredicate
type TargetPredicate func(*Target) bool
