/* This is free and unencumbered software released into the public domain. */

package gdk

// Object
type Object interface{}

// ObjectPredicate
type ObjectPredicate func(Object) bool

// ObjectPredicates
func ObjectPredicates() map[string]ObjectPredicate {
	return map[string]ObjectPredicate{}
}
