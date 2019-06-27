/* This is free and unencumbered software released into the public domain. */

package gdk

// Agent
type Agent struct {
	Object Object
}

// AgentPredicate
type AgentPredicate func(*Agent) bool

// AgentPredicates
func AgentPredicates() map[string]AgentPredicate {
	return map[string]AgentPredicate{
		"is_player":  (*Agent).IsPlayer,
		"is_robot":   (*Agent).IsRobot,
		"can_fly":    (*Agent).CanFly,
		"can_move":   (*Agent).CanMove,
		"has_legs":   (*Agent).HasLegs,
		"has_wings":  (*Agent).HasWings,
		"has_wheels": (*Agent).HasWheels,
	}
}

// IsPlayer
func (agent *Agent) IsPlayer() bool {
	return false // TODO
}

// IsRobot
func (agent *Agent) IsRobot() bool {
	return false // TODO
}

// CanFly
func (agent *Agent) CanFly() bool {
	return false // TODO
}

// CanMove
func (agent *Agent) CanMove() bool {
	return false // TODO
}

// HasLegs
func (agent *Agent) HasLegs() bool {
	return false // TODO
}

// HasWings
func (agent *Agent) HasWings() bool {
	return false // TODO
}

// HasWheels
func (agent *Agent) HasWheels() bool {
	return false // TODO
}
